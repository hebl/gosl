package matrix

import (
	"github.com/hebl/gosl/errors"
	"github.com/hebl/gosl/vector"
)

//Row Get row r
func (m *Matrix) Row(r int) *vector.Vector {
	v, err := vector.Zeros(m.size2)
	if err != nil {
		return nil
	}
	offset := r * m.size2
	for i := 0; i < m.size2; i++ {
		v.Set(i, m.data[offset+i])
	}

	return v
}

//Col Get Col c
func (m *Matrix) Col(c int) *vector.Vector {
	v, err := vector.Zeros(m.size1)
	if err != nil {
		return nil
	}
	offset := c
	for i := 0; i < m.size1; i++ {
		v.Set(i, m.data[i*m.tda+offset])
	}

	return v
}

//Copy matrix hard copy
func (m *Matrix) Copy() *Matrix {
	m2 := New(m.size1, m.size2)
	for i := range m.data {
		m2.data[i] = m.data[i]
	}
	return m2
}

//T matrix transpose
func (m *Matrix) T() *Matrix {
	m2 := New(m.size2, m.size1)

	for i := 0; i < m.size1; i++ {
		for j := 0; j < m.size2; j++ {
			m2.Set(j, i, m.At(i, j))
		}
	}

	return m2
}

//SwapRows swap row i, j
func (m *Matrix) SwapRows(i, j int) error {
	if i >= m.size1 || j >= m.size1 {
		return errors.ErrINVAL
	}

	for k := 0; k < m.size2; k++ {
		m.data[i*m.tda+k], m.data[j*m.tda+k] = m.data[j*m.tda+k], m.data[i*m.tda+k]
	}

	return nil
}

//SwapCols swap col i, j
func (m *Matrix) SwapCols(i, j int) error {
	if i >= m.size2 || j >= m.size2 {
		return errors.ErrINVAL
	}

	for k := 0; k < m.size1; k++ {
		m.data[k*m.tda+i], m.data[k*m.tda+j] = m.data[k*m.tda+j], m.data[k*m.tda+i]
	}

	return nil
}

//Singular is singular?
func (m *Matrix) Singular() int {
	return 0
}

//Norm1 norm
//TODO
func (m *Matrix) Norm1() float64 {
	return 0
}

//Norm2 norm
//TODO
func (m *Matrix) Norm2() float64 {
	return 0
}

//NormInf norm
//TODO
func (m *Matrix) NormInf() float64 {
	return 0

}
