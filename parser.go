package mdParser

import (
	"github.com/adreasnow/mdParser/rules"
	"github.com/yosssi/gohtml"
)

func (p *Parser) Parse(in string) (string, error) {
	var out string
	var err error

	for _, rule := range rules.Rules {
		out, err = rule.Parse(in, p.templates, p.styles)
		if err != nil {
			return "", err
		}
		in = out
	}

	return gohtml.Format(out), nil
}
