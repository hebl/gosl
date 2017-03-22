package linalg

import (
	"math"

	"github.com/hebl/gosl/errors"
	"github.com/hebl/gosl/matrix"
	"github.com/hebl/gosl/permutation"
)

/**
LU_decomp
*/

//LU LU decomp
func LU(A *matrix.Matrix) (*permutation.Permutation, error) {
	size1, size2 := A.Shape()
	if size1 != size2 {
		return nil, errors.ErrBADLEN
	}

	N := size1

	p := permutation.New(N)

	for j := 0; j < N-1; j++ {
		//var ajj float64
		max := math.Abs(A.At(j, j))
		ipivot := j

		for i := j + 1; i < N; i++ {
			aij := math.Abs(A.At(i, j))

			if aij > max {
				max = aij
				ipivot = i
			}
		}

		if ipivot != j {
			A.SwapRows(j, ipivot)
			p.Swap(j, ipivot)
		}
		ajj := A.At(j, j)

		if ajj != 0.0 {
			for i := j + 1; i < N; i++ {
				aij := A.At(i, j) / ajj
				A.Set(i, j, aij)

				for k := j + 1; k < N; k++ {
					aik := A.At(i, k)
					ajk := A.At(j, k)
					A.Set(i, k, aik-aij*ajk)
				}
			}
		}
	}

	return p, nil
}
