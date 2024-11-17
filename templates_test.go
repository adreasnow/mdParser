package mdParser

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed exampleConfig/htmlFragments.html
var templatesFile []byte

func TestLoadTemplates(t *testing.T) {
	t.Parallel()

	p := new(Parser)

	templatesReader := bytes.NewReader(templatesFile)

	err := p.loadTemplates(templatesReader)
	require.NoError(t, err)

	templates := p.templates.DefinedTemplates()
	for _, expected := range []string{
		`"h1"`,
		`"h2"`,
		`"h3"`,
		`"h4"`,
		`"p"`,
		`"code"`,
		`"link"`,
	} {
		assert.Contains(t, templates, expected)
	}
}
