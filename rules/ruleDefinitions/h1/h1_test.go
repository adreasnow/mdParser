package h1

import (
	"html/template"
	"testing"

	"github.com/adreasnow/mdParser/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetContents(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in       string
		expected string
	}{
		{
			in:       "# Hello",
			expected: "Hello",
		},
		{
			in:       "## Hello",
			expected: " Hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			out := getContents(tt.in)
			assert.Equal(t, tt.expected, out)
		})
	}
}

func TestH1(t *testing.T) {
	t.Parallel()

	templ := template.Must(template.New("h1").Parse(`<h1 style="{{ .Style }}">{{ .Content }}</h1>`))
	styles := &types.Styles{H1: "style-h1"}
	tests := []struct {
		in       string
		expected string
	}{
		{
			in:       "# heading 1",
			expected: `<h1 style="style-h1">heading 1</h1>`,
		},
		{
			in:       "### heading 3",
			expected: `### heading 3`,
		},
		{
			in:       " # space+#",
			expected: ` # space+#`,
		},
		{
			in:       "# heading 1\n### heading 3\n # space+#",
			expected: "<h1 style=\"style-h1\">heading 1</h1>\n### heading 3\n # space+#",
		},
	}

	h1 := H1{}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			out, err := h1.Parse(tt.in, templ, styles)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}
