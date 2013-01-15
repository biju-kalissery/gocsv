package gocsv

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

type a interface{
	GetRows()
	GetColumns()
	GetRow(row int)
	GetColumn(col int)
	Next()
	GetString(row int)
}


