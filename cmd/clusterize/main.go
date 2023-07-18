package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/xavier268/autocluster/cluster"
	"github.com/xavier268/autocluster/distance"
)

func main() {

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
	flag.StringVar(&distance.CACHEFILENAME, "cache", filepath.Join(os.TempDir(), "fileDistance.cache"), "cache file location")
	flag.BoolVar(&FLAGHELP, "h", false, "print version, usage and exit")
	flag.BoolVar(&FLAGDEND, "d", true, "print dendrogramme")
	flag.BoolVar(&FLAGTREE, "t", false, "print tree")
	flag.BoolVar(&FLAGNAMES, "f", true, "list file names")
	flag.BoolVar(&FLAGMATRIX, "dm", true, "print distance matrix (truncated)")
	flag.BoolVar(&cluster.FLAGMEDOID, "med", true, "compute and print medoids (dendrogramme)")

	flag.Parse()
	args := flag.Args()

	if FLAGHELP {
		fmt.Printf("Unsupervised clustering of files\n\tdistance version : %s\n\tcluster version  : %s\nUsage :\n", distance.VERSION, cluster.VERSION)
		flag.PrintDefaults()
		return
	}

	files := []string{}
	for _, f := range args {
		files = append(files, distance.FilesInFolder(f)...)
	}
	if FLAGNAMES {
		fmt.Println("Processing files :")
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
		cc = cluster.NewCContexMatrix(ctx, mat, cluster.SingleLinkage)
	case "complete":
		cc = cluster.NewCContexMatrix(ctx, mat, cluster.CompleteLinkage)
	case "upgma":
		cc = cluster.NewCContexMatrix(ctx, mat, cluster.UPGMALinkage)
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

	if cluster.FLAGMEDOID {
		fmt.Println()
		cc.DumpMedoids()
	}

}
