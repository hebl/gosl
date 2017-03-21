package matrix

import "github.com/hebl/gosl/errors"

//Add two matrix add
func (m *Matrix) Add(m2 *Matrix) error {
	if (m.size1 != m2.size1) && (m.size2 != m2.size2) {
		return errors.ErrBADLEN
	}
	totalsize := m.size1 * m.size2
	for i := 0; i < totalsize; i++ {
		m.data[i] += m2.data[i]
	}
	return nil
}

//Sub two matrix sub
func (m *Matrix) Sub(m2 *Matrix) error {
	if (m.size1 != m2.size1) && (m.size2 != m2.size2) {
		return errors.ErrBADLEN
	}
	totalsize := m.size1 * m.size2
	for i := 0; i < totalsize; i++ {
		m.data[i] -= m2.data[i]
	}
	return nil
}

//MulElement two matrix element mul
func (m *Matrix) MulElement(m2 *Matrix) error {
	if (m.size1 != m2.size1) && (m.size2 != m2.size2) {
		return errors.ErrBADLEN
	}
	totalsize := m.size1 * m.size2
	for i := 0; i < totalsize; i++ {
		m.data[i] *= m2.data[i]
	}
	return nil
}

//DivElement two matrix element mul
func (m *Matrix) DivElement(m2 *Matrix) error {
	if (m.size1 != m2.size1) && (m.size2 != m2.size2) {
		return errors.ErrBADLEN
	}
	totalsize := m.size1 * m.size2
	for i := 0; i < totalsize; i++ {
		m.data[i] /= m2.data[i]
	}
	return nil
}

//Scale two matrix element mul
func (m *Matrix) Scale(x float64) {
	totalsize := m.size1 * m.size2
	for i := 0; i < totalsize; i++ {
		m.data[i] *= x
	}
}

//AddConstant two matrix element mul
func (m *Matrix) AddConstant(x float64) {
	totalsize := m.size1 * m.size2
	for i := 0; i < totalsize; i++ {
		m.data[i] += x
	}
}
