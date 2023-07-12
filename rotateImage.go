package main

func rotateImage(matrix [][]int) {
	rotateCircle(matrix, 0)
}

func rotateCircle(matrix [][]int, index int) {
	sizeToRotate := len(matrix) - 1 - index*2
	if sizeToRotate <= 0 {
		return
	}
	temp := 0
	for i := 0; i < sizeToRotate; i++ {
		temp = matrix[index][index+i]
		matrix[index][index+i] = matrix[len(matrix)-index-1-i][index]
		matrix[len(matrix)-index-1-i][index] = matrix[len(matrix)-index-1][len(matrix)-index-1-i]
		matrix[len(matrix)-index-1][len(matrix)-index-1-i] = matrix[index+i][len(matrix)-index-1]
		matrix[index+i][len(matrix)-index-1] = temp
	}
	rotateCircle(matrix, index+1)

}

func main() {

}
