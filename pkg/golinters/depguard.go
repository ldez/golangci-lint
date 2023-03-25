package golinters

import (
	//nolint:staticcheck // require changes in github.com/OpenPeeDeeP/depguard

	"fmt"

	"github.com/OpenPeeDeeP/depguard/v2"
	"github.com/mitchellh/hashstructure"
	"golang.org/x/tools/go/analysis"

	"github.com/golangci/golangci-lint/pkg/config"
	"github.com/golangci/golangci-lint/pkg/golinters/goanalysis"
)

func NewDepguard(settings *config.DepGuardSettings) *goanalysis.Linter {
	cfg := depGuardSettings(settings)

	a, err := depguard.NewAnalyzer(&cfg)
	if err != nil {
		linterLogger.Fatalf("depguard: create analyzer: %v", err)
	}

	a.Name = "depguard"

	return goanalysis.NewLinter(
		"depguard",
		a.Doc,
		[]*analysis.Analyzer{a},
		nil,
	).WithLoadMode(goanalysis.LoadModeSyntax)
}

func depGuardSettings(settings *config.DepGuardSettings) depguard.LinterSettings {
	cfg := depguard.LinterSettings{}
	if settings == nil {
		return cfg
	}

	depGuardCompatibility("Main", cfg, *settings)

	return cfg
}

func depGuardCompatibility(name string, acc depguard.LinterSettings, settings config.DepGuardSettings) {
	switch settings.ListType {
	case "allowlist", "whitelist":
		acc[name] = &depguard.List{Allow: settings.Packages, Files: []string{"*.go", "**/*.go"}}

	case "denylist", "blacklist", "":
		l := &depguard.List{Deny: settings.PackagesWithErrorMessage, Files: []string{"*.go", "**/*.go"}}
		if l.Deny == nil {
			l.Deny = map[string]string{}
		}

		for _, p := range settings.Packages {
			l.Deny[p] = "in the denylist"
		}

		acc[name] = l
	}

	for _, guard := range settings.AdditionalGuards {
		h, _ := hashstructure.Hash(guard, nil)
		depGuardCompatibility(fmt.Sprintf("Additional_%d", h), acc, guard)
	}
}
