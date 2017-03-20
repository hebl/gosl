package matrix

//Copy matrix hard copy
func (m *Matrix) Copy() *Matrix {
	m2 := New(m.size1, m.size2)
	for i := range m.data {
		m2.data[i] = m.data[i]
	}
	return m2
}

//T matrix transport
func (m *Matrix) T() *Matrix {
	m2 := New(m.size2, m.size1)

	for i := 0; i < m.size1; i++ {
		for j := 0; j < m.size2; j++ {
			m2.Set(j, i, m.At(i, j))
		}
	}

	return m2
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
