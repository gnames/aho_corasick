package aho_corasick_test

import (
	"os"
	"strings"
	"testing"

	"github.com/gnames/aho_corasick"
	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	tests := []struct {
		msg      string
		patterns []string
		nodesNum int
	}{
		{"nodes", []string{"aba", "cla", "ac", "gee", "lan"}, 13},
		{"nodes", []string{}, 0},
		{"nodes", []string{"3jsl", "", " "}, 5},
	}

	for _, v := range tests {
		ac := aho_corasick.New()
		num := ac.Setup(v.patterns)
		assert.Equal(t, num, v.nodesNum)
	}
}

func TestSearch(t *testing.T) {
	ac := aho_corasick.New()
	ac.Setup([]string{"aba", "cla", "ac", "gee", "lan"})
	matches := ac.Search("aclant")
	assert.Equal(t, len(matches), 3)
}

func TestLargeTree(t *testing.T) {
	ac := aho_corasick.New()
	file, err := os.ReadFile("testdata/patterns_test.txt")
	assert.Nil(t, err)
	patterns := strings.Split(strings.TrimSpace(string(file)), "\n")
	for i, p := range patterns {
		patterns[i] = strings.TrimSpace(p)
	}
	ac.Setup(patterns)
	tests := []struct {
		msg, haystack string
		lenRes        int
	}{
		{"1", "aclant", 7},
		{"2", "888adedbofdjdrdpehdpclpidpnaioe333", 30},
	}

	for _, v := range tests {
		matches := ac.Search(v.haystack)
		assert.Equal(t, len(matches), v.lenRes, v.msg)
	}
}
