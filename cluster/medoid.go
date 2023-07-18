package cluster

import (
	"fmt"
	"math"
)

var FLAGMEDOID bool // to compute or not the medoids

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

func (cc *CContext) DumpMedoids() {

	fmt.Println("\nList of medoids and average internal distance per cluster :\n\n\tlevel\tmedoïd\tmedoïd dist\t\t\t\tcluster")
	for k, v := range cc.cls {
		if len(k.obj) == 1 {
			// do not dump single clusters
			continue
		}
		if v {
			fmt.Printf("root%5d\t", k.level)
		} else {
			fmt.Printf("    %5d\t", k.level)
		}
		fmt.Printf("[%d]\t%2.6f \t --medoïd for--> \t%v \n", cc.medoids[k].m, cc.medoids[k].d, k.obj)
	}
	fmt.Println()

}
