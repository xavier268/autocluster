package distance

import (
	"archive/zip"
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"encoding/xml"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"golang.org/x/net/html"
)

// Get (recursiveley) all files in folder
func FilesInFolder(folder string) []string {

	files := []string{}
	visit := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Print(err)
			return nil
		}
		if !info.IsDir() {
			ap, _ := filepath.Abs(path)
			files = append(files, ap)
		}
		return nil
	}
	err := filepath.Walk(folder, visit)
	if err != nil {
		log.Print(err)
	}
	return files
}

// Compute the distance matrix for all files in a folder
func ComputeFolder(folder string) *Matrix {
	return ComputeFiles(FilesInFolder(folder)...)
}

// Compute the distance matrix for a group of files
func ComputeFiles(fnames ...string) *Matrix {
	mat := new(Matrix)
	for i := 0; i < len(fnames); i++ {
		for j := i + 1; j < len(fnames); j++ {
			mat.Set(i, j, DistString(fnames[i], fnames[j]))
		}
	}
	return mat
}

// Distance between two files, given by their path names.
// Usefull text content will be extracted first.
func DistFile(f1, f2 string) float64 {

	x, y := ExtractText(f1), ExtractText(f2)
	return DistBytes(x, y)
}

func ExtractText(fname string) []byte {
	x, err := os.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	x = ungzip(x)
	x = unzlib(x)
	x = unzip(x)
	x = unxml(x)
	x = unhtml(x)
	x = trimSpaces(x)

	return x
}

func trimSpaces(src []byte) []byte {
	sp := regexp.MustCompile(`\s+`)
	return sp.ReplaceAll(src, []byte(" "))
}

// Extract byte content from zip file.
// If not zip, return original content.
func ungzip(source []byte) []byte {

	// try to unzip
	zr, err := gzip.NewReader(bytes.NewReader(source))
	if err == nil { // its a zip file, try to unzip it ...
		zr, err := io.ReadAll(zr)
		if err == nil {
			return zr // unzipped content
		}
	}
	log.Println(err)
	return source
}

// same than ungzip, but with zlib
func unzlib(source []byte) []byte {
	// try to unzip
	zr, err := zlib.NewReader(bytes.NewReader(source))
	if err == nil { // its a zlib file, try to unzip it ...
		zr, err := io.ReadAll(zr)
		if err == nil {
			return zr // unzipped content
		}
	}
	log.Println(err)
	return source
}

func unzip(source []byte) []byte {
	zr, err := zip.NewReader(bytes.NewReader(source), int64(len(source)))
	if err != nil { // not a zip archive,
		log.Println(err)
		return source
	}
	// Iterate through the files in the archive
	res := make([]byte, 0, len(source))
	for _, f := range zr.File {
		rc, err := f.Open()
		if err != nil {
			log.Println(err)
			rc.Close()
			return source
		}
		ct, err := io.ReadAll(rc)
		if err != nil {
			log.Println(err)
			rc.Close()
			return source
		}
		res = append(res, ct...)
		rc.Close()
	}
	return res

}

// If it is a valid html, returns the text content.
// If invalid html, return the source.
func unhtml(source []byte) []byte {

	doc, err := html.Parse(bytes.NewReader(source))
	if err != nil {
		log.Print(err)
		return source // not valid html !
	}

	// assume valid utf8 html5 encoded data
	text := make([]byte, 0, len(source))

	// f will recurse over the dom tree
	var f func(*html.Node)
	f = func(n *html.Node) {
		switch n.Type {
		case html.TextNode:
			text = append(text, []byte(n.Data)...)
		case html.ElementNode:
			if n.Data == "script" || n.Data == "style" || n.Data == "img" {
				return // skip text content for style, scripts, ...
			}
		}
		// if we are still there, recurse ...
		for nn := n.FirstChild; nn != nil; nn = nn.NextSibling {
			f(nn)
		}
	}
	// do it !
	f(doc)
	return text
}

// Extract text from a valid xml file.
// Will return source if invalid.
func unxml(source []byte) []byte {
	doc := new(any)
	err := xml.Unmarshal(source, doc)
	if err != nil {
		log.Println(err)
		return source
	}

	// assume valid xml now
	text := make([]byte, 0, len(source))

	dec := xml.NewDecoder(bytes.NewReader(source))
	dec.Strict = true

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			return text // all fine !
		}
		if err != nil {
			log.Println(err)
			return source // invalid xml ?
		}
		switch v := tok.(type) {
		case xml.CharData:
			text = append(text, v...)
		}
	}
}
