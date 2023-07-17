package cluster

import "math"

// Find the Medoïd and the average distance from the cluster elements to the medoid (including the medoid itself).
func (cc *CContext) Medoid(c *Cluster) (med int, dist float64) {

	if c == nil || len(c.obj) == 0 {
		panic("cannot find medoid of an empty cluster)")
	}

	if len(c.obj) == 1 {
		return c.obj[0], 0.
	}

	med = c.obj[0]
	dist = math.Inf(+1)
	for _, m := range c.obj {
		ds := c.sumDistFrom(m, cc.ed)
		if ds < dist {
			med = m
			dist = ds
		}
	}
	return med, dist / float64(len(c.obj))
}

func (c *Cluster) sumDistFrom(x int, ed Dist) float64 {
	d := 0.
	for _, e := range c.obj {
		d += ed(x, e)
	}
	return d
}
