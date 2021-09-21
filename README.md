# aho_corasick

A Go implementation of Aho-Corasick algorithm for efficient multiple pattern
matching within a string.

## Introduction

This project implements the powerful string searching [Aho-Corasick
algorithm](https://dl.acm.org/doi/10.1145/360825.360855) invented by Alfred V.
Aho and Margaret J. Corasick in the Go programming language. The Aho-Corasick
algorithm is useful because it efficiently indexes all occurrences of a list of
keywords within a text string.

This implementation searches at the letter level instead of the word level.
Both [failure links](https://www.youtube.com/watch?v=O7_w001f58c) and
[dictionary links](https://www.youtube.com/watch?v=OFKxWFew_L0) are
implemented.

## Installation

The Go module is installable by running:
```
go get github.com/gnames/aho_corasick
```

## Usage

ereate a new aho_corasick instance with `aho_corasick.New()` and setup the
automaton with the search patterns with `ac.Setup(patterns)`. Run search with
`ac.Setup(patterns)`, which returns an array of matches.

```go
ac := aho_corasick.New()
patterns := []string{"aba", "cla", "ac", "gee", "lan"}
ac.Setup(patterns)
haystack := "abacgeeaba"
matches := ac.Search(haystack)

```

## Development

If you find a bug, please open an
[issue](https://github.com/gnames/aho_corasick/issues) ticket. Pull requests
are welcome.

Tests can be run with `go test` which will produce a text visual of the trie:

```
******* Trie *******

haystack: abacgeeaba

root->root ┬─ a->root ┬─ b->root ── a*->a
           │          └─ c*->c
           ├─ c->root ── l->l ── a*->a
           ├─ g->root ── e->root ── e*->root
           └─ l->root ── a->a ── n*->root

********************
PASS
```

In the trie output, `root` refers to the root node, `*` represents word nodes,
`->` indicates the failure links, `|` indicates dictionary links.


Trie output can also be produced with the debugger, which can be run with:

```go
ac := aho_corasick.New()
haystack := "geeabaclaba"
patterns := []string{"aba", "cla", "ac", "gee", "lan"}
ac.Setup(patterns)
ac.Debug(haystack)
```


## License

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Authors


* [Dmitry Mozzherin]
* [Geoff Ower]

[Dmitry Mozzherin]: https://github.com/dimus
[Geoff Ower]: https://github.com/gdower
