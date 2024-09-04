package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/iancoleman/strcase"
)

func main() {
	var (
		servicePath   string
		removeHandler bool
	)

	flag.StringVar(&servicePath, "service_path", "", "path to service directory")
	flag.BoolVar(&removeHandler, "remove_handler", false, "remove handler source file")
	flag.Parse()

	handlerFilePath := filepath.Join(servicePath, "handler.go")
	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset, handlerFilePath, nil, parser.ParseComments)
	if err != nil {
		panic(fmt.Errorf("parse handler.go error: %w", err))
	}

	handlerDirPath := filepath.Join(servicePath, "handler")

	dir := filepath.Join(servicePath, "handler")
	if err = os.MkdirAll(dir, 0o755); err != nil {
		panic(fmt.Errorf("create dir %s: %w", dir, err))
	}

	var importNode *ast.GenDecl

	for _, f := range node.Decls {
		switch decl := f.(type) {
		case *ast.GenDecl:
			switch decl.Tok {
			case token.IMPORT:
				importNode = decl
			case token.TYPE:
				writeNodes(filepath.Join(handlerDirPath, "handler.go"), importNode, decl)
			default:
			}
		case *ast.FuncDecl:
			path := filepath.Join(handlerDirPath, strcase.ToSnake(decl.Name.Name)+".go")
			writeNodes(path, importNode, decl)
		}
	}

	err = exec.Command("goimports", "-w", handlerDirPath).Run()
	if err != nil {
		panic(fmt.Errorf("goimports format error: %w", err))
	}

	if removeHandler {
		err = os.Remove(handlerFilePath)
		if err != nil {
			panic(fmt.Errorf("remove handler.go error: %w", err))
		}
	}
}

func writeNodes(path string, nodes ...ast.Decl) {
	if _, err := os.Stat(path); err == nil {
		return
	} else if !os.IsNotExist(err) {
		panic(fmt.Errorf("check %s error: %w", path, err))
	}

	file, err := os.Create(path)
	if err != nil {
		panic(fmt.Errorf("create %s error: %w", path, err))
	}

	defer func() { _ = file.Close() }()

	writeFile(file, "package handler\n\n")

	for _, node := range nodes {
		writeDoc(file, node)

		err = format.Node(file, token.NewFileSet(), node)
		if err != nil {
			panic(fmt.Errorf("format %s error: %w", path, err))
		}

		writeFile(file, "\n")
	}
}

func writeFile(file *os.File, s string) {
	_, err := file.WriteString(s)
	if err != nil {
		panic(fmt.Errorf("write file %s error: %w", file.Name(), err))
	}
}

func writeDoc(file *os.File, node ast.Decl) {
	switch n := node.(type) {
	case *ast.FuncDecl:
		if n.Doc != nil {
			writeFile(file, "\n// ")
			writeFile(file, n.Doc.Text())
			n.Doc = nil
		}
	case *ast.GenDecl:
		if n.Doc != nil {
			writeFile(file, "\n// ")
			writeFile(file, n.Doc.Text())
			n.Doc = nil
		}
	}
}
