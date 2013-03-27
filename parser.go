package gocsv

import(
	"encoding/csv"
	"strings"
	"io/ioutil"
	"fmt"
)

var(
	row []string
	goCSV *CSV
)

// return  *CSV 
func New(file string) *CSV{
	return LoadFile(file)
}

// load csv from file
// return *CSV
func LoadFile(file string) *CSV{
	goCSV = &CSV{}

	fileContent, err := ioutil.ReadFile(file)

	if err != nil{
		fmt.Println("Read File Error")
	}

	csvReader := csv.NewReader(strings.NewReader(string(fileContent)))

	headerRow, errRead := csvReader.Read()

	if errRead != nil{
		fmt.Println("Read Header Error")
	}

	goCSV.headerRow=headerRow


	data, errReadAll :=csvReader.ReadAll()

	if errReadAll != nil {
		fmt.Printf("%v",errReadAll)
	}

	goCSV.data=data

	return goCSV
}

func Parse(){

}

// return CSV HeaderRow
func (csv CSV) GetHeaderRow()[]string{
	return csv.headerRow
}

// return CSV Data
func (csv CSV) GetData()[][]string{
	return csv.data
}

// return CSV Rows
func (csv CSV) GetRows() int {
	return len(csv.data)
}

// return CSV Columns
func (csv CSV) GetColumns() int {
	return len(csv.headerRow)
}

// parameter row
// return CSV Row
func (csv CSV) GetRow(row int)[]string{
	if row>csv.GetRows() {
		return nil
	}
	return csv.GetData()[row]
}


// return CSV row and column value
func (csv CSV)GetString(row int, column int)string{

	if row > csv.GetRows(){
		return ""
	}

	if column > csv.GetColumns(){
		return ""
	}

	return csv.GetRow(row)[column]
}

// show CSV Content
func (csv CSV)ShowContent(){
	fmt.Println(csv.GetHeaderRow())

	rows := csv.GetRows()

	for row := 0 ;row < rows; row++{
		fmt.Println(csv.GetRow(row))
	}
}
