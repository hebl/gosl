// Copyright 2017 HE Boliang. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vector

import "testing"

var (
	v = Rand(10)
	a = Rand(10)
)

func TestAdd(t *testing.T) {
	v.Add(a)
}

func TestSub(t *testing.T) {

	v.Sub(a)
}

func TestMul(t *testing.T) {

	v.Mul(a)
}

func TestDiv(t *testing.T) {

	v.Div(a)
}

func TestScale(t *testing.T) {

	v.Scale(10.213)
}

func TestAddConstant(t *testing.T) {

	v.AddConstant(10.213)
}

func TestNorm(t *testing.T) {

	n1 := v.Norm1()
	n2 := v.Norm2()

	t.Log(n1, n2)
}

func TestSort(t *testing.T) {
	n1 := Rand(20)
	t.Log("Before Sorted")
	n1.Print("%7.2f")
	n1.Sort()
	t.Log("Sorted")
	n1.Print("%7.2f")

}

func TestMedian(t *testing.T) {
	n1 := Rand(20)
	t.Log("Before Sorted")
	n1.Print("%7.2f")

	t.Log(n1.Median())
	n1.Sort()
	t.Log("Sorted")
	n1.Print("%7.2f")
}
