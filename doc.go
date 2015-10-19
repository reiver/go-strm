/*
Package strm provides an opinionated framework for processing both finite and infinite data tables as streams.

For example:

	package main
	
	// Note that dot imports are being used.
	import (
		. "github.com/reiver/go-strm"
		. "github.com/reiver/strm-append"
		. "github.com/reiver/strm-csv"
		. "github.com/reiver/strm-head"
		. "github.com/reiver/strm-select"
		. "github.com/reiver/strm-stdout"
	)
	
	func main() {
		Begin(CSV, "table.csv").
			Strm(HEAD, 100).
			Strm(SELECT, "color", "size")
			Strm(APPEND, "blue",  5)
		End(STDOUT, "tsv")
	}

*/
package strm
