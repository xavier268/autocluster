// Package cluster provides clustering in a distance agnostic way
package cluster

import (
	"context"
	"fmt"
	"math"
	"os"
	"sort"

	"github.com/xavier268/autocluster/distance"
)

type Cluster struct {
	obj     []int    // list of objects (0 based)
	left    *Cluster // left sub-tree of hierarchical clusters
	right   *Cluster // right sub-tree of hierarchical clusters
	linkd   float64  // distance between left and right sub clusters
	level   int      // from the leaf = 0
	med     int      // medoid for cluster
	meddist float64  // average medoid distance
}

// Cluster context
type CContext struct {
	ctx   context.Context
	cls   []*Cluster // a slice of free clusters - always at least one !
	ld    LinkDist   // Link distance to use
	ed    Dist       // Element distance
	names []string   // Name associated with each object
}

func newEmptyCContext(ctx context.Context) *CContext {
	return &CContext{
		ctx: ctx,
		cls: []*Cluster{},
		ld:  nil,
		ed:  nil,
	}
}

// Defines how to converts the element distance into a linkage distance.
type LinkOption func(Dist) LinkDist

// Creates a new cluster context with a linkage distance,
// based upon a distance matrix and a list of names for each object.
func NewCContex(ctx context.Context, mat *distance.Matrix, linkOption LinkOption, names []string) *CContext {
	cc := newEmptyCContext(ctx)
	cc.ld = linkOption(mat.Dist)
	cc.ed = mat.Dist
	for i := 0; i < mat.Size(); i++ {
		cc.addObject(i)
	}
	if len(names) == mat.Size() {
		cc.names = names
	} else {
		cc.names = make([]string, mat.Size())
	}
	return cc
}

// Add a new single object cluster
func (cc *CContext) addObject(obj int) {
	c := &Cluster{
		obj:     []int{obj},
		med:     obj,
		meddist: 0,
	}
	cc.cls = append(cc.cls, c)
}

// Merge 2 clusters. Old clusters removed from free cluster list, new cluster added.
func (cc *CContext) merge2clusters(c1, c2 *Cluster, d float64) {
	c := &Cluster{
		obj:   append(c1.obj, c2.obj...),
		left:  c1,
		right: c2,
		linkd: d,
	}
	sort.Ints(c.obj)
	c.med, c.meddist = cc.medoid(c)
	if c1.level > c2.level {
		c.level = c1.level + 1
	} else {
		c.level = c2.level + 1
	}
	// replace c1 by c
	for i, v := range cc.cls {
		if v == c1 {
			cc.cls[i] = c
			break
		}

	}
	// remove c2
	for i, v := range cc.cls {
		if v == c2 {
			if i+1 < len(cc.cls) {
				cc.cls = append(cc.cls[:i], cc.cls[i+1:]...)
				break
			} else {
				cc.cls = cc.cls[:i]
				break
			}
		}
	}
}

// Merge until there is only 1 cluster left.
func (cc *CContext) Merge() {
	nb := len(cc.cls) // initial length will change !
	fmt.Fprint(os.Stderr, "\n")
	for i := 0; !cc.mergeStep(); i++ {
		fmt.Fprintf(os.Stderr, "\rComputing clusters %d/%d        ", i+2, nb)
	}
	fmt.Println()
}

// Get the root cluster
func (cc *CContext) Root() *Cluster {
	return cc.cls[0]
}

// Make a single merge step. Return true when finished (only 1 cluster left)
func (cc *CContext) mergeStep() (finished bool) {

	if len(cc.cls) <= 1 {
		return true // done !
	}
	var dmin float64 = math.Inf(+1)
	var c1, c2 *Cluster = nil, nil
	// Only compare 0 <= i < j < len(cc.cls)
	for i := 0; i < len(cc.cls)-1; i++ {
		for j := i + 1; j < len(cc.cls); j++ {
			d := cc.ld(cc.cls[i], cc.cls[j])
			if c1 == nil || d < dmin {
				c1, c2 = cc.cls[i], cc.cls[j]
				dmin = d
			}
		}
	}
	if c1 == nil {
		panic("internal error - merge should have happened")
	}
	cc.merge2clusters(c1, c2, dmin)
	return false
}
