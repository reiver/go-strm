# go-strm

A library that provides an opinionated framework for processing both finite and infinite data tables as streams, for the Go programming language.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-strm

[![GoDoc](https://godoc.org/github.com/reiver/go-strm?status.svg)](https://godoc.org/github.com/reiver/go-strm)

## Example
```
package main

import (
	. "github.com/reiver/go-strm"
	. "github.com/reiver/strm-csv"
	. "github.com/reiver/strm-head"
	. "github.com/reiver/strm-stdout"
)

func main() {
	Begin(CSV, "table.csv").
		Strm(HEAD, 12).
	End(STDOUT, "tsv")
}
```

(Note that in that example dot imports were used.)
