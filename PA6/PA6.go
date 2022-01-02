package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)

	// Read file size(string)
	size, errr := reader.ReadString('\n')  
	check(errr)

	// turn string to int
	var l = len(size)
	size = size[:l-1]
	file_size, _ := strconv.Atoi(size)  

	f_out, _ := os.Create("whatever.txt")
	defer f_out.Close()

	count := 1
	readlen := 0
	for readlen < file_size {
		// reads from the socket one line at a time
		line, _ := reader.ReadString('\n')
		readlen += len(line)
		fmt.Fprintf(f_out, "%d ", count)
		fmt.Fprintf(f_out, "%s", line)
		count++
	}

	// get file information
	fi, _ := os.Stat("whatever.txt")
	fmt.Printf("Upload file size: %d\nOutput file size: %d\n", file_size, fi.Size())

	// send a message back that tells the client the original file and the new file size
	writer := bufio.NewWriter(conn)
	newline := fmt.Sprintf("Original file size: %d, New file size: %d\n", file_size, fi.Size())
	_, errw := writer.WriteString(newline)
	check(errw)
	writer.Flush()

	// Add a 5 second delay
	time.Sleep(5 * time.Second)
}

func main() {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":12015")
	defer ln.Close()
	for {
		conn, _ := ln.Accept()
		defer conn.Close()

		// make conn concurrent
		go handleConnection(conn)
	}
}
