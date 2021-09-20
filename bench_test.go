package aho_corasick_test

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gnames/aho_corasick"
	"github.com/gnames/aho_corasick/ent/match"
)

func BenchmarkSearch(b *testing.B) {
	path := filepath.Join("testdata", "patterns_test.txt")
	count := 1000
	patterns := make([]string, count)
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if count == 0 {
			break
		}
		pattern := strings.TrimSpace(scanner.Text())
		patterns = append(patterns, pattern)
		count--
	}

	ac := aho_corasick.New()
	ac.Setup(patterns)

	b.Run("setup", func(b *testing.B) {

		for i := 0; i < b.N; i++ {
			ac = aho_corasick.New()
			ac.Setup(patterns)
		}

		_ = fmt.Sprintf("%v", len(ac.Search("one")))
	})

	b.Run("search string", func(b *testing.B) {
		haystack := "adjdlsnmasothlskjdsallsjsl&ddjdllslajajaldpppeewrwwopwowowo"
		var matches []match.Match

		for i := 0; i < b.N; i++ {
			ac.Search(haystack)
		}

		_ = fmt.Sprintf("%d", len(matches))
	})

	b.Run("search uniq", func(b *testing.B) {
		haystack := "adjdlsnmasothlskjdsallsjsl&ddjdllslajajaldpppeewrwwopwowowo"
		var matches []match.Match
		for i := 0; i < b.N; i++ {
			matches = ac.SearchUniq(haystack)
		}
		_ = fmt.Sprintf("%d", len(matches))
	})
}
