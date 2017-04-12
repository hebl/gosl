// Copyright 2017 HE Boliang. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vector

import (
	"fmt"

	"math"

	"github.com/hebl/gosl/errors"
)

//Equal if vectors v and a are equal?
func (v *Vector) Equal(a *Vector) bool {
	if a.size != v.size {
		panic(fmt.Sprintf("vector size %d <> vector a size %d", v.size, a.size))
	}
	return false
}

//Min returns the minimum value
func (v *Vector) Min() float64 {
	min := v.data[0]
	for i := 1; i < v.size; i++ {
		if v.data[i] < min {
			min = v.data[i]
		}
	}
	return min
}

//MinIndex returns the minimum value
func (v *Vector) MinIndex() int {
	min := v.data[0]
	idx := 0
	for i := 1; i < v.size; i++ {
		if v.data[i] < min {
			min = v.data[i]
			idx = i
		}
	}
	return idx
}

//AMin returns the minimum abs value
func (v *Vector) AMin() float64 {
	min := math.Abs(v.data[0])
	for i := 1; i < v.size; i++ {
		if math.Abs(v.data[i]) < min {
			min = math.Abs(v.data[i])
		}
	}
	return min
}

//AMinIndex returns the minimum abs value
func (v *Vector) AMinIndex() int {
	min := math.Abs(v.data[0])
	idx := 0
	for i := 1; i < v.size; i++ {
		if math.Abs(v.data[i]) < min {
			min = math.Abs(v.data[i])
			idx = i
		}
	}
	return idx
}

//Max returns the maximum value
func (v *Vector) Max() float64 {
	max := v.data[0]
	for i := 1; i < v.size; i++ {
		if v.data[i] > max {
			max = v.data[i]
		}
	}
	return max
}

//MaxIndex returns the maximum value
func (v *Vector) MaxIndex() int {
	max := v.data[0]
	idx := 0
	for i := 1; i < v.size; i++ {
		if v.data[i] > max {
			max = v.data[i]
			idx = i
		}
	}
	return idx
}

//AMax returns the maximum abs value
func (v *Vector) AMax() float64 {
	max := math.Abs(v.data[0])
	for i := 1; i < v.size; i++ {
		if math.Abs(v.data[i]) > max {
			max = math.Abs(v.data[i])
		}
	}
	return max
}

//AMaxIndex returns the maximum abs value
func (v *Vector) AMaxIndex() int {
	max := math.Abs(v.data[0])
	idx := 0
	for i := 1; i < v.size; i++ {
		if math.Abs(v.data[i]) > max {
			max = math.Abs(v.data[i])
			idx = i
		}
	}
	return idx
}

//Add add
func (v *Vector) Add(a *Vector) error {
	if a.size != v.size {
		return errors.ErrBADLEN
	}
	for i := 0; i < v.size; i++ {
		v.data[i] += a.data[i]
	}
	return nil
}

//Sub sub
func (v *Vector) Sub(a *Vector) error {
	if a.size != v.size {
		return errors.ErrBADLEN
	}
	for i := 0; i < v.size; i++ {
		v.data[i] -= a.data[i]
	}
	return nil
}

//Mul multiplies the elements of vector v by the elements of vector a
func (v *Vector) Mul(a *Vector) error {
	if a.size != v.size {
		return errors.ErrBADLEN
	}
	for i := 0; i < v.size; i++ {
		v.data[i] = v.data[i] * a.data[i]
	}
	return nil
}

//Div divides the elements of vector v by the elements of vector a
func (v *Vector) Div(a *Vector) error {
	if a.size != v.size {
		return errors.ErrBADLEN
	}
	for i := 0; i < v.size; i++ {
		v.data[i] = v.data[i] / a.data[i]
	}
	return nil
}

//Scale multiplies the elements of vector v by the constant factor x.
func (v *Vector) Scale(x float64) {
	for i := 0; i < v.size; i++ {
		v.data[i] = v.data[i] * x
	}
}

//AddConstant adds the constant value x to the elements of the vector v
func (v *Vector) AddConstant(x float64) {
	for i := 0; i < v.size; i++ {
		v.data[i] = v.data[i] + x
	}
}

//Dot dot
func (v *Vector) Dot(a *Vector) (float64, error) {
	if a.size != v.size {
		return 0, errors.ErrBADLEN
	}
	var sum float64
	sum = 0
	for i := range v.data {
		sum += v.data[i] * a.data[i]
	}
	return sum, nil
}
