// Copyright 2017 HE Boliang. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Variance, Sd

package statistics

import "math"

// Variance returns the estimated, or sample, variance of data, a dataset of length n
// with stride stride. The estimated variance is denoted by σˆ2
func Variance(data []float64, stride int, size int) float64 {
	mean := Mean(data, stride, size)
	return VarianceM(data, stride, size, mean)
}

// VarianceM returns the estimated, or sample, variance of data, a dataset of length n
// with stride stride. The estimated variance is denoted by σˆ2
func VarianceM(data []float64, stride int, size int, mean float64) float64 {
	var σ2 float64

	for i := 0; i < size; i++ {
		delta := (data[i*stride] - mean)
		σ2 += (delta*delta - σ2) / float64(i+1)
	}
	return σ2
}

// Sd The standard deviation is defined as the square root of the variance.
// These functions return the square root of the corresponding variance functions above.
func Sd(data []float64, stride int, size int) float64 {
	mean := Mean(data, stride, size)
	return SdM(data, stride, size, mean)
}

// SdM The standard deviation is defined as the square root of the variance.
// These functions return the square root of the corresponding variance functions above.
func SdM(data []float64, stride int, size int, mean float64) float64 {

	σ2 := VarianceM(data, stride, size, mean)
	return math.Sqrt(σ2 * float64(size) / float64(size-1))
}

// Tss return the total sum of squares (TSS) of data about the mean
func Tss(data []float64, stride int, size int) float64 {
	mean := Mean(data, stride, size)
	return TssM(data, stride, size, mean)
}

// TssM return the total sum of squares (TSS) of data about the mean
func TssM(data []float64, stride int, size int, mean float64) float64 {
	var tss float64
	for i := 0; i < size; i++ {
		delta := data[i*stride] - mean
		tss += delta * delta
	}
	return tss
}
