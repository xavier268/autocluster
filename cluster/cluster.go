package cluster

import (
	"context"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/xavier268/autocluster/distance"
)

type Cluster struct {
	obj   []int    // list of objects (0 based)
	left  *Cluster // tree of hierarchical clusters
	right *Cluster // tree of hierarchical clusters
	linkd float64  // distance between left and right sub clusters
	level int      // from the leaf = 0
}

// Cluster context
type CContext struct {
	ctx     context.Context
	cls     map[*Cluster]bool     // Set of free clusters, ie clusters that can be merged further
	ld      LinkDist              // Link distance to use
	ed      Dist                  // Element distance
	medoids map[*Cluster]struct { // medoid and distance - may stay unitiallized
		m int
		d float64
	}
}

func NewEmptyCContext(ctx context.Context) *CContext {
	return &CContext{
		ctx: ctx,
		cls: make(map[*Cluster]bool),
		ld:  nil,
		ed:  nil,
	}
}

// Defines how to converts the element distance into a linkage distance.
type LinkOption func(Dist) LinkDist

func NewCContexMatrix(ctx context.Context, mat *distance.Matrix, linkOption LinkOption) *CContext {
	cc := NewEmptyCContext(ctx)
	cc.ld = linkOption(mat.Dist)
	cc.ed = mat.Dist
	cc.medoids = make(map[*Cluster]struct {
		m int
		d float64
	})
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
	cc.medoids[c] = struct {
		m int
		d float64
	}{obj, 0.}
}

// Merge 2 clusters. Old clusters become inactive, new cluster is now active.
func (cc *CContext) merge(c1, c2 *Cluster, d float64) {
	c := &Cluster{
		obj:   append(c1.obj, c2.obj...),
		left:  c1,
		right: c2,
		linkd: d,
	}
	if c1.level > c2.level {
		c.level = c1.level + 1
	} else {
		c.level = c2.level + 1
	}
	cc.cls[c1] = false
	cc.cls[c2] = false
	cc.cls[c] = true
	if FLAGMEDOID {
		m, d := cc.Medoid(c)
		cc.medoids[c] = struct {
			m int
			d float64
		}{m, d}
	}
}

// Merge until there is only 1 cluster left.
func (cc *CContext) MergeAll() {

	fmt.Fprint(os.Stderr, "\n")
	for i := 0; !cc.Merge(); i++ {
		fmt.Fprintf(os.Stderr, "\rComputing clusters %d/%d        ", i+2, len(cc.cls)-i-1)
	}
	fmt.Println()
}

// Get the root cluster
func (cc *CContext) Root() *Cluster {
	for k, v := range cc.cls {
		if v {
			return k
		}
	}
	return nil
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
	var dmin float64 = math.Inf(+1)
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

// String representation of a single cluster, for debugging.
func (c *Cluster) String() string {
	return fmt.Sprintf("%p\t%v\t%p , %p\t d=%2.6f", c, c.obj, c.left, c.right, c.linkd)
}

func (cc *CContext) Dump() {
	for k, v := range cc.cls {
		fmt.Printf("%v:%p\t%s\n", v, k, k)
	}
}

// Tree representation of a group of clusters
func (c *Cluster) Tree() string {
	sb := new(strings.Builder)
	fmt.Fprintln(sb, "Table of all clusters :\n\nlink dist.\tlevel\tcluster content .....................")
	c.tree(sb, "", false) // do not skip single nodes
	return sb.String()
}

func (c *Cluster) tree(sb *strings.Builder, prefix string, skipSingle bool) {
	if skipSingle && len(c.obj) <= 1 {
		return // skip single nodes ...
	}
	const inc = "\t"
	fmt.Fprintf(sb, "%2.6f\t%d\t%s%v\n", c.linkd, c.level, prefix, c.obj)
	if c.left != nil {
		c.left.tree(sb, inc+prefix, skipSingle)
		c.right.tree(sb, inc+prefix, skipSingle)
	}
}

func (c *Cluster) Dendrogram(names []string, minsize int) string {
	if c == nil {
		return ""
	}
	sb := new(strings.Builder)
	fmt.Fprintln(sb, "Annotated dendrogramme of clusters :\n(cluster content ....)\t\t( level / link distance )")
	c.dendrogram(sb, "", names, true, minsize, true)
	return sb.String()
}

func (c *Cluster) dendrogram(sb *strings.Builder, prefix string, names []string, isTail bool, minsize int, truncate bool) {
	if c == nil {
		return
	}
	if len(c.obj) == 0 {
		panic("internal error - dentogram with empty cluster")
	}
	if len(c.obj) == 1 || len(c.obj) < minsize {
		for _, obj := range c.obj {
			pp := fmt.Sprintf("%s+---[%d]%s", prefix, obj, strings.Repeat(" - ", 80))
			fmt.Fprintf(sb, "%s  %s\n", pp[:100], names[obj])
		}
		return
	}
	if truncate && len(c.obj) > 30 {
		fmt.Fprintf(sb, "%s+---%v[...]\t(%d / %2.6f)\n", prefix, c.obj[:28], c.level, c.linkd)
	} else {
		fmt.Fprintf(sb, "%s+---%v\t(%d / %2.6f)\n", prefix, c.obj, c.level, c.linkd)
	}
	if isTail {
		prefix += "   "
	} else {
		prefix += "|  "
	}
	//fmt.Fprintln(sb, prefix)
	c.left.dendrogram(sb, prefix, names, false, minsize, truncate)
	c.right.dendrogram(sb, prefix, names, true, minsize, truncate)
	fmt.Fprintln(sb, prefix)
}

func isIn(v int, sl []int) bool {
	for _, s := range sl {
		if v == s {
			return true
		}
	}
	return false
}

var _ = isIn
