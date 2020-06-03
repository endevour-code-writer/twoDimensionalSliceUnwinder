package twoDimensionalSliceUnwinder

import (
	"fmt"
	"math/rand"
	"strconv"
)

func Unwind(args []string) []int {
	var unwind []int
	matrix := composeTwoDimensionalSlice(args)
	_, unwind = unwindTwoDimensionalSlice(matrix, unwind)

	return unwind
}

func unwindTwoDimensionalSlice(twoDimensional [][]int, unwind []int) ([][]int, []int) {
	head, tail := twoDimensional[0], twoDimensional[1:]
	unwind = append(unwind, head...)

	slicedValues := make([][]int, len(tail[0]))

	for index, sliceValue := range tail {
		for i := range sliceValue {

			if len(slicedValues[i]) <= 0 {
				slicedValues[i] = make([]int, len(tail))
			}

			slicedValues[i][index] = sliceValue[i]
		}
	}

	for i := len(slicedValues)/2 - 1; i >= 0; i-- {
		opp := len(slicedValues) - 1 - i
		slicedValues[i], slicedValues[opp] = slicedValues[opp], slicedValues[i]
	}

	if 1 == len(slicedValues) {
		unwind = append(unwind, slicedValues[0]...)

		return slicedValues, unwind
	}

	return unwindTwoDimensionalSlice(slicedValues, unwind)

}

func composeTwoDimensionalSlice(args []string) [][]int {
	if 0 == len(args) {
		return createDefaultTwoDimensionalSlice()
	}

	n, err := strconv.Atoi(args[:1][0])

	if err != nil || 1 == n {
		fmt.Printf("The input arg %v is incorrect", n)
		fmt.Println()

		return createDefaultTwoDimensionalSlice()
	}

	matrix := make([][]int, n)

	for i := 0; i < n; i++ {
		rows := make([]int, n)
		for index := range rows {
			rows[index] = rand.Intn(10)
		}

		matrix[i] = rows
	}

	fmt.Printf("The matrix %v was created", matrix)
	fmt.Println()

	return matrix
}

func createDefaultTwoDimensionalSlice() [][]int {
	twoDimensionalSlice := [][]int{
			{1, 2, 3, 1},
			{4, 5, 6, 4},
			{7, 8, 9, 7},
			{7, 8, 9, 7},
	}

	fmt.Println("Default 2D slice was created:")
	fmt.Println()
	fmt.Printf("%v", twoDimensionalSlice)
	fmt.Println()

	return twoDimensionalSlice
}
