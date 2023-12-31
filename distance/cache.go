package distance

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Default. Can be changed before use, see the CLI flags in main.go.
var CACHEFILENAME string = filepath.Join(os.TempDir(), "fileDistance.cache")

// Cache for file to file distance.
// It is a very expensive calculation, since we check for word, excel, zip, etc ... files, so caching makes sense.
// We do not use filenames, but the hash of both files, to ensure propoer handling of file name or content changes.
type Cache struct {
	M map[[sha256.Size * 2]byte]float64 // should not be used directly, nor relied upon. Public only because required for ease of saving as gob.
}

// Create a new cache.
// Load from previously saved cache if there is one.
// Not thread safe.
func NewCache() *Cache {
	c := new(Cache)
	file, err := os.Open(CACHEFILENAME)
	if err != nil {
		// No file exists, just use a new empty cache
		fmt.Fprintf(os.Stderr, "%v\nno cache available at %s\n", err, CACHEFILENAME)
		c.M = make(map[[sha256.Size * 2]byte]float64)
		return c
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	if err := decoder.Decode(&c); err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stderr, "\ncache loaded from %s (%d values) \n", CACHEFILENAME, c.Size())
	return c
}

// Save cache to file.
// Not thread safe.
func (c *Cache) Save() {
	file, err := os.Create(CACHEFILENAME)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	if err := encoder.Encode(c); err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stderr, "\ncache saved to %s ( %d values)\n", CACHEFILENAME, c.Size())

}

// Clear cache in memory.
// Cache on file will be erased on next save.
func (c *Cache) Clear() {
	c.M = make(map[[sha256.Size * 2]byte]float64)
}

// Try to read from cache, if cache misses, compute, store and return result.
func (c *Cache) Get(f1, f2 string) float64 {
	idx := cacheindex(f1, f2)
	v, ok := c.M[idx]
	if ok {
		return v
	}
	// cache miss
	v = DistFile(f1, f2)
	c.M[idx] = v
	return v
}

// Number of distinct pair of files whose distance is cached.
func (c *Cache) Size() int {
	return len(c.M)
}

// index ensures that (f1, f2) and (f2,f1) will point to the same value.
// It also ensures result is the zero-value if f1 and f2 have identical contents,
// even if one is a zipped version of the other.
func cacheindex(f1 string, f2 string) (idx [sha256.Size * 2]byte) {

	d1, d2 := digest(f1), digest(f2)
	c := bytes.Compare(d1[:], d2[:])
	switch {
	case c == 0: // contents are the same, use key 0 by convention
		return idx
	case c < 0:
		for i := 0; i < sha256.Size; i++ {
			idx[i], idx[i+sha256.Size] = d1[i], d2[i]
		}
		return idx
	case c > 0:
		for i := 0; i < sha256.Size; i++ {
			idx[i], idx[i+sha256.Size] = d2[i], d1[i]
		}
		return idx
	}
	panic("internal error")
}

func digest(fname string) []byte {
	file, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		panic(err)
	}
	return hash.Sum(nil)
}
