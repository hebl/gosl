package matrix

import "github.com/hebl/gosl/errors"

//Product The product of two matrices
func Product(A, B *Matrix) (*Matrix, error) {
	if A.size2 != B.size1 {
		return nil, errors.ErrBADLEN
	}

	C := Zeros(A.size1, B.size2)

	for i := 0; i < A.size1; i++ {
		for j := 0; j < B.size2; j++ {
			v := 0.0
			for k := 0; k < A.size2; k++ {
				v += A.At(i, k) * B.At(k, j)
			}
			C.Set(i, j, v)
		}
	}

	return C, nil
}
