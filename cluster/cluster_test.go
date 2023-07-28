package cluster

import (
	"context"
	"fmt"
	"testing"

	"github.com/xavier268/autocluster/distance"
)

var data = []string{
	"Il fait beau",
	"Il fait beau et le ciel est bleu",
	"Il ne fait pas beau du tout",
	"Le petit chien se promène",
	"Hellow World, how are you doing today ?",
	"Il était une fois une princesse et un chevalier, qui aimait bien chasser les dragons. La princesse vivait dans un chateau fort.",
	"lkjsd qsdv:lkjqsdvl  lkjqdsfv lkjqsdf   qlkjqdfvljkqdf qldkfj vlkqdjfv  lkjqdflkjqdf  lkjqdflkj",
}

func TestStringDistance(t *testing.T) {
	displayData()
	VerifyStringMatrix(t)
}

func TestStringCluster(t *testing.T) {

	var cc *CContext
	var k *Cluster

	// Prepare test matrix
	displayData()
	mat := distance.ComputeString(data)
	fmt.Println(mat)

	fmt.Println("Computing clusters - single linkage")
	cc = NewCContexMatrix(context.Background(), mat, SingleLinkage)
	cc.MergeAll()
	k = cc.Root()
	fmt.Println(k.Tree())

	fmt.Println("Computing clusters - complete linkage")
	cc = NewCContexMatrix(context.Background(), mat, CompleteLinkage)
	cc.MergeAll()
	k = cc.Root()
	fmt.Println(k.Tree())

	fmt.Println("Computing clusters - UPGMA linkage")
	cc = NewCContexMatrix(context.Background(), mat, UPGMALinkage)
	cc.MergeAll()
	k = cc.Root()
	fmt.Println(k.Tree())

}

func displayData() {
	fmt.Println("Test data :")
	for i, x := range data {
		fmt.Printf("%5d:\t%s\n", i, x)
	}
	fmt.Println()
}

func VerifyStringMatrix(t *testing.T) {
	for _, x := range data {
		for _, y := range data {
			display(x, y)
			if distance.DistString(x, y) < 0 {
				t.Fatalf("Distance should not be negative between <%s> and <%s>", x, y)
			}
			if distance.DistString(x, y) != distance.DistString(y, x) {
				t.Fatalf("Distance should be symetric")
			}
			if x == y && distance.DistString(x, y) != 0. {
				t.Fatalf("Distance should be exactly 0. for identical strings")
			}
		}
	}

	fmt.Println(distance.ComputeString(data))
}

func display(x, y string) {
	fmt.Printf("<%s>\n<%s>\n\t distance => \t %0.7f\n", x, y, distance.DistString(x, y))
}
