package gocsv

// define CSV
type CSV struct{
	data [][]string
	headerRow []string
}

type Reader interface{
	Read(file string)
}


type Writer interface{
	Write(file string,v ...interface{})
}

// interface of CSV
type ICSV interface{
	GetRows() int
	GetColumns() int
	GetRow(row int) []string
	GetColumn(col int) []string
	Next()bool
	GetString(row int,column int) string
}

