package matrix

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
