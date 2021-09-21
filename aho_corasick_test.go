package aho_corasick_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/gnames/aho_corasick"
	"github.com/gnames/aho_corasick/ent/match"
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
	haystack := "abacgeeaba"
	ac.Debug(haystack)
	matches := ac.Search(haystack)
	assert.Equal(t, len(matches), 4)

	res := []match.Match{
		{Pattern: "aba", Pos: &match.Pos{Start: 0, End: 2}, PatternIndex: 0},
		{Pattern: "ac", Pos: &match.Pos{Start: 2, End: 3}, PatternIndex: 2},
		{Pattern: "gee", Pos: &match.Pos{Start: 4, End: 6}, PatternIndex: 3},
		{Pattern: "aba", Pos: &match.Pos{Start: 7, End: 9}, PatternIndex: 0},
	}
	for i, v := range matches {
		assert.Equal(t, v.Pattern, res[i].Pattern)
		assert.Equal(t, v.Start, res[i].Start)
		assert.Equal(t, v.End, res[i].End)
		assert.Equal(t, v.PatternIndex, res[i].PatternIndex)
	}
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
		{"1", "aclant", 8},
		{"2", "888adedbofdjdrdpehdpclpidpnaioe333", 30},
	}

	for _, v := range tests {
		matches := ac.Search(v.haystack)
		assert.Equal(t, len(matches), v.lenRes, v.msg)
	}
}

func Example() {
	// create AhoCorasick instance
	ac := aho_corasick.New()
	// initialize it with string patterns that need to be matched
	ac.Setup([]string{"aba", "cla", "ac", "gee", "lan"})

	// Search for the patterns, providing information of a start and end
	// positions for every found pattern
	matches := ac.Search("aclant")
	fmt.Printf(
		"Pattern: %s, Pos: %d:%d\n",
		matches[0].Pattern,
		matches[0].Start,
		matches[0].End+1,
	)

	// Search and return a unique list of matched patterns
	matchesUnique := ac.SearchUniq("abacgeeaba")
	patterns := make([]string, len(matchesUnique))
	for i, v := range matchesUnique {
		patterns[i] = v.Pattern
	}
	fmt.Printf("Patterns: %s", strings.Join(patterns, ", "))

	// Output:
	// Pattern: ac, Pos: 0:2
	// Patterns: aba, ac, gee
}
