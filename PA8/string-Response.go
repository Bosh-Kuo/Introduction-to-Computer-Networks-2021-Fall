package main

import "fmt"
import "bufio"
import "net"
import "net/http"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":12015")
	defer ln.Close()
	conn, _ := ln.Accept()
	defer conn.Close()

	reader := bufio.NewReader(conn)
	req, err := http.ReadRequest(reader)  // 回傳 type Request 
	check(err)
	fmt.Printf("Method: %s\n", req.Method)

	// Header
	fmt.Fprintf(conn, "HTTP/1.1 404 Not Found\r\n")  // status line
	fmt.Fprintf(conn, "Date: ...\r\n")  //  lame header line
	fmt.Fprintf(conn, "\r\n")  // empty line that signals the end of the header lines

	// Content
	fmt.Fprintf(conn, "File not found\r\n")
	fmt.Fprintf(conn, "笑鼠\r\n")
	fmt.Fprintf(conn, "\r\n")
}
