package h2

import (
	"bytes"
	"html/template"
	"strings"

	"github.com/adreasnow/mdParser/types"
)

type H2 struct {
	Style   template.HTML
	Content template.HTML
}

func getContents(line string) string {
	return line[3:]
}

func (h1 *H2) Parse(in string, t *template.Template, styles *types.Styles) (string, error) {
	var out []string
	for _, line := range strings.Split(in, "\n") {
		if strings.HasPrefix(line, "## ") {
			content := getContents(line)

			h2 := H2{
				Style:   template.HTML(styles.H2),
				Content: template.HTML(content),
			}

			buf := new(bytes.Buffer)

			err := t.ExecuteTemplate(buf, "h2", h2)
			if err != nil {
				return "", err
			}
			out = append(out, buf.String())
		} else {
			out = append(out, line)
		}
	}
	return strings.Join(out, "\n"), nil
}
