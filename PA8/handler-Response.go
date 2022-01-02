package main

import "fmt"
import "net/http"

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
func main() {
	fmt.Println("Launching server...")
	hh := http.HandlerFunc(helloHandler)  // adapting a programmer-defined function to a handler function
	http.Handle("/hello", hh)  // associates a prefix to its handler, (handle "/hello" command with hh handler)
	fs := http.FileServer(http.Dir("."))  // It's also a handler
	http.Handle("/", http.StripPrefix("/", fs))
	http.ListenAndServe(":12015", nil)  // set no handler
}
