package main

import "fmt"

func generate(numRows int) [][]int {
	var res = make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		res[i] = make([]int, i+1)
		res[i][0], res[i][i] = 1, 1
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j] + res[i-1][j-1]
		}
	}
	return res
}

func GanjilGenap(n int) bool {
	if n%2 != 0 {
		return true
	}
	return false
}

func SAGanjilGenap() {
	pascal := generate(15)
	for _, a := range pascal {
		for _, b := range a {

			str := "gn"
			val := GanjilGenap(b)
			if val {
				str = "gj"
			}
			fmt.Print(b)
			if b != 1 {
				fmt.Print(str)
			}
			fmt.Print(" ")
		}
		fmt.Println("")
	}
}

func SACount() {
	pascal := generate(15)
	for _, a := range pascal {
		var val int
		for _, b := range a {
			fmt.Print(b)

			if b != 1 {
				val += b
			}
			fmt.Print(" ")
		}
		fmt.Print(" = ")
		fmt.Print(val)
		fmt.Println("")
	}
}

func SARevers(row int) {

	pascal := generate(row)

	result := make([][]int, row)
	for i := row; i > 1; i-- {
		result[row-i] = pascal[i-1]
	}

	for _, a := range result {
		for _, b := range a {
			fmt.Print(b)
			fmt.Print(" ")
		}
		fmt.Print(" = ")
		fmt.Println("")
	}
}
