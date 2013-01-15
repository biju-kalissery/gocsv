Mabetle GoCSV
=============

CSV format file parse utils. Very useful in import ant export data between database


Install
-------

	go get -u github.com/mabetle/gocsv

Usage
-----

	import "github.com/mabetle/gocsv"

	csv:=gocsv.LoadFile(file)
	csv.GetRows()			//return rows, int
	csv.GetColumns()		//return columns, int
	csv.GetHeaderRow()		//return header row, string[]
	csv.GetString(row,column) //return specific row col value, string



Run Demo
--------

	cd MABETLE_GOCSV_PATH
	go run cmd/gocsv_demo/main.go

Bugs
----


Todo
-----


