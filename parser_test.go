package mdParser

import (
	"bytes"
	"embed"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yosssi/gohtml"
)

//go:embed tests/*
var testFiles embed.FS

func TestParse(t *testing.T) {
	t.Parallel()

	stylesReader := bytes.NewReader(stylesFile)
	templatesReader := bytes.NewReader(templatesFile)

	p, err := New(stylesReader, templatesReader)
	require.NoError(t, err)

	t.Run("headings", func(t *testing.T) {
		t.Parallel()

		md, err := testFiles.ReadFile("tests/headings.md")
		require.NoError(t, err)
		expectedHtml, err := testFiles.ReadFile("tests/headings.html")
		require.NoError(t, err)

		html, err := p.Parse(string(md))
		require.NoError(t, err)

		assert.Equal(t, gohtml.Format(string(expectedHtml)), html)
	})
}
