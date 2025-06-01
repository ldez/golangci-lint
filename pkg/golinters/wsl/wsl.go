package wsl

import (
	"bytes"
	"slices"

	wslv4 "github.com/bombsimon/wsl/v4"
	wslv5 "github.com/bombsimon/wsl/v5"
	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
	"github.com/golangci/golangci-lint/v2/pkg/golinters/internal"
	"gopkg.in/yaml.v3"
)

// Deprecated
func NewV4(settings *config.WSLv4Settings) *goanalysis.Linter {
	var conf *wslv4.Configuration

	if settings != nil {
		conf = &wslv4.Configuration{
			StrictAppend:                     settings.StrictAppend,
			AllowAssignAndCallCuddle:         settings.AllowAssignAndCallCuddle,
			AllowAssignAndAnythingCuddle:     settings.AllowAssignAndAnythingCuddle,
			AllowMultiLineAssignCuddle:       settings.AllowMultiLineAssignCuddle,
			ForceCaseTrailingWhitespaceLimit: settings.ForceCaseTrailingWhitespaceLimit,
			AllowTrailingComment:             settings.AllowTrailingComment,
			AllowSeparatedLeadingComment:     settings.AllowSeparatedLeadingComment,
			AllowCuddleDeclaration:           settings.AllowCuddleDeclaration,
			AllowCuddleWithCalls:             settings.AllowCuddleWithCalls,
			AllowCuddleWithRHS:               settings.AllowCuddleWithRHS,
			ForceCuddleErrCheckAndAssign:     settings.ForceCuddleErrCheckAndAssign,
			AllowCuddleUsedInBlock:           settings.AllowCuddleUsedInBlock,
			ErrorVariableNames:               settings.ErrorVariableNames,
			ForceExclusiveShortDeclarations:  settings.ForceExclusiveShortDeclarations,
			IncludeGenerated:                 true, // force to true because golangci-lint already have a way to filter generated files.
		}
	}

	return goanalysis.
		NewLinterFromAnalyzer(wslv4.NewAnalyzer(conf)).
		WithLoadMode(goanalysis.LoadModeSyntax)
}

type fakeRoot struct {
	Settings v5YAML `yaml:"wsl_v5"`
}

type v5YAML struct {
	AllowFirstInBlock bool     `yaml:"allow-first-in-block,omitempty"`
	AllowWholeBlock   bool     `yaml:"allow-whole-block,omitempty"`
	BranchMaxLines    int      `yaml:"branch-max-lines,omitempty"`
	CaseMaxLines      int      `yaml:"case-max-lines,omitempty"`
	Enable            []string `yaml:"enable,omitempty"`
	Disable           []string `yaml:"disable,omitempty"`
}

func Migration(old *config.WSLv4Settings) string {
	cfg := fakeRoot{Settings: v5YAML{
		AllowFirstInBlock: true,
		AllowWholeBlock:   false,
		BranchMaxLines:    2,
		CaseMaxLines:      old.ForceCaseTrailingWhitespaceLimit,
	}}

	if !old.StrictAppend {
		cfg.Settings.Disable = append(cfg.Settings.Disable, wslv5.CheckAppend.String())
	}

	if old.AllowAssignAndAnythingCuddle {
		cfg.Settings.Disable = append(cfg.Settings.Disable, wslv5.CheckAssign.String())
	}

	if old.AllowMultiLineAssignCuddle {
		internal.LinterLogger.Warnf("`allow-multiline-assign` is deprecated and always allowed in wsl >= v5")
	}

	if old.AllowTrailingComment {
		internal.LinterLogger.Warnf("`allow-trailing-comment` is deprecated and always allowed in wsl >= v5")
	}

	if old.AllowSeparatedLeadingComment {
		internal.LinterLogger.Warnf("`allow-separated-leading-comment` is deprecated and always allowed in wsl >= v5")
	}

	if old.AllowCuddleDeclaration {
		cfg.Settings.Disable = append(cfg.Settings.Disable, wslv5.CheckDecl.String())
	}

	if old.ForceCuddleErrCheckAndAssign {
		cfg.Settings.Enable = append(cfg.Settings.Enable, wslv5.CheckErr.String())
	}

	if old.ForceExclusiveShortDeclarations {
		cfg.Settings.Enable = append(cfg.Settings.Enable, wslv5.CheckAssignExclusive.String())
	}

	if !old.AllowAssignAndCallCuddle {
		cfg.Settings.Enable = append(cfg.Settings.Enable, wslv5.CheckAssignExpr.String())
	}

	slices.Sort(cfg.Settings.Enable)
	slices.Sort(cfg.Settings.Disable)

	buf := bytes.NewBuffer([]byte{})
	encoder := yaml.NewEncoder(buf)
	encoder.SetIndent(2)

	err := encoder.Encode(cfg)
	if err != nil {
		internal.LinterLogger.Fatalf("wsl: invalid config: %v", err)
	}

	return buf.String()
}
