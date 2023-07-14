package distance

import (
	"math"
	"testing"
)

func TestMatrix(t *testing.T) {

	m := &Matrix{}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			m.Set(i, j, math.Abs(float64(i-j)))
		}
	}

	if m.Dist(1, 1) != 0 || m.Dist(6, 6) != 0 {
		t.FailNow()
	}

	if m.Dist(1, 3) != m.Dist(3, 1) {
		t.FailNow()
	}

	if m.Dist(2, 4) != 2. || m.Dist(4, 2) != 2. {
		t.FailNow()
	}

	if m.Dist(8, 2) != 0. {
		t.FailNow()
	}
}
