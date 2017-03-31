// Copyright 2017 HE Boliang. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Mean Median Min Max

package statistics

// Mean arithmetic mean of data, a dataset of length n with stride stride.
func Mean(data []float64, stride int, size int) float64 {
	var μ float64
	for i := 0; i < size; i++ {
		μ += (data[i*stride] - μ) / float64(i+1)
	}
	return μ
}

// Median returns the median value of sorted data, a dataset of length n with stride stride.
// data must be sorted
func Median(data []float64, stride int, size int) float64 {
	var median float64
	lhs := (size - 1) / 2
	rhs := size / 2

	if size == 0 {
		return 0.0
	}

	if lhs == rhs {
		median = data[lhs*stride]
	} else {
		median = (data[lhs*stride] + data[rhs*stride]) / 2.0
	}

	return median
}

// Min returns the minimum value in data
func Min(data []float64, stride int, size int) float64 {
	var min float64
	for i := 0; i < size; i++ {
		if data[i*stride] < min {
			min = data[i*stride]
		}
	}
	return min
}

// Max returns the maximum value in data
func Max(data []float64, stride int, size int) float64 {
	var max float64
	for i := 0; i < size; i++ {
		if data[i*stride] > max {
			max = data[i*stride]
		}
	}
	return max
}

// MinMax finds both the minimum and maximum
func MinMax(data []float64, stride int, size int) (float64, float64) {
	return Min(data, stride, size), Max(data, stride, size)
}
