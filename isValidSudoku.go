package main

import "fmt"

func isValidSudoku(board [][]byte) bool {

	result := checkRow(board)
	result = result && checkCol(board)
	return result && checkSquare(board)
}

func checkRow(board [][]byte) bool {
	result := true
	for _, row := range board {
		result = result && isValidGroup(row)
	}
	return result
}

func checkCol(board [][]byte) bool {
	col := make([]byte, 9)
	result := true
	for i := 0; i < 9; i++ {
		for index, _ := range board {
			col[index] = board[index][i]
		}
		result = result && isValidGroup(col)
	}

	return result
}

func checkSquare(board [][]byte) bool {
	sqr := make([]byte, 9)
	result := true
	for i := 0; i < 9; i++ {
		/*		fmt.Println("sqr ", i)
		 */for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				/*				fmt.Println(board[i/3*3+j][i%3*3+k])
				 */sqr[3*j+k] = board[i/3*3+j][i%3*3+k]
			}
		}
		result = result && isValidGroup(sqr)
	}
	return result
}
func isValidGroup(group []byte) bool {
	//for _, b := range group {
	//	fmt.Printf("%c ", b)
	//}
	//
	//fmt.Println()
	occurrence := [9]bool{}
	for _, val := range group {
		if val >= 49 {
			val -= 49
		} else {
			continue
		}
		if occurrence[val] {
			//fmt.Println(false)
			return false
		}
		occurrence[val] = true
	}
	//fmt.Println(true)
	return true
}

func main() {
	input := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	fmt.Println(isValidSudoku(input))

}
