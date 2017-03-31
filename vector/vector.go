// Copyright 2017 HE Boliang. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vector

import (
	"math/rand"

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
func New(n int) *Vector {

	data := make([]float64, n)

	v := &Vector{
		size:   n,
		stride: 1,
		data:   data,
	}
	return v
}

//Size vector length
func (v *Vector) Size() int {
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

//GetData Get vector data
func (v *Vector) GetData() *[]float64 {
	return &v.data
}

//Copy vector copy
func (v *Vector) Copy() *Vector {
	v2 := New(v.size)
	copy(v2.data, v.data)
	return v2
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

//Rand rand value
func Rand(n int) *Vector {
	v := New(n)
	for i := range v.data {
		v.data[i] = rand.Float64()
	}
	return v
}
