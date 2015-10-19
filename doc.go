/*
Package strm provides an opinionated framework for processing both finite and infinite data tables as streams.

For example:

	package main
	
	// Note that dot imports are being used.
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

*/
package strm
