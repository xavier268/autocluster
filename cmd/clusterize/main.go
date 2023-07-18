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

	link := ""
	flag.StringVar(&link, "link", "complete", "select type of linkage from single, complete, upgma")
	minsize := 0
	flag.IntVar(&minsize, "min", 0, "minimum cluster size to display")
	flag.Parse()
	args := flag.Args()
	// fmt.Println("Args :", args)

	files := []string{}
	for _, f := range args {
		files = append(files, distance.FilesInFolder(f)...)
	}
	fmt.Println("Processing files :")
	for i, f := range files {
		fmt.Printf("%4d\t%s\n", i, f)
	}
	mat := distance.ComputeFiles(files...)
	fmt.Println()
	fmt.Println(mat)
	var cc *cluster.CContext
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	switch strings.ToLower(link) {
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
	fmt.Println()
	root := cc.Root()
	fmt.Println(root.Dendrogram(files, minsize))
}
