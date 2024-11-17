package mdParser

import (
	"io"
	"log/slog"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/adreasnow/mdParser/types"
)

func (p *Parser) loadStyles(stylesFile io.Reader) error {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, stylesFile)
	if err != nil {
		slog.Error("could not load styles from the reader", slog.Any("error", err))
	}

	s := new(types.Styles)
	_, err = toml.Decode(buf.String(), s)
	if err != nil {
		slog.Error("could not decode styles from toml", slog.Any("error", err))
		return err
	}

	p.styles = s

	return nil
}
