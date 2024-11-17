package mdParser

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Parallel()

	stylesReader := bytes.NewReader(stylesFile)
	templatesReader := bytes.NewReader(templatesFile)

	p, err := New(stylesReader, templatesReader)
	require.NoError(t, err)

	assert.Equal(t, "style 1", p.styles.H1)
	assert.Contains(t, p.templates.DefinedTemplates(), "h1")
}
