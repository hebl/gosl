package vector

import "testing"

func TestAdd(t *testing.T) {
	v := NewVector(10, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	a := NewVector(10, []float64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})

	v.Add(a)
}

func TestSub(t *testing.T) {
	v := NewVector(10, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	a := NewVector(10, []float64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})

	v.Add(a)
}

func TestMul(t *testing.T) {
	v := NewVector(10, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	a := NewVector(10, []float64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})

	v.Add(a)
}

func TestDiv(t *testing.T) {
	v := NewVector(10, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	a := NewVector(10, []float64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})

	v.Add(a)
}

func TestScale(t *testing.T) {
	v := NewVector(10, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	v.Scale(10.213)
}

func TestAddConstant(t *testing.T) {
	v := NewVector(10, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	v.AddConstant(10.213)
}
