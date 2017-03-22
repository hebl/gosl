package permutation

import "github.com/hebl/gosl/errors"
import "fmt"

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

//Print print
func (p *Permutation) Print() {
	fmt.Println(p.data)
}
