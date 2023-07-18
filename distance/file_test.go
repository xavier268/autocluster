package distance

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestExtractTextTestFiles(t *testing.T) {

	data := FilesInFolder(filepath.Join("..", "testFiles"))

	for i, d := range data {

		dumpFile(i, d)
	}

	mat := ComputeFiles(data...)
	fmt.Println("Files processed :")
	for i, d := range data {
		fmt.Printf("%2d\t%s\n", i, d)
	}
	fmt.Println(mat)

}

func TestExtractProjectFiles(t *testing.T) {

	data := FilesInFolder("..")

	for i, d := range data {

		dumpFile(i, d)
	}

	mat := ComputeFiles(data...)
	fmt.Println("Files processed :")
	for i, d := range data {
		fmt.Printf("%2d\t%s\n", i, d)
	}
	fmt.Println(mat)

}

func dumpFile(i int, fname string) {
	fmt.Printf("\n\n%d)\t======= Dumping extracted text content of %s =========\n", i, fname)
	x := ExtractText(fname)
	fmt.Println(string(x))
}
