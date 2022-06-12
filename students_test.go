package coverage

import (
	"log"
	"os"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

var p = People{
		{
			firstName: "A",
			lastName: "B",
			birthDay: time.Now(),
		},
		{
			firstName: "C",
			lastName: "D",
			birthDay: time.Now().Add(20000),
		},
	}	


func TestLen(t *testing.T) {
	people := p
	if people.Len() != 2 {
		t.Error("not correct func Len()")
	}
}

func TestLess(t *testing.T) {
	b := time.Now()
	people := p
	people[0].birthDay = b
	people[1].birthDay = b
	
	if !p.Less(0, 1) {
		t.Error("error occured in func Less")
	}

	p[0].firstName = p[1].firstName
	if !p.Less(0, 1) {
		t.Error("error occured in func Less")
	}

	p[0].birthDay = time.Now().Add(100)
	if !p.Less(0, 1) {
		t.Error("error occured in func Less")
	}
}

func TestSwap(t *testing.T) {
	people := p
	p1 := people[0]
	p2 := people[1]
	people.Swap(0, 1)
	log.Println(people[0], people[1], p1, p2)
	if people[0] != p2 || people[1] != p1 {
		t.Error("error occured in func Swap(i, j)")
	}
}

func TestNew(t *testing.T) {
	strmtr := `1 2 3
	4 5 6
	7 8 9`
	mtr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	matrix, err := New(strmtr)
	if err != nil {
		t.Errorf("there should be no errors")
	}
	if matrix.rows != 3 || matrix.cols != 3 {
		t.Error("dimensions are not correct")
	}
	for i := range matrix.data {
		if mtr[i] != matrix.data[i] {
			t.Errorf("data is not equal")
		}
	}
	
	strmtr = 
	`
	1
	2 3
	`
	matrix, err = New(strmtr)
	if err == nil {
		t.Errorf("there should be an error here, dimensions are not correct")
	}

	strmtr = 
	`
	1 deef
	eerg 5`
	matrix, err = New(strmtr)
	if err == nil {
		t.Errorf("there should be an error here, data is not an integer")
	}
}

func TestRows(t *testing.T) {
	strmtr := `1 2 3
	4 5 6
	7 8 9`

	mtr := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	
	matrix, _ := New(strmtr)
	rows := matrix.Rows()
	for i := range rows {
		for j := range rows[i] {
			if mtr[i][j] != rows[i][j] {
				t.Error("matrix rows are not equal")
			}
		}
	}
}

func TestCols(t *testing.T) {
	strmtr := `1 2 3
	4 5 6
	7 8 9`

	mtr := [][]int{
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 9},
	}
	
	matrix, _ := New(strmtr)
	cols := matrix.Cols()
	for i := range cols {
		for j := range cols[i] {
			if mtr[i][j] != cols[i][j] {
				t.Error("matrix cols are not equal")
			}
		}
	}
}

func TestSet(t *testing.T) {
	strmtr := `1 2 3
	4 5 6
	7 8 9`

	matrix, _ := New(strmtr)

	if !matrix.Set(0, 0, 5) {
		t.Error("unable to change cell")
	}

	rows := matrix.Rows()
	if rows[0][0] != 5 {
		t.Error("cell did not change")
	}

	if matrix.Set(3, 0, 10) {
		t.Error("set data should not work")
	}

	if matrix.Set(0, 3, 10) {
		t.Error("set data should not work")
	}
}