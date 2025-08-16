package goarima

import (
	"errors"
	"math"
)

// Solve solves the linear system Ax = b using Gaussian elimination.
// A is an n x n matrix, b is a vector of length n.
// It returns the solution vector x or an error if the matrix is singular.
func Solve(A [][]float64, b []float64) ([]float64, error) {
	n := len(b)
	// Build augmented matrix
	aug := make([][]float64, n)
	for i := 0; i < n; i++ {
		aug[i] = make([]float64, n+1)
		copy(aug[i], A[i])
		aug[i][n] = b[i]
	}

	// Forward elimination
	for i := 0; i < n; i++ {
		// Pivot
		maxRow := i
		for r := i + 1; r < n; r++ {
			if math.Abs(aug[r][i]) > math.Abs(aug[maxRow][i]) {
				maxRow = r
			}
		}
		if math.Abs(aug[maxRow][i]) < 1e-12 {
			return nil, errors.New("singular matrix")
		}
		aug[i], aug[maxRow] = aug[maxRow], aug[i]

		// Normalize pivot row
		piv := aug[i][i]
		for c := i; c <= n; c++ {
			aug[i][c] /= piv
		}

		// Eliminate below
		for r := i + 1; r < n; r++ {
			factor := aug[r][i]
			for c := i; c <= n; c++ {
				aug[r][c] -= factor * aug[i][c]
			}
		}
	}

	// Back substitution
	x := make([]float64, n)
	for i := n - 1; i >= 0; i-- {
		x[i] = aug[i][n]
		for j := i + 1; j < n; j++ {
			x[i] -= aug[i][j] * x[j]
		}
	}
	return x, nil
}
