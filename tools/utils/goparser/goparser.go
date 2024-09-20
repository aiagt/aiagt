package goparser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

type StructMethod struct {
	StructName string
	Methods    []string
}

func ParseStructMethods(filePath string) (map[string]*StructMethod, error) {
	structMethods := make(map[string]*StructMethod)

	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset, filePath, nil, parser.AllErrors)
	if err != nil {
		return nil, fmt.Errorf("parse file %s error: %v", filePath, err)
	}

	ast.Inspect(node, func(n ast.Node) bool {
		if funcDecl, ok := n.(*ast.FuncDecl); ok && funcDecl.Recv != nil {
			if recvList := funcDecl.Recv.List; len(recvList) > 0 {
				recvType := recvList[0].Type
				var structName string

				switch expr := recvType.(type) {
				case *ast.Ident:
					structName = expr.Name
				case *ast.StarExpr:
					if ident, ok := expr.X.(*ast.Ident); ok {
						structName = ident.Name
					}
				}

				if structName != "" {
					methodName := funcDecl.Name.Name

					if _, exists := structMethods[structName]; !exists {
						structMethods[structName] = &StructMethod{
							StructName: structName,
							Methods:    []string{},
						}
					}

					structMethods[structName].Methods = append(structMethods[structName].Methods, methodName)
				}
			}
		}

		return true
	})

	return structMethods, nil
}

func ParseGoFilesInDir(dir string) (map[string]*StructMethod, error) {
	structMethods := make(map[string]*StructMethod)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			fileStructMethods, err := ParseStructMethods(path)
			if err != nil {
				return err
			}

			for structName, methods := range fileStructMethods {
				if _, exists := structMethods[structName]; !exists {
					structMethods[structName] = methods
				} else {
					structMethods[structName].Methods = append(structMethods[structName].Methods, methods.Methods...)
				}
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return structMethods, nil
}
