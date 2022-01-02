package main

import "fmt"
import "net/http"

func main() {
	fmt.Println("Launching server...")
	http.ListenAndServe(":12015",
		http.FileServer(http.Dir(".")))
	// func ListenAndServe(addr string, handler Handler) error
	// func FileServer(root FileSystem) Handler
	// Dir is a data type defined in the http package. It stores a directory in the file system as a string
	// built-in server will look from the server’s home directory for the file being requested
}
