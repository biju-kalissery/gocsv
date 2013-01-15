package gocsv

import(
	"encoding/csv"
	"strings"
	"io/ioutil"
)

var(
	row []string
	goCSV *CSV
)


// load csv from file
func LoadFile(file string) *CSV{
	goCSV = &CSV{}

	fileContent, _ := ioutil.ReadFile(file)

	csvReader := csv.NewReader(strings.NewReader(string(fileContent)))

	goCSV.headerRow, _ = csvReader.Read()

	goCSV.data, _ =csvReader.ReadAll()

	return goCSV
}

func Parse(){

}



func (csv CSV) GetHeaderRow()[]string{
	return csv.headerRow
}

func (csv CSV) GetData()[][]string{
	return csv.data
}




func (csv CSV) GetRows() int {
	return len(csv.data)
}

func (csv CSV) GetColumns() int {
	return len(csv.headerRow)
}



func (csv CSV) GetRow(row int)[]string{
	if row>csv.GetRows() {
		return nil
	}

	return csv.GetData()[row]

}

func (csv CSV)GetString(row int, column int)string{

	if row > csv.GetRows(){
		return ""
	}

	if column > csv.GetColumns(){
		return ""
	}

	return csv.GetRow(row)[column]
}

