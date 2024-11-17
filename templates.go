package mdParser

import (
	"html/template"
	"io"
	"log/slog"
	"strings"
)

func (p *Parser) loadTemplates(templates io.Reader) error {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, templates)
	if err != nil {
		slog.Error("could not load templates from the reader", slog.Any("error", err))
	}

	p.templates, err = template.New("htmlFragments").Parse(buf.String())
	if err != nil {
		slog.Error("could not parse templates", slog.Any("error", err))
		return err
	}

	return nil
}
