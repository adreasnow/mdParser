package h1

import (
	"bytes"
	"html/template"
	"strings"

	"github.com/adreasnow/mdParser/types"
)

type H1 struct {
	Style   template.HTML
	Content template.HTML
}

func getContents(line string) string {
	return line[2:]
}

func (h1 *H1) Parse(in string, t *template.Template, styles *types.Styles) (string, error) {
	var out []string
	for _, line := range strings.Split(in, "\n") {
		if strings.HasPrefix(line, "# ") {
			content := getContents(line)

			h1 := H1{
				Style:   template.HTML(styles.H1),
				Content: template.HTML(content),
			}

			buf := new(bytes.Buffer)

			err := t.ExecuteTemplate(buf, "h1", h1)
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
