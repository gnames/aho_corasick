// Package match provides a structure for Aho-Corasick output.
package match

// Match contains data of one pattern matched to a string.
type Match struct {
	// Pattern is a string value of a matched pattern.
	Pattern string
	// PatternIndex is a position of a pattern in the patterns slice
	// used for initialization of the trie.
	PatternIndex int
	// Pos is an optional field that provides start and end positions of the
	// pattern on the sarched string.
	*Pos
}

// Pos contains positional data of a pattern on the search string.
type Pos struct {
	// Start is the byte offset of the start of a pattern on the searched string.
	Start int
	// End is the byte offset of the end of a pattern on the searched string.
	// If End is used for providing a slice, it has to be incremented by one:
	// `haystack[pos.Start:pos.End+1]`
	End int
}

// New creates a new instance of a Match entity. It takes the pattern string,
// its index in the slice used for the trie creation, and the end position of
// the detected pattern occurence.
func New(pattern string, index, end int) Match {
	start := end - len(pattern) + 1
	return Match{
		Pattern:      pattern,
		PatternIndex: index,
		Pos:          &Pos{Start: start, End: end},
	}
}

// Uniq takes a list of matched patterns occurences and returns back unique
// list of them.
func Uniq(matches []Match) []Match {
	matchMap := make(map[string]struct{})
	res := make([]Match, len(matches))

	var count int
	for i := range matches {
		pattern := matches[i].Pattern

		if _, ok := matchMap[pattern]; !ok {
			matchMap[pattern] = struct{}{}
			matches[i].Pos = nil
			res[count] = matches[i]
			count++
		}
	}
	return res[:count]
}
