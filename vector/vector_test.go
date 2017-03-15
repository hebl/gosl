package vector

import "testing"

var (
	v, _ = New(10, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	a, _ = New(10, []float64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
)

func TestAdd(t *testing.T) {
	v.Add(a)
}

func TestSub(t *testing.T) {

	v.Sub(a)
}

func TestMul(t *testing.T) {

	v.Mul(a)
}

func TestDiv(t *testing.T) {

	v.Div(a)
}

func TestScale(t *testing.T) {

	v.Scale(10.213)
}

func TestAddConstant(t *testing.T) {

	v.AddConstant(10.213)
}

func TestNorm(t *testing.T) {

	n1 := v.Norm1()
	n2 := v.Norm2()

	t.Log(n1, n2)
}
