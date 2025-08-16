# Gaussian Elimination in Go

[![Build Status](https://drone.albertyw.com/api/badges/albertyw/go-gaussian-elimination/status.svg)](https://drone.albertyw.com/albertyw/go-gaussian-elimination)

Solve linear equations using Gaussian Elimination.

See [Wikipedia](https://en.wikipedia.org/wiki/Gaussian_elimination) for more information about Gaussian elimination.

## Usage

For three linear equations:

```
3a  + 2b   - 1c = 1
2a  - 2b   + 4c = -2
-1a + 0.5b - 1c = 0
```

This can be converted into a coefficient matrix `A`:

```
| 3   2    -1 |
| 2  -2     4 |
| -1  0.5  -1 |
```

and a right-hand side vector `B`:

```
| 1 -2 0 |
```

Solving the linear equations would be done with:

```go
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
    fmt.Println(solution) // [1, -2, -2]
}
```
