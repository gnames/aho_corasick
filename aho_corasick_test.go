package aho_corasick_test

import (
	"testing"

	"github.com/gnames/aho_corasick"
	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	ac := aho_corasick.New()
	ac.Setup([]string{"aba", "cla", "ac", "gee", "lan"})
	matches := ac.Search("aclant", false)
	assert.Equal(t, len(matches), 3)
}
