Mabetle GoCSV
=============

CSV format file utils. It's very useful when import and export data between database.


Install
-------

	`go get -u github.com/mabetle/gocsv`

Usage
-----

	import "github.com/mabetle/gocsv"

	file:="demo.csv"

	csv:=gocsv.LoadFile(file)
	csv.GetRows()			//return rows, int
	csv.GetColumns()		//return columns, int
	csv.GetHeaderRow()		//return header row, string[]
	csv.GetString(row,column) //return specific row col value, string

Run Demo
--------

	`go run cmd/gocsv_demo/main.go`

BUGS
----


TODO
-----


