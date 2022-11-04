package linters

import (
	"fmt"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

var MaxLines int = 800

var FileLinesAnalyzer = &analysis.Analyzer{
	Name: "fll",
	Doc:  "max file line length",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	pass.Fset.Iterate(func(file *token.File) bool {
		lines := file.LineCount()
		if file.LineCount() > MaxLines {
			pass.Report(analysis.Diagnostic{
				Pos:            0,
				End:            0,
				Category:       "file max lines",
				Message:        fmt.Sprintf("file has %d lines which is out of limit %d lines", lines, MaxLines),
				SuggestedFixes: nil,
			})
		}
		return true
	})

	return nil, nil
}
