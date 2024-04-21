package main

// import "fmt"

// func main() {
// 	// declare an int array of length 5
// 	//element type and length are part of the arrays TYPE
// 	var a [5]int
// 	// prints empy: [0 0 0 0 0], arrays are zero valued by default
// 	fmt.Println("empy:", a)
// 	var b [5]bool
// 	// prints empy: [false false false false false], since arrays are zero valued by default
// 	fmt.Println("empy:", b)

// 	//values can be assigned and retrieved with square bracket index notation
// 	a[0] = 12
// 	a[4] = 100

// 	//built-in for determining length of an array
// 	fmt.Println("length of array a:", len(a))

// 	//single line array declaration/initializing
// 	//uses squarebrackets to declare array length
// 	//curly brace to declare values. Undeclared values will be initialized as appropriate 0 value for data type
// 	var g = [5]int{1, 2, 3, 4}
// 	fmt.Println(g)

// 	// 2d array initialization and iterative traversal
// 	// print returns stringified data structure
// 	var twoDimensionalArray [2][3]int
// 	for i := 0; i < 2; i++ {
// 		for j := 0; j < 3; j++ {
// 			twoDimensionalArray[i][j] = i + j
// 		}
// 	}
// 	//prints 2d arr: [[0 1 2] [1 2 3]]
// 	fmt.Println("2d arr:", twoDimensionalArray)
// }
//
