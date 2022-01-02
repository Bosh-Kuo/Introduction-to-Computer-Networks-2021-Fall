package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Who's there?\n")
	text := ""
	text2 := ""

	fmt.Scanf("%s %s", &text, &text2)

	fmt.Printf("Hello, %s\n", text)
	fmt.Println("Hello,", text)
	fmt.Fprintf(os.Stdout, "Hello, %s\n", text)

}


