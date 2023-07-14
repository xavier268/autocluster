package distance

// A distance matrix
type Matrix struct {
	data []float64 // stores the value as (0,1),     (0,2),(1,2),     (0,3),(1,3),(2,3),     (0,4),(1,4), 2,4),(3,4),    ...
}

// Get distance between i and j.
func (m *Matrix) At(i, j int) float64 {
	if i == j {
		return 0.
	}
	return m.data[index(i, j)]
}

func index(i, j int) int {
	if i < j {
		return i + j*(j-1)/2
	} else {
		return index(j, i)
	}
}

// Set a distance between i and j
func (m *Matrix) Set(i, j int, d float64) {
	if i == j {
		return
	}
	if m.data == nil {
		m.data = make([]float64, 0, index(i+1, j+1))
	}
	m.data[index(i, j)] = d
}
