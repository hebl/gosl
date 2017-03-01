package vector

import (
	"fmt"
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

//Add add
func (v *Vector) Add(a *Vector) {
	if a.size != v.size {
		panic(fmt.Sprintf("vector size %d <> vector a size %d", v.size, a.size))
	}
	for i := 0; i < v.size; i++ {
		v.data[i] += a.data[i]
	}
}

//Sub sub
func (v *Vector) Sub(a *Vector) {
	if a.size != v.size {
		panic(fmt.Sprintf("vector size %d <> vector a size %d", v.size, a.size))
	}
	for i := 0; i < v.size; i++ {
		v.data[i] -= a.data[i]
	}
}

//Mul multiplies the elements of vector v by the elements of vector a
func (v *Vector) Mul(a *Vector) {
	if a.size != v.size {
		panic(fmt.Sprintf("vector size %d <> vector1 size %d", v.size, a.size))
	}
	for i := 0; i < v.size; i++ {
		v.data[i] = v.data[i] * a.data[i]
	}
}

//Div divides the elements of vector v by the elements of vector a
func (v *Vector) Div(a *Vector) {
	if a.size != v.size {
		panic(fmt.Sprintf("vector size %d <> vector1 size %d", v.size, a.size))
	}
	for i := 0; i < v.size; i++ {
		v.data[i] = v.data[i] / a.data[i]
	}
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
