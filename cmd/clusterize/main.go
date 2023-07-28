package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/xavier268/autocluster/cluster"
	"github.com/xavier268/autocluster/distance"
)

func main() {

	const VERSION = "0.8.5"
	const COPYRIGHT = "(c) xavier268<at>github.com, 2023"

	var ( // flags
		FLAGLINK    string // type of linkage
		FLAGMINSIZE int    // minimum cluster size (for display)
		FLAGHELP    bool
		FLAGDEND    bool // print dendrogramme
		FLAGTREE    bool // print tree
		FLAGNAMES   bool // print file names
		FLAGMATRIX  bool // print distance matrix (truncated)

	)

	flag.StringVar(&FLAGLINK, "link", "complete", "select type of linkage from single, complete, upgma")
	flag.IntVar(&FLAGMINSIZE, "min", 0, "set this value to a high number to get less clusters")
	// flag.StringVar(&distance.CACHEFILENAME, "cache", distance.CACHEFILENAME, "cache file location")
	flag.BoolVar(&FLAGHELP, "h", false, "print version, usage and exit")
	flag.BoolVar(&FLAGDEND, "d", true, "print dendrogramme view")
	flag.BoolVar(&FLAGTREE, "t", true, "print tree view")
	flag.BoolVar(&FLAGNAMES, "f", true, "list file names")
	flag.BoolVar(&FLAGMATRIX, "dm", true, "print distance matrix")
	// flag.BoolVar(&cluster.FLAGMEDOID, "med", true, "compute and print medoïd view") // always do it !

	flag.Parse()
	args := flag.Args()

	if FLAGHELP {
		fmt.Printf("Unsupervised clustering of files\n%s - v%s\nUsage :\n", COPYRIGHT, VERSION)
		fmt.Printf("Using cache file : %s\n", distance.CACHEFILENAME)
		flag.PrintDefaults()
		return
	}

	files := []string{}
	for _, f := range args {
		files = append(files, distance.FilesInFolder(f)...)
	}
	if FLAGNAMES {
		fmt.Println("\nList of files to process :\n\n  id\tfull path name")
		for i, f := range files {
			fmt.Printf("%4d\t%s\n", i, f)
		}
	}
	mat := distance.ComputeFiles(files...)
	if FLAGMATRIX {
		fmt.Println()
		fmt.Println(mat)
	}
	var cc *cluster.CContext
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	switch strings.ToLower(FLAGLINK) {
	case "single":
		cc = cluster.NewCContex(ctx, mat, cluster.SingleLinkage, files)
	case "complete":
		cc = cluster.NewCContex(ctx, mat, cluster.CompleteLinkage, files)
	case "upgma":
		cc = cluster.NewCContex(ctx, mat, cluster.UPGMALinkage, files)
	default:
		fmt.Fprintln(os.Stderr, "you selected an invalid linkage option")
		flag.PrintDefaults()
		return
	}
	cc.MergeAll()
	root := cc.Root()
	if FLAGDEND {
		fmt.Println()
		fmt.Println(root.Dendrogram(files, FLAGMINSIZE))
	}
	if FLAGTREE {
		fmt.Println()
		fmt.Println(root.Tree())
	}

}
