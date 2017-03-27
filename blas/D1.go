package blas

import (
	"github.com/hebl/gosl/errors"
	"github.com/hebl/gosl/vector"
)

/**
Double
Level 1
*/

//DDot scalar product x^T y for the vectors x and y
func DDot(x, y *vector.Vector) (float64, error) {
	if x.Size() != y.Size() {
		return 0.0, errors.ErrINVIDX
	}

	var ddot float64
	for i := 0; i < x.Size(); i++ {
		ddot += x.At(i) * y.At(i)
	}

	return ddot, nil
}

//DCopy copy the elements of the vector x into the vector y
func DCopy(x, y *vector.Vector) error {
	if x.Size() == y.Size() {
		copy(*y.GetData(), *x.GetData())
		return nil
	}

	return errors.ErrINVIDX
}

//DSwap exchange the elements of the vectors x and y
func DSwap(x, y *vector.Vector) error {
	if x.Size() == y.Size() {
		for i := 0; i < x.Size(); i++ {
			xi, yi := x.At(i), y.At(i)
			x.Set(i, yi)
			y.Set(i, xi)
		}
		return nil
	}

	return errors.ErrINVIDX
}
