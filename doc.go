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
			Strm(SELECT, "color", "size").
			Strm(APPEND, "blue",  5).
		End(STDOUT, "tsv")
	}

If in this example "table.csv" contained the data:

	name,size,color
	apple,5,red
	banana,12,yellow
	cherry,1,red

Then Begin(CSV, "table.csv") would cause the following rows to be generated:

	[]interface{"name",   "size", "color"}
	[]interface{"apple",  "5",    "red"}
	[]interface{"banana", "12",   "yellow"}
	[]interface{"cherry", "1",    "red"}

Then Strm(HEAD, 100) says take the first 100 rows. But since there are less than
100 rows it returns what's there.

	[]interface{"name",   "size", "color"}
	[]interface{"apple",  "5",    "red"}
	[]interface{"banana", "12",   "yellow"}
	[]interface{"cherry", "1",    "red"}

Then Strm(SELECT, "color", "size") says only return the "color" and "size" columns,
and in the order "color" 1st and then "size" 2nd.

	[]interface{"color",  "size"}
	[]interface{"red",    "5"}
	[]interface{"yellow", "12"}
	[]interface{"red",    "1"}

Then Strm(APPEND, "blue",  5) says append the row specified to the end, giving us:

	[]interface{"color",  "size"}
	[]interface{"red",    "5"}
	[]interface{"yellow", "12"}
	[]interface{"red",    "1"}
	[]interface{"blue",   5}

And finally End(STDOUT, "tsv") sends the results to the STDOUT (standard out) in
TSV format.

	color	size
	red	5
	yellow	12
	red	1
	blue	5
*/
package strm
