package trie_test

import (
	"fmt"
	"sort"
	"testing"

	triepkg "github.com/gnames/aho_corasick/ent/trie"
	"github.com/stretchr/testify/assert"
)

func TestFailLinks(t *testing.T) {
	patterns := []string{"ACC", "ATC", "CAT", "GCG"}
	haystack := "GCATCG"
	trie := triepkg.New(patterns)
	trie.Debug(haystack)
	matches := trie.Search(haystack)
	assert.Equal(t, len(matches), 2)

	res := make([]string, len(matches))
	for i := range matches {
		res[i] = matches[i].Pattern
	}
	sort.Strings(res)
	assert.Equal(t, res, []string{"ATC", "CAT"})
}

func TestDictLinks(t *testing.T) {
	patterns := []string{"A", "AG", "C", "CAA", "GAG", "GC", "GCA"}
	haystack := "GCAA"
	trie := triepkg.New(patterns)
	trie.Debug(haystack)
	matches := trie.Search(haystack)
	fmt.Println(matches)
	assert.Equal(t, len(matches), 6)

	res := make([]string, len(matches))
	for i := range matches {
		res[i] = matches[i].Pattern
	}
	sort.Strings(res)
	assert.Equal(t, res, []string{"A", "A", "C", "CAA", "GC", "GCA"})
}
