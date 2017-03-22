package matrix

import (
	"fmt"
)

//Matrix matrix
type Matrix struct {
	size1 int
	size2 int
	tda   int
	data  []float64
}

//New create a matrix
func New(i, j int) *Matrix {

	m := &Matrix{
		size1: i,
		size2: j,
	}
	m.tda = j
	m.data = make([]float64, i*j)

	return m
}

//NewS create a square matrix
func NewS(n int) *Matrix {

	m := &Matrix{
		size1: n,
		size2: n,
	}
	m.tda = n
	m.data = make([]float64, n*n)

	return m
}

//At Get (i,j) value
func (m *Matrix) At(i, j int) float64 {
	return m.data[i*m.tda+j]
}

//Get Get (i,j) value
func (m *Matrix) Get(i, j int) float64 {
	return m.At(i, j)
}

//Set Set (i,j) value for v
func (m *Matrix) Set(i, j int, v float64) {
	m.data[i*m.tda+j] = v
}

//Shape matrix shape
func (m *Matrix) Shape() (int, int) {
	return m.size1, m.size2
}

//Print print matrix
func (m *Matrix) Print() {
	for i := 0; i < m.size1; i++ {
		for j := 0; j < m.size2-1; j++ {
			fmt.Printf("%.04f  ", m.At(i, j))
		}
		fmt.Printf("%.04f\n", m.At(i, m.size2-1))
	}
}
