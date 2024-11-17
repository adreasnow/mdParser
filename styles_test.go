package mdParser

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed exampleConfig/styles.toml
var stylesFile []byte

func TestLoadStyles(t *testing.T) {
	t.Parallel()

	p := new(Parser)

	stylesReader := bytes.NewReader(stylesFile)

	err := p.loadStyles(stylesReader)
	require.NoError(t, err)

	assert.Equal(t, "style 1", p.styles.H1)
	assert.Equal(t, "style 2", p.styles.H2)
	assert.Equal(t, "style 3", p.styles.H3)
	assert.Equal(t, "style 4", p.styles.H4)
	assert.Equal(t, "style p", p.styles.P)
	assert.Equal(t, "style code", p.styles.Code)
}
