// This must be package main
package main

import (
	"flag"

	linters "github.com/dbraley/example-linter"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/singlechecker"
)

type analyzerPlugin struct{}

// This must be implemented
func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		linters.FileLinesAnalyzer,
	}
}

// This must be defined and named 'AnalyzerPlugin'
var AnalyzerPlugin analyzerPlugin

func main() {
	flag.IntVar(&linters.MaxLines, "max-lines", 800, "max lines of file")
	singlechecker.Main(linters.FileLinesAnalyzer)
}
