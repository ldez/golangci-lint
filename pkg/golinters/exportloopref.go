package golinters

import (
	"github.com/golangci/golangci-lint/pkg/config"
	"github.com/kyoh86/exportloopref"
	"github.com/kyoh86/looppointer"
	"golang.org/x/tools/go/analysis"

	"github.com/golangci/golangci-lint/pkg/golinters/goanalysis"
)

func NewExportLoopRef(settings *config.ExportLoopRefSettings) *goanalysis.Linter {
	a := exportloopref.Analyzer
	if settings != nil && settings.LooppointerMode {
		a = looppointer.Analyzer
	}

	return goanalysis.NewLinter(
		exportloopref.Analyzer.Name,
		exportloopref.Analyzer.Doc,
		[]*analysis.Analyzer{a},
		nil,
	).WithLoadMode(goanalysis.LoadModeTypesInfo)
}
