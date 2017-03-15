package vector

import (
	"math"
)

//Norm1 norm
func (v *Vector) Norm1() float64 {
	var sum float64
	for _, v := range v.data {
		sum += math.Abs(v)
	}
	return sum
}

//Norm2 norm
func (v *Vector) Norm2() float64 {
	var sum float64
	for _, v := range v.data {
		sum += v * v
	}
	return math.Sqrt(sum)
}

//NormInf norm
func (v *Vector) NormInf() float64 {
	return 0
}

//Mean Mean of vector
func (v *Vector) Mean() float64 {
	var sum float64
	for _, v := range v.data {
		sum += v
	}
	return sum / float64(v.size)
}

//Median Median of vector
//TODO
func (v *Vector) Median() float64 {
	return 0
}

//Variance variance of vector
func (v *Vector) Variance() float64 {
	var variance float64
	mean := v.Mean()
	for i := range v.data {
		delta := v.data[i] - mean
		variance += (delta*delta - variance) / float64(i+1)
	}
	return variance
}

//Std Std of vector
func (v *Vector) Std() float64 {
	return math.Sqrt(v.Variance())
}
