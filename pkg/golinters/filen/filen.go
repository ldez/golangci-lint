package filen

import (
	"github.com/DanilXO/filen/pkg/filen"
	"golang.org/x/tools/go/analysis"

	"github.com/golangci/golangci-lint/pkg/config"
	"github.com/golangci/golangci-lint/pkg/goanalysis"
)

func New(settings *config.FilenSettings) *goanalysis.Linter {
	a := filen.NewAnalyzer(&filen.Runner{
		MaxLines:       settings.MaxLines,
		MinLines:       settings.MinLines,
		IgnoreComments: settings.IgnoreComments,
	})

	return goanalysis.NewLinter(
		a.Name,
		a.Doc,
		[]*analysis.Analyzer{a},
		nil,
	).WithLoadMode(goanalysis.LoadModeSyntax)
}
