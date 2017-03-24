package matrix

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
func (m *Matrix) Fprintf(w io.Writer, format string) {
	var s string
	for i := 0; i < m.size1; i++ {
		w.Write(tab)
		for j := 0; j < m.size2; j++ {
			s = fmt.Sprintf(" "+format, m.At(i, j))
			w.Write([]byte(s))
		}
		w.Write(ln)
	}

}

//Printf print matrix
func (m *Matrix) Printf(format string) {
	os.Stdout.Write(nn)
	m.Fprintf(os.Stdout, format)
	os.Stdout.Write(nn)
}

//Print print matrix to stdout
func (m *Matrix) Print() {
	m.Printf("%10.3f")
}
