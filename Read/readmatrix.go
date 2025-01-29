package read

import (
	gauss "Gauss/Gauss"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadMatrixFromFile(filename string) (*gauss.Matrix, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	matrix := &gauss.Matrix{}

	// Чтение строк по одной
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue // пропускаем пустые строки
		}

		switch parts[0] {
		case "Rows:":
			rows, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, fmt.Errorf("неверное число строк: %v", err)
			}
			matrix.Rows = rows
		case "Cols:":
			cols, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, fmt.Errorf("невероное число столбцов: %v", err)
			}
			matrix.Cols = cols
		case "Data:":
			// Пропускаем заголовок "Data:"
			if matrix.Rows == 0 || matrix.Cols == 0 {
				return nil, fmt.Errorf("это Rows и Cols должны быть до Data")
			}
			matrix.Data = make([][]float64, matrix.Rows)
			for i := 0; i < matrix.Rows; i++ {
				if !scanner.Scan() {
					return nil, fmt.Errorf("нет данных в строках")
				}
				line = scanner.Text()
				numbers := strings.Fields(line)
				if len(numbers) != matrix.Cols {
					return nil, fmt.Errorf("ожидалось %d столбцов, получено %d", matrix.Cols, len(numbers))
				}
				matrix.Data[i] = make([]float64, matrix.Cols)
				for j, numStr := range numbers {
					num, err := strconv.ParseFloat(numStr, 64)
					if err != nil {
						return nil, fmt.Errorf("невероное float значение в строке %d, столбце %d: %v", i, j, err)
					}
					matrix.Data[i][j] = num
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if matrix.Rows == 0 || matrix.Cols == 0 || len(matrix.Data) != matrix.Rows {
		return nil, fmt.Errorf("incomplete matrix data")
	}

	return matrix, nil
}
