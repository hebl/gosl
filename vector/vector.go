package vector

import (
	"github.com/hebl/gosl/errors"
)

//Vector32 vector struct
type Vector32 struct {
	size   int
	stride int
	data   []float32
}

//Vector vector struct
type Vector struct {
	size   int
	stride int
	data   []float64
}

//New new vector
func New(n int, data []float64) (*Vector, error) {
	if len(data) != n && data != nil {
		return nil, errors.RETFAILURE
	}

	if data == nil {
		data = make([]float64, n)
	}
	v := &Vector{
		size:   len(data),
		stride: 1,
		data:   data,
	}
	return v, nil
}

//Len vector length
func (v *Vector) Len() int {
	return v.size
}

func (v *Vector) checkidx(i int) error {
	if i < 0 || i >= v.size {
		return errors.ErrINVIDX
	}
	return nil
}

//At Get index value
func (v *Vector) At(i int) float64 {
	v.checkidx(i)
	return v.data[i]

}

//Set set index value
func (v *Vector) Set(i int, val float64) {
	v.checkidx(i)
	v.data[i] = val
}

//========================================

//Ones vector filled with 1
func Ones(n int) (*Vector, error) {
	if n < 0 {
		return nil, errors.RETFAILURE
	}

	data := make([]float64, n)

	for i := range data {
		data[i] = 1
	}

	v := &Vector{
		size:   n,
		stride: 1,
		data:   data,
	}

	return v, nil
}

//Zeros vector filled with 0
func Zeros(n int) (*Vector, error) {
	if n < 0 {
		return nil, errors.RETFAILURE
	}

	v := &Vector{
		size:   n,
		stride: 1,
		data:   make([]float64, n),
	}

	return v, nil
}
