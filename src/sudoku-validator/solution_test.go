package sudoku_validator

import (
	"fmt"
	"testing"
)

type TestCase struct {
	name     string
	sudoku   [9][9]int32
	expected bool
}

func TestSudokuValidator(t *testing.T) {
	dataTable := []TestCase{
		{
			"Empty",
			[9][9]int32{
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			true,
		},
		{
			"Completed",
			[9][9]int32{
				{4, 3, 5, 2, 6, 9, 7, 8, 1},
				{6, 8, 2, 5, 7, 1, 4, 9, 3},
				{1, 9, 7, 8, 3, 4, 5, 6, 2},
				{8, 2, 6, 1, 9, 5, 3, 4, 7},
				{3, 7, 4, 6, 8, 2, 9, 1, 5},
				{9, 5, 1, 7, 4, 3, 6, 2, 8},
				{5, 1, 9, 3, 2, 6, 8, 7, 4},
				{2, 4, 8, 9, 5, 7, 1, 3, 6},
				{7, 6, 3, 4, 1, 8, 2, 5, 9},
			},
			true,
		},
		{
			"Incomplete",
			[9][9]int32{
				{4, 0, 5, 0, 6, 0, 0, 8, 0},
				{6, 0, 2, 5, 7, 0, 4, 9, 3},
				{1, 9, 0, 8, 0, 0, 0, 6, 0},
				{8, 2, 0, 1, 9, 0, 3, 4, 7},
				{3, 0, 0, 6, 0, 2, 0, 1, 0},
				{0, 0, 1, 0, 4, 0, 0, 2, 0},
				{5, 1, 9, 0, 0, 6, 0, 7, 4},
				{2, 0, 0, 9, 0, 0, 1, 0, 0},
				{7, 6, 0, 4, 0, 0, 0, 5, 9},
			},
			true,
		},
		{
			"Repeated on first row",
			[9][9]int32{
				{4, 0, 5, 0, 6, 0, 0, 8, 4},
				{6, 0, 2, 5, 7, 0, 4, 9, 3},
				{1, 9, 0, 8, 0, 0, 0, 6, 0},
				{8, 2, 0, 1, 9, 0, 3, 4, 7},
				{3, 0, 0, 6, 0, 2, 0, 1, 0},
				{0, 0, 1, 0, 4, 0, 0, 2, 0},
				{5, 1, 9, 0, 0, 6, 0, 7, 4},
				{2, 0, 0, 9, 0, 0, 1, 0, 0},
				{7, 6, 0, 4, 0, 0, 0, 5, 9},
			},
			false,
		},
		{
			"Repeated on last row",
			[9][9]int32{
				{4, 0, 5, 0, 6, 0, 0, 8, 0},
				{6, 0, 2, 5, 7, 0, 4, 9, 3},
				{1, 9, 0, 8, 0, 0, 0, 6, 0},
				{8, 2, 0, 1, 9, 0, 3, 4, 7},
				{3, 0, 0, 6, 0, 2, 0, 1, 0},
				{0, 0, 1, 0, 4, 0, 0, 2, 0},
				{5, 1, 9, 0, 0, 6, 0, 7, 4},
				{2, 0, 0, 9, 0, 0, 1, 0, 0},
				{7, 6, 0, 4, 0, 4, 0, 5, 9},
			},
			false,
		},
		{
			"Repeated on first column",
			[9][9]int32{
				{4, 0, 5, 0, 6, 0, 0, 8, 0},
				{6, 0, 2, 5, 7, 0, 4, 9, 3},
				{1, 9, 0, 8, 0, 0, 0, 6, 0},
				{8, 2, 0, 1, 9, 0, 3, 4, 7},
				{3, 0, 0, 6, 0, 2, 0, 1, 0},
				{0, 0, 1, 0, 4, 0, 0, 2, 0},
				{5, 1, 9, 0, 0, 6, 0, 7, 4},
				{2, 0, 0, 9, 0, 0, 1, 0, 0},
				{1, 6, 0, 4, 0, 0, 0, 5, 9},
			},
			false,
		},
		{
			"Repeated on last column",
			[9][9]int32{
				{4, 0, 5, 0, 6, 0, 0, 8, 0},
				{6, 0, 2, 5, 7, 0, 4, 9, 3},
				{1, 9, 0, 8, 0, 0, 0, 6, 0},
				{8, 2, 0, 1, 9, 0, 3, 4, 7},
				{3, 0, 0, 6, 0, 2, 0, 1, 0},
				{0, 0, 1, 0, 4, 0, 0, 2, 0},
				{5, 1, 9, 0, 0, 6, 0, 7, 4},
				{2, 0, 0, 9, 0, 0, 1, 0, 0},
				{7, 6, 0, 4, 0, 0, 0, 5, 3},
			},
			false,
		},
	}

	for _, testCase := range dataTable {
		t.Run(fmt.Sprintf("%s", testCase.name), func(t *testing.T) {
			isValid := SudokuValidator(testCase.sudoku)

			if isValid != testCase.expected {
				t.Errorf("expected %t to be %t", isValid, testCase.expected)
			}
		})
	}
}
