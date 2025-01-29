package main

import (
	gauss "Gauss/Gauss"
	read "Gauss/Read"
	"fmt"
)

func main() {
	matrix, err := read.ReadMatrixFromFile("matrix.txt")
	if err != nil {
		fmt.Printf("Error reading matrix from file: %v\n", err)
		return
	}
	solution := gauss.SolveGauss(matrix)
	fmt.Println("Решение системы:")
	for i, val := range solution {
		fmt.Printf("x%d = %.2f\n", i+1, val)
	}
}
