package internal

import (
	"fmt"
	"sort"
	"testing"
)

type TaskTestCase struct {
	PackSizes    []int
	ItemsOrdered int
	Result       []int
}

func isSlicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Ints(a)
	sort.Ints(b)

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

//func sum(a []int) int {
//	s := 0
//	for i := range s
//}

func TestTask(t *testing.T) {
	var taskTestCases = []TaskTestCase{
		{
			PackSizes:    []int{23, 31, 53},
			ItemsOrdered: 263,
			Result:       []int{23, 23, 31, 31, 31, 31, 31, 31, 31},
		},
		{
			PackSizes:    []int{23, 31, 53},
			ItemsOrdered: 106,
			Result:       []int{53, 53},
		},
		{
			PackSizes:    []int{23, 31, 53},
			ItemsOrdered: 12,
			Result:       []int{23},
		},
		{
			PackSizes:    []int{23, 31, 53},
			ItemsOrdered: 55,
			Result:       []int{31, 31},
		},
		{
			PackSizes:    []int{250, 500, 1000, 2000, 5000},
			ItemsOrdered: 251,
			Result:       []int{500},
		},
		{
			PackSizes:    []int{250, 500, 1000, 2000, 5000},
			ItemsOrdered: 1,
			Result:       []int{250},
		},
		{
			PackSizes:    []int{250, 500, 1000, 2000, 5000},
			ItemsOrdered: 250,
			Result:       []int{250},
		},
		{
			PackSizes:    []int{250, 500, 1000, 2000, 5000},
			ItemsOrdered: 501,
			Result:       []int{250, 500},
		},
		{
			PackSizes:    []int{250, 500, 1000, 2000, 5000},
			ItemsOrdered: 12001,
			Result:       []int{250, 2000, 5000, 5000},
		},
	}
	for i := range taskTestCases {

		testname := fmt.Sprintf("%d", i)
		t.Run(testname, func(t *testing.T) {
			result := Task(taskTestCases[i].PackSizes, taskTestCases[i].ItemsOrdered)
			if !isSlicesEqual(result, taskTestCases[i].Result) {
				t.Errorf("got: %v, expected: %v", result, taskTestCases[i].Result)
			}
		})
	}
}
