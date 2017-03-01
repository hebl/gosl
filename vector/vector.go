package vector

import (
	"fmt"
)

//Vector vector struct
type Vector struct {
	size int
	data []float64
}

//NewVector new vector
func NewVector(n int, data []float64) *Vector {
	if len(data) != n && data != nil {
		panic("")
	}

	if data == nil {
		data = make([]float64, n)
	}
	v := &Vector{
		size: len(data),
		data: data,
	}
	return v
}

//Len vector length
func (v *Vector) Len() int {
	return v.size
}

func (v *Vector) checkidx(i int) {
	if i < 0 || i >= v.size {
		panic(fmt.Sprintf("%d out of vector range 0-%d", i, v.size))
	}
}

//At Get index value
func (v *Vector) At(i int) float64 {
	v.checkidx(i)
	return v.data[i]

}

//Set set index value
func (v *Vector) Set(i int, val float64) {
	if i < 0 || i >= v.size {
		panic("")
	}

}
