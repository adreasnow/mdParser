package mdParser

import (
	"html/template"
	"io"

	"github.com/adreasnow/mdParser/types"
)

type Parser struct {
	templates *template.Template
	styles    *types.Styles
}

func New(styles io.Reader, templates io.Reader) (*Parser, error) {
	p := new(Parser)
	err := p.loadTemplates(templates)
	if err != nil {
		return p, err
	}

	err = p.loadStyles(styles)
	if err != nil {
		return p, err
	}

	return p, nil
}
