package distance

import (
	"compress/gzip"
)

// Compute a normalized compression distance between two strings using gzip.
// Should normally be between 0.0 and 1.0.
// See article attached, annexe A.
func DistString(x, y string) float64 {

	if x == y {
		return 0.
	}

	dx := ziplen(x)
	dy := ziplen(y)

	if dx == 0 && dy == 0 {
		return 0.
	}

	if dx > dy {
		return (ziplen(y, x) - dy) / dx
	}
	return (ziplen(x, y) - dx) / dy

}

// implements a file that stores nothing but just count bytes.
type counterFile struct {
	count int
}

// Write implements io.Writer.
func (f *counterFile) Write(p []byte) (n int, err error) {
	f.count = f.count + len(p)
	return len(p), nil
}

// return the zip length of x
func ziplen(xx ...string) float64 {
	if len(xx) == 0 {
		return 0.
	}
	c := new(counterFile)
	w := gzip.NewWriter(c)
	for _, x := range xx {
		w.Write([]byte(x))
	}
	w.Close() // require to flush data !
	return float64(c.count)
}

// Compute the distance matrix between strings
func ComputeString(ss []string) (mat *Matrix) {
	mat = new(Matrix)
	for i := 0; i < len(ss); i++ {
		for j := i + 1; j < len(ss); j++ {
			mat.Set(i, j, DistString(ss[i], ss[j]))
		}
	}
	return mat
}
