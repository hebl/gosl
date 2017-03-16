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
