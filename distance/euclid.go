package distance

import "math"

type Vect = []float64

func dist(v, w Vect) float64 {
	d := 0.
	for i := 0; i < len(v) && i < len(w); i++ {
		d = d + (v[i]-w[i])*(v[i]-w[i])
	}
	return math.Sqrt(d)
}

func ComputeEuclid(vects []Vect) (mat *Matrix) {
	mat = new(Matrix)
	for i := 0; i < len(vects); i++ {
		for j := i + 1; j < len(vects); j++ {
			mat.Set(i, j, dist(vects[i], vects[j]))
		}
	}
	return mat
}
