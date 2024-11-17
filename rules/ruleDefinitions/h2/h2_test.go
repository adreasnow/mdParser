package h2

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
			expected: "ello",
		},
		{
			in:       "## Hello",
			expected: "Hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			out := getContents(tt.in)
			assert.Equal(t, tt.expected, out)
		})
	}
}

func TestH2(t *testing.T) {
	t.Parallel()

	templ := template.Must(template.New("h2").Parse(`<h2 style="{{ .Style }}">{{ .Content }}</h2>`))
	styles := &types.Styles{H2: "style-h2"}
	tests := []struct {
		in       string
		expected string
	}{
		{
			in:       "## heading 2",
			expected: `<h2 style="style-h2">heading 2</h2>`,
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
			in:       "## heading 2\n### heading 3\n # space+#",
			expected: "<h2 style=\"style-h2\">heading 2</h2>\n### heading 3\n # space+#",
		},
	}

	h2 := H2{}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			out, err := h2.Parse(tt.in, templ, styles)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}
