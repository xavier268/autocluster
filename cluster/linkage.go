package cluster

import "math"

// LinkDist is used to measure linkage distance between clusters
type LinkDist func(c1, c2 *Cluster) float64

// Element wise distance. Typically from a distance matrix.
type Dist func(i, j int) float64

// Single linkage (min of d(a,b))
func SingleLinkage(dist Dist) LinkDist {

	return func(c1, c2 *Cluster) float64 {
		d := math.Inf(+1)
		for _, e1 := range c1.obj {
			for _, e2 := range c2.obj {
				dd := dist(e1, e2)
				if dd < d {
					d = dd
				}
			}
		}
		return d
	}
}

// Complete Linkage (max(a,b))
func CompleteLinkage(dist Dist) LinkDist {

	return func(c1, c2 *Cluster) float64 {
		var d float64
		for _, e1 := range c1.obj {
			for _, e2 := range c2.obj {
				dd := dist(e1, e2)
				if dd > d {
					d = dd
				}
			}
		}
		return d
	}
}

// UPGMA Linkage
func UPGMALinkage(dist Dist) LinkDist {
	return func(c1, c2 *Cluster) float64 {
		var d float64
		for _, e1 := range c1.obj {
			for _, e2 := range c2.obj {
				d = d + dist(e1, e2)
			}
		}
		return d / float64(len(c1.obj)*len(c2.obj))
	}
}

// Mini-Max linkage
// TODO

// Hausdorff Linkage
// TODO

// Medoid linkage
// TODO

// Minimum Sum Medoid linkage
// TODO
