// Copyright 2017 HE Boliang. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vector

import "sort"

//Len sort Len interface
func (v *Vector) Len() int {
	return v.size
}

//Swap sort swap element
func (v *Vector) Swap(i, j int) {
	v.data[i], v.data[j] = v.data[j], v.data[i]
}

//Less sort less interface
func (v *Vector) Less(i, j int) bool {
	return v.data[i] < v.data[j]
}

//Sort sort vector
func (v *Vector) Sort() {
	sort.Sort(v)
}

//IsSorted is sorted?
func (v *Vector) IsSorted() bool {
	return sort.IsSorted(v)
}
