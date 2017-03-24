package permutation

import (
	"fmt"

	"github.com/hebl/gosl/errors"
	"github.com/hebl/gosl/matrix"
)

//Permutation permutations
type Permutation struct {
	size int
	data []int
}

//New new Permutation
func New(n int) *Permutation {
	p := &Permutation{
		size: n,
	}
	p.data = make([]int, n)
	for i := 0; i < n; i++ {
		p.data[i] = i
	}
	return p
}

//Size Permutation size
func (p *Permutation) Size() int { return p.size }

//Swap swap i, j
func (p *Permutation) Swap(i, j int) error {
	if i >= p.size || j >= p.size {
		return errors.ErrINVAL
	}

	p.data[i], p.data[j] = p.data[j], p.data[i]

	return nil
}

//Matrix matrix view
func (p *Permutation) Matrix() *matrix.Matrix {
	m := matrix.Zeros(p.size, p.size)
	for i := 0; i < p.size; i++ {
		m.Set(i, p.data[i], 1)
	}
	return m
}

//Print print
func (p *Permutation) Print() {
	fmt.Println(p.data)
}

//Mprint print
func (p *Permutation) Mprint() {
	fmt.Print("\n")
	for i := 0; i < p.size; i++ {
		fmt.Print("   ")
		for j := 0; j < p.size; j++ {
			if j == p.data[i] {
				fmt.Printf("%4d", 1)
			} else {
				fmt.Printf("%4d", 0)
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}
