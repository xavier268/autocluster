package distance

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestExtractText(t *testing.T) {

	data := []string{
		"test.html.zip",
		"test.html",
		"test.docx",
		"test.xlsx",
	}

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
	f := filepath.Join("..", "testFiles", fname)
	x := ExtractText(f)
	fmt.Println(string(x))
}
