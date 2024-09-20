package logger

import (
	"fmt"
	"os"
	"strings"
)

func Infof(format string, args ...interface{}) {
	var formatBuilder strings.Builder

	formatBuilder.WriteString("[INFO] ")
	formatBuilder.WriteString(format)

	if !strings.HasSuffix(format, "\n") {
		formatBuilder.WriteString("\n")
	}

	_, _ = fmt.Fprintf(os.Stdout, formatBuilder.String(), args...)
}

func Warnf(format string, args ...interface{}) {
	var formatBuilder strings.Builder

	formatBuilder.WriteString("[WARN] ")
	formatBuilder.WriteString(format)

	if !strings.HasSuffix(format, "\n") {
		formatBuilder.WriteString("\n")
	}

	_, _ = fmt.Fprintf(os.Stderr, formatBuilder.String(), args...)
}

func Errorf(format string, args ...interface{}) {
	var formatBuilder strings.Builder

	formatBuilder.WriteString("[ERROR] ")
	formatBuilder.WriteString(format)

	if !strings.HasSuffix(format, "\n") {
		formatBuilder.WriteString("\n")
	}

	_, _ = fmt.Fprintf(os.Stderr, formatBuilder.String(), args...)
	os.Exit(1)
}
