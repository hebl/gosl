package linalg

import (
	"testing"

	"github.com/hebl/gosl/matrix"
)

func TestLU(t *testing.T) {
	A := matrix.NewS(3)
	A.Set(0, 0, 2)
	A.Set(0, 1, 1)
	A.Set(0, 2, 5)
	A.Set(1, 0, 4)
	A.Set(1, 1, 4)
	A.Set(1, 2, -4)
	A.Set(2, 0, 1)
	A.Set(2, 1, 3)
	A.Set(2, 2, 1)

	B := A.Copy()

	A.Print()
	//t.Log("\n")

	p, _ := LU(A)

	A.Print()
	p.Print()

	m := p.Matrix()
	m.Print()

	L, U, P := LUP(B)
	//t.Log("\n")
	L.Print()
	//t.Log("\n")
	P.Print()
	//t.Log("\n")
	U.Print()

}
