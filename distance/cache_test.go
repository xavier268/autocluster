package distance

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestCache(t *testing.T) {

	CACHEFILENAME = filepath.Join(os.TempDir(), "debug.cache")

	fmt.Println("Using temporary cache in ", CACHEFILENAME)
	os.Remove(CACHEFILENAME)
	defer os.Remove(CACHEFILENAME)

	c1 := NewCache()
	files := FilesInFolder(".")
	for _, f1 := range files {
		for _, f2 := range files {
			d := c1.Get(f1, f2)
			fmt.Printf("%s\t%s -> \t%2.6f\n", f1, f2, d)
		}
	}

	c1.Save()
	c2 := NewCache()
	if c1.Size() != c2.Size() {
		t.Fatalf("Save restore of cache failed")
	}

}
