package aho_corasick_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gnames/aho_corasick"
	"github.com/gnames/gnsys"
	"github.com/stretchr/testify/assert"
)

func TestDumpLoad(t *testing.T) {
	var err error
	patterns := []string{"aba", "cla", "ac", "gee", "lan"}
	ac := aho_corasick.New()
	ac.Setup(patterns)
	haystack := "abacgeeaba"
	matches := ac.Search(haystack)
	assert.Equal(t, len(matches), 4)

	dir := filepath.Join(os.TempDir(), "ahoco")
	err = gnsys.MakeDir(dir)
	assert.Nil(t, err)
	err = ac.Dump(dir)
	assert.Nil(t, err)
}
