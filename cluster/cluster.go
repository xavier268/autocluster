package cluster

import (
	"context"

	"github.com/xavier268/autocluster/distance"
)

type Cluster struct {
	obj   []int    // list of objects (0 based)
	left  *Cluster // tree of hierarchical clusters
	right *Cluster // tree of hierarchical clusters
	linkd float64  // distance between left and right sub clusters
}

// Cluster context
type CContext struct {
	ctx context.Context
	cls map[*Cluster]bool // Set of free clusters, ie clusters that can be merged further
	ld  LinkDist          // Link distance to use
}

func NewEmptyCContext(ctx context.Context) *CContext {
	return &CContext{
		ctx: ctx,
		cls: make(map[*Cluster]bool),
		ld:  nil,
	}
}

// Defines how to converts the element distance into a linkage distance.
type LinkOption func(Dist) LinkDist

func NewCContexMatrix(ctx context.Context, mat *distance.Matrix, linkOption LinkOption) *CContext {
	cc := NewEmptyCContext(ctx)
	cc.ld = linkOption(mat.Dist)
	for i := 0; i < mat.Size(); i++ {
		cc.AddObject(i)
	}
	return cc
}

// Add a new single object cluster
func (cc *CContext) AddObject(obj int) {
	c := &Cluster{
		obj: []int{obj},
	}
	cc.cls[c] = true
}

// Merge 2 clusters. Old clusters become inactive, new cluster is now active.
func (cc *CContext) merge(c1, c2 *Cluster, d float64) {
	c := &Cluster{
		obj:   append(c1.obj, c2.obj...),
		left:  c1,
		right: c2,
		linkd: d,
	}
	cc.cls[c1] = false
	cc.cls[c2] = false
	cc.cls[c] = true
}

// Make a single merge step. Return true when finished (only 1 cluster left)
func (cc *CContext) Merge() (finished bool) {

	var free []*Cluster // collect free clusters that could be merged
	for k, v := range cc.cls {
		if v {
			free = append(free, k)
		}
	}
	if len(free) <= 1 {
		return true
	}
	var dmin float64 = MAXDISTANCE
	var c1, c2 *Cluster = nil, nil
	// Only compare 0 <= i < j < len(free)
	for i := 0; i < len(free)-1; i++ {
		for j := i + 1; j < len(free); j++ {
			d := cc.ld(free[i], free[j])
			if c1 == nil || d < dmin {
				c1, c2 = free[i], free[j]
				dmin = d
			}
		}
	}
	if c1 == nil {
		panic("internal error - merge should have happened")
	}
	cc.merge(c1, c2, dmin)
	return false
}
