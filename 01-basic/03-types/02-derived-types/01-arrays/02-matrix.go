package main

import "fmt"

func main()  {
	mymatrix := [2][3]int{
		{3, 4, 5},
		{2, 3, 6},
	}
	fmt.Println(mymatrix)

	// Indexing a Matrix:
	// 	matrix[rowIndex][columnIndex]
	fmt.Println("Element row 20, col 2", mymatrix[0][1])
	fmt.Println("Element in row 1, col 2", mymatrix[1][2])

	//print each element
	// looping for rows
	// Loop through rows
	for i := 0; i < len(mymatrix); i++ {
		// Loop through columns
		for j := 0; j < len(mymatrix[i]); j++ {
			// Use fmt.Printf() for formatted output
			fmt.Printf("Element at [%d][%d]: %d\n", i, j, mymatrix[i][j])
		}
	}
}
