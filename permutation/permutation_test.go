package permutation

import "testing"

func TestPerm(t *testing.T) {
	p := New(10)
	p.Swap(3, 7)
	p.Mprint()
}
