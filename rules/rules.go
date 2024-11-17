package rules

import (
	"html/template"

	"github.com/adreasnow/mdParser/rules/ruleDefinitions/h1"
	"github.com/adreasnow/mdParser/rules/ruleDefinitions/h2"
	"github.com/adreasnow/mdParser/types"
)

type rules interface {
	Parse(input string, t *template.Template, styles *types.Styles) (string, error)
}

var Rules []rules = []rules{
	&h1.H1{},
	&h2.H2{},
}
