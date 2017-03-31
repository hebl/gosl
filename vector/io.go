// Copyright 2017 HE Boliang. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vector

import (
	"fmt"
	"io"
	"os"
)

var (
	tab = []byte("   ")
	nn  = []byte("\n")
	ln  = []byte("\n")
)

//Fprintf write format
func (v *Vector) Fprintf(w io.Writer, format string) {
	var s string
	for i := 0; i < v.size; i++ {
		w.Write(tab)

		s = fmt.Sprintf(" "+format, v.At(i))
		w.Write([]byte(s))

		w.Write(ln)
	}
}

//Print write format
func (v *Vector) Print(format string) {
	os.Stdout.Write(ln)
	v.Fprintf(os.Stdout, format)
	os.Stdout.Write(ln)
}
