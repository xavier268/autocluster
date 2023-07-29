package cluster

import (
	"fmt"
	"sort"
	"strings"
)

// A map of all medoids with the number of clusters they appear in
func (cc *CContext) ListMedoids() map[int]int {
	m := make(map[int]int, 0)
	for _, c := range cc.ListClusters() {
		m[c.med] = m[c.med] + 1
	}
	return m
}

func (cc *CContext) ListClusters() []*Cluster {
	return listClusters(cc.Root())
}

// recursively get a list of all clusters
func listClusters(root *Cluster) []*Cluster {
	if root == nil {
		return nil
	}
	list := append(listClusters(root.left), listClusters(root.right)...)
	return append(list, root)
}

func (cc *CContext) ViewByLinkDistance(list []*Cluster) string {
	sb := new(strings.Builder)
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].linkd < list[j].linkd
	})

	fmt.Fprintln(sb, "Clusters ranked by link distance\nLinkDist\tMedoid\tMedoidDist\tObject ids")
	for _, c := range list {
		if len(c.obj) > 1 {
			fmt.Fprintf(sb, "%1.6f\t%4d\t%1.6f\t%v\n", c.linkd, c.med, c.meddist, c.obj)
		}
	}
	return sb.String()
}

func (cc *CContext) ViewByMedoidDistance(list []*Cluster) string {
	sb := new(strings.Builder)
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].meddist < list[j].meddist
	})

	fmt.Fprintln(sb, "Clusters ranked by medoïd distance\nMedoid\tMedoidDist\tObject ids")
	for _, c := range list {
		if len(c.obj) > 1 {
			fmt.Fprintf(sb, "%4d\t%1.6f\t%v\n", c.med, c.meddist, c.obj)
		}
	}
	return sb.String()
}

func (cc *CContext) ViewMedoids(mm map[int]int) string {
	sb := new(strings.Builder)
	meds := []int{}
	for med := range mm {
		meds = append(meds, med)
	}
	sort.SliceStable(meds, func(i, j int) bool {
		return mm[meds[i]] > mm[meds[j]]
	})

	fmt.Fprintln(sb, "List of medoïds\nMedoid\tCount\tName")
	for _, m := range meds {
		if mm[m] > 1 {
			fmt.Fprintf(sb, "%4d\t%4d\t%s\n", m, mm[m], cc.names[m])
		}
	}
	return sb.String()
}

// String representation of a single cluster, for debugging.
func (c *Cluster) String() string {
	return fmt.Sprintf("%p\t%v\t%p , %p\t d=%2.6f", c, c.obj, c.left, c.right, c.linkd)
}

// ViewAsTree representation of a group of clusters
func (c *Cluster) ViewAsTree() string {
	sb := new(strings.Builder)
	fmt.Fprintln(sb, "Tree view of all clusters :\nLink dist.\tLevel\tCluster content")
	c.tree(sb, "", true) //  skip single nodes
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

// obsolete ...
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
