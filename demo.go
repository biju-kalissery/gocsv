package gocsv

import(
	"fmt"
)

func Demo(){

	file:="demo_read.csv"
	csv:=LoadFile(file)
	csv.ShowContent()
	fmt.Println(csv.GetHeaderRow())

	fmt.Println("Columns: ", csv.GetColumns())
	fmt.Println("Rows: ", csv.GetRows())
	fmt.Println("Row 2: ", csv.GetRow(2))

	fmt.Println("Row 2,Column 2: ", csv.GetString(2,2))

}


