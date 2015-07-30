package gocsv

import (
	"testing"
)

// test parse
func TestParse(t *testing.T) {

	file := "demo_read.csv"
	csv := LoadFile(file)
	csv.ShowContent()

	refHeaders := []string{
		"ID", "UserName", "Email", "RealName", "Age",
	}

	actualHeaders := csv.GetHeaderRow()
	for idx, _ := range refHeaders {
		t.Log("Testing header at index:", idx, " for value:", refHeaders[idx])
		if refHeaders[idx] != actualHeaders[idx] {
			t.Error("Expected: ", refHeaders[idx], " got: ", actualHeaders[idx])
		}
	}

	actualColCount := csv.GetColumns()
	expectedColCnt := len(refHeaders)
	if actualColCount != expectedColCnt {
		t.Error("Column count expected: ", expectedColCnt, " got: ", actualColCount)
	}

	actualRowCount := csv.GetRows()
	expectedRowCnt := 3
	if actualRowCount != expectedRowCnt {
		t.Error("Column count expected: ", expectedRowCnt, " got: ", actualRowCount)
	}

	actualRow2Vals := csv.GetRow(2)
	expectedRow2Vals := []string{
		"2", "demo", "demo@demo.com", "Demo测试", "33",
	}

	for idx2, _ := range expectedRow2Vals {
		t.Log("Testing row2 at index:", idx2, " for value:", expectedRow2Vals[idx2])
		if expectedRow2Vals[idx2] != actualRow2Vals[idx2] {
			t.Error("Expected: ", expectedRow2Vals[idx2], " got: ", actualRow2Vals[idx2])
		}
	}

	expectedCell22 := "demo@demo.com"
	actualCell22 := csv.GetString(2, 2)
	if expectedCell22 != actualCell22 {
		t.Error("Expected: ", expectedCell22, " actual: ", actualCell22)
	}
}
