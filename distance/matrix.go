package distance

import (
	"fmt"
	"strings"
)

// A distance matrix
// Optimised for storage efficiency. Zero value can be sused immediately.
type Matrix struct {
	size int
	data []float64 // stores the value as (0,1),     (0,2),(1,2),     (0,3),(1,3),(2,3),     (0,4),(1,4), 2,4),(3,4),    ...
}

// Get distance between i and j.
func (m *Matrix) Dist(i, j int) float64 {
	if i == j {
		return 0.
	}
	idx := index(i, j)
	if idx >= len(m.data) {
		return 0.
	}
	return m.data[idx]
}

func index(i, j int) int {
	if i < j {
		return i + j*(j-1)/2
	} else {
		return index(j, i)
	}
}

// Set a distance between i and j.
// Size increases as needed.
func (m *Matrix) Set(i, j int, d float64) {
	if i == j {
		return
	}
	idx := index(i, j)
	if i >= m.size || j >= m.size || idx >= len(m.data) { // auto extend matrix when needed and adjust size
		m.data = append(m.data, make([]float64, 1+idx-len(m.data))...)
		if i > j {
			m.size = i + 1
		} else {
			m.size = j + 1
		}
	}
	m.data[index(i, j)] = d
}

// Provide current size n of matrix (n x n)
// May dynamically increase when elements are added.
func (m *Matrix) Size() int {
	return m.size
}

// String for display
func (m *Matrix) String() string {
	if m == nil || m.size == 0 {
		return "<empty matrix>"
	}
	sb := new(strings.Builder)
	for j := 0; j < m.size; j++ {
		if j == 11 {
			fmt.Fprintf(sb, "\t[...%d]", m.size-1)
			break
		}
		fmt.Fprintf(sb, "\t%8d", j)
	}
	for i := 0; i < m.size; i++ {
		fmt.Fprintf(sb, "\n%5d\t", i)
		for j := 0; j < m.size; j++ {
			if j > 10 {
				break
			}
			fmt.Fprintf(sb, "%02.6f\t", m.Dist(i, j))
		}
		if i == 10 {
			fmt.Fprintf(sb, "\n[...%d]", m.size-1)
			break
		}
	}
	fmt.Fprintln(sb)
	return sb.String()
}
