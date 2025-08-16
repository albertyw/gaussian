package main

import (
	"fmt"

	gaussian "github.com/albertyw/go-gaussian-elimination"
)

func main() {
	a := [][]float64{
		{3, 2, -1},
		{2, -2, 4},
		{-1, 0.5, -1},
	}
	b := []float64{1, -2, 0}
	solution, err := gaussian.Solve(a, b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(solution)
}
