package main

import (
	"fmt"
	"net/http"
	"os"
)

type appHandler func(http.ResponseWriter, *http.Request) int

func myHandler(w http.ResponseWriter, r *http.Request) int {
	reqFilePath := fmt.Sprintf(".%s", r.URL.Path)
	if reqFilePath[len(reqFilePath)-1] == '/' {
		return http.StatusNotFound
	}
	if _, err := os.Stat(reqFilePath); err == nil {
		http.ServeFile(w, r, reqFilePath)
		return http.StatusOK
	} else {
		return http.StatusInternalServerError
	}
}

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if stat := fn(w, r); stat != http.StatusOK {
		fmt.Fprintln(w, "File not found")
	}
}

func main() {
	fmt.Println("Launching server...")
	http.Handle("/", appHandler(myHandler))
	http.ListenAndServeTLS(":12015", "server.cer", "server.key", nil)
}
