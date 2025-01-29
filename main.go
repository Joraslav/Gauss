package main

import (
	gauss "Gauss/Gauss"
	"fmt"
)

func main() {
	// Пример системы уравнений:
	// 2x + y - z = 8
	// -3x - y + 2z = -11
	// -2x + y + 2z = -3
	matrix := gauss.NewMatrix(3, 4)
	matrix.Data = [][]float64{
		{2, 1, -1, 8},
		{-3, -1, 2, -11},
		{-2, 1, 2, -3},
	}
	solution := gauss.SolveGauss(matrix)
	fmt.Println("Решение системы:")
	for i, val := range solution {
		fmt.Printf("x%d = %.2f\n", i+1, val)
	}
}
