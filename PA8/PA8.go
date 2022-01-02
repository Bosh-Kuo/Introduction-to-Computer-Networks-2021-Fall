package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func findFileInRoot(root string) []string {
	var files []string
	err := filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}

func customizedHandler(w http.ResponseWriter, r *http.Request) {
	// find all the files in root
	filesInRoot := findFileInRoot(".")
	reqFilePath := strings.TrimPrefix(r.URL.Path, "/")

	// In filesInRoot array, Looking for the same file name with reqFilePath
	exists := false
	for _, f := range filesInRoot {
		if f == reqFilePath || reqFilePath == "" || (reqFilePath[len(reqFilePath)-1] == '/' && f == reqFilePath[:len(reqFilePath)-1]) {
			fs := http.FileServer(http.Dir("."))
			r2 := new(http.Request)
			*r2 = *r
			r2.URL = new(url.URL)
			*r2.URL = *r.URL
			r2.URL.Path = reqFilePath
			fs.ServeHTTP(w, r2)
			exists = true
			break
		}
	}
	if !exists {
		fmt.Fprintln(w, "File Not found")
	}
}

func main() {
	fmt.Println("Launching server...")
	hd := http.HandlerFunc(customizedHandler)
	http.Handle("/", hd)
	http.ListenAndServe(":12015", nil)
}
