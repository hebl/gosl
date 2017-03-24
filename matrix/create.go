package matrix

import (
	"fmt"
	"math/rand"
	"strings"
)

//Ones matrix filled with 1
func Ones(n, m int) *Matrix {
	mat := New(n, m)
	//data := make([]float64, n*m)
	for i := range mat.data {
		mat.data[i] = 1.0
	}
	return mat
}

//Zeros matrix filled with 0
func Zeros(n, m int) *Matrix {
	return New(n, m)
}

//Identity Return the identity square matrix
func Identity(n int) *Matrix {
	mat := NewS(n)
	for i := 0; i < n; i++ {
		mat.data[i*mat.tda+i] = 1.0
	}
	return mat
}

//Rand Rand float64 [0.0,1.0)
func Rand(n, m int) *Matrix {
	mat := New(n, m)
	for i := 0; i < n*m; i++ {
		mat.data[i] = rand.Float64()
	}
	return mat
}

//----------------------------------

//LoadFromStr load from str
// load from [1 2 3; 4 5 6; 7 8 9]
func LoadFromStr(str string) *Matrix {
	str = strings.TrimSpace(str)
	str = strings.TrimPrefix(str, "[")
	str = strings.TrimSuffix(str, "]")
	rows := strings.Split(str, ";")
	size1 := len(rows)
	size2 := 0
	for i := 0; i < size1; i++ {
		cs := strings.Fields(rows[i])
		if size2 < len(cs) {
			size2 = len(cs)
		}
	}
	data := make([]float64, size1*size2)
	for i := 0; i < size1; i++ {

		rd := make([]interface{}, size2)
		for j := 0; j < size2; j++ {
			rd[j] = &data[i*size2+j]

		}
		fmt.Sscan(rows[i], rd...)
	}

	return &Matrix{
		size1: size1,
		size2: size2,
		tda:   size2,
		data:  data,
	}
}

//LoadFromArray  load from a array
func LoadFromArray(size1, size2 int, data []float64) *Matrix {
	tsize := size1 * size2

	m := &Matrix{
		size1: size1,
		size2: size2,
		tda:   size2,
	}

	if len(data) > tsize {
		data = data[:tsize]
	}

	if len(data) < tsize {
		data = append(data, make([]float64, tsize-len(data))...)
	}

	m.data = data

	return m
}
