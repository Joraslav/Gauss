package gauss

import (
	"math"
	"sync"
)

// Matrix представляет расширенную матрицу системы уравнений
type Matrix struct {
	Rows int
	Cols int
	Data [][]float64
}

// NewMatrix создает новую матрицу заданного размера
func NewMatrix(rows, cols int) *Matrix {
	data := make([][]float64, rows)
	for i := range data {
		data[i] = make([]float64, cols)
	}
	return &Matrix{Rows: rows, Cols: cols, Data: data}
}

// SwapRows меняет местами две строки матрицы
func (m *Matrix) SwapRows(i, j int) {
	m.Data[i], m.Data[j] = m.Data[j], m.Data[i]
}

// DirectPass реализует прямой ход метода Гаусса с распараллеливанием
func (m *Matrix) DirectPass() {
	for i := 0; i < m.Rows; i++ {
		// 1. Выбор ведущего элемента (pivot) для улучшения стабильности
		pivotRow := i
		maxVal := math.Abs(m.Data[i][i])
		for j := i + 1; j < m.Rows; j++ {
			if currentVal := math.Abs(m.Data[j][i]); currentVal >
				maxVal {
				maxVal = currentVal
				pivotRow = j
			}
		}
		// Перестановка строк, если найден больший ведущий элемент
		if pivotRow != i {
			m.SwapRows(i, pivotRow)
		}
		// 2. Нормировка текущей строки
		divisor := m.Data[i][i]
		if divisor == 0 {
			panic("Матрица вырождена, система не имеет единственного решения")
		}
		for col := i; col < m.Cols; col++ {
			m.Data[i][col] /= divisor
		}
		// 3. Распараллеленное вычитание текущей строки из последующих
		var wg sync.WaitGroup
		for j := i + 1; j < m.Rows; j++ {
			wg.Add(1)
			go func(j int) {
				defer wg.Done()
				multiplier := m.Data[j][i]
				for col := i; col < m.Cols; col++ {
					m.Data[j][col] -= multiplier * m.Data[i][col]
				}
			}(j)
		}
		wg.Wait() // Ожидаем завершения всех горутин
	}
}

// BackSubstitution реализует обратную подстановку
func (m *Matrix) BackSubstitution() []float64 {
	solution := make([]float64, m.Rows)
	for i := m.Rows - 1; i >= 0; i-- {
		sum := 0.0
		for j := i + 1; j < m.Rows; j++ {
			sum += m.Data[i][j] * solution[j]
		}
		solution[i] = m.Data[i][m.Cols-1] - sum
	}
	return solution
}

// SolveGauss решает систему уравнений методом Гаусса
func SolveGauss(m *Matrix) []float64 {
	m.DirectPass()
	return m.BackSubstitution()
}
