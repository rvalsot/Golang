package main

import "fmt"

func main() {
	fmt.Println("Multidimensional Slice:")

	records := make([][]string, 0)

	// student 1
	student1 := make([]string, 4)
	student1[0] = "Pijamas"
	student1[1] = "Juan"
	student1[2] = "Math"
	student1[3] = "100"

	records = append(records, student1)

	// student 2
	student2 := make([]string, 4)
	student2[0] = "Man"
	student2[1] = "Bear"
	student2[2] = "Being Russian"
	student2[3] = "100"

	records = append(records, student2)

	//
	// fmt.Println(records[1])
	matrixSlice := make([][]int, 0)
	for i := 0; i < 3; i++ {
		yAxis := make([]int, 0)
		for j := 0; j < 3; j++ {
			yAxis = append(yAxis, j)
		}
		matrixSlice = append(matrixSlice, yAxis)
	}
	fmt.Println(matrixSlice)
}
