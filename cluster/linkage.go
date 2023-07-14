package cluster

// LinkDist is used to measure linkage distance between clusters
type LinkDist func(c1, c2 *Cluster) float64

// Element wise distance. Typically from a distance matrix.
type Dist func(i, j int) float64

const MAXDISTANCE float64 = 1e99

// Single linkage (min of d(a,b))
func SingleLinkage(dist Dist) LinkDist {

	return func(c1, c2 *Cluster) float64 {
		d := MAXDISTANCE
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

// Find the medoid of a cluster.
// The medoid is the object that minimize the sum of the distance from it to all other elements of the cluster.
func Medoid(c *Cluster) int {
	l := len(c.obj)
	if l == 0 {
		panic("empty cluster")
	}
	if l == 1 {
		return c.obj[0]
	}

	panic("todo")
}
