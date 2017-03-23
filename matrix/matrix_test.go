package matrix

import (
	"testing"
)

func TestMatrix1(t *testing.T) {
	A := New(10, 10)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			A.Set(i, j, float64(i*j)+0.1)
		}
	}
	A.Print()
}

func TestOnes(t *testing.T) {
	A := Identity(10)
	A.Print()
}

func TestProduct(t *testing.T) {
	A := New(3, 2)
	B := New(2, 3)

	A.Set(0, 0, 3)
	A.Set(0, 1, -2)
	A.Set(1, 0, 2)
	A.Set(1, 1, 4)
	A.Set(2, 0, 1)
	A.Set(2, 1, -3)

	A.Print()

	B.Set(0, 0, -2)
	B.Set(0, 1, 1)
	B.Set(0, 2, 3)
	B.Set(1, 0, 4)
	B.Set(1, 1, 1)
	B.Set(1, 2, 6)

	B.Print()

	C, _ := Product(A, B)
	C.Print()
}

func TestTransport(t *testing.T) {
	A := New(3, 2)
	A.Set(0, 0, 3)
	A.Set(0, 1, -2)
	A.Set(1, 0, 2)
	A.Set(1, 1, 4)
	A.Set(2, 0, 1)
	A.Set(2, 1, -3)

	A.Print()

	B := A.T()

	B.Print()
}

func TestCopy(t *testing.T) {
	A := New(3, 2)
	A.Set(0, 0, 3)
	A.Set(0, 1, -2)
	A.Set(1, 0, 2)
	A.Set(1, 1, 4)
	A.Set(2, 0, 1)
	A.Set(2, 1, -3)

	A.Print()

	B := A.Copy()

	B.Print()
}

func TestSwapRows(t *testing.T) {

	A := LoadFromArray(3, 2, []float64{3, -2, 2, 4, 1, -3})

	A.Print()

	A.SwapRows(1, 2)
	A.Print()

	A.SwapCols(0, 1)
	A.Print()
}

func TestLoadFromArray(t *testing.T) {
	A := LoadFromArray(2, 2, []float64{1, 2, 3, 4, 5})
	A.Print()
}

func TestLoadFromStr(t *testing.T) {
	str := "[1 2 3   9; 4 5    7;9]"
	A := LoadFromStr(str)
	A.Print()
}
