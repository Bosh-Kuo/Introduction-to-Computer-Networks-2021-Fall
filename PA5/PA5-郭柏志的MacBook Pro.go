package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":12015")
	defer ln.Close()
	for {
		conn, _ := ln.Accept()
		defer conn.Close()

		reader := bufio.NewReader(conn)
		size, errr := reader.ReadString('\n')
		check(errr)

		var l = len(size)
		size = size[:l-1]
		file_size, _ := strconv.Atoi(size)

		f_out, _ := os.Create("whatever.txt")
		defer f_out.Close()

		count := 1
		readlen := 0
		for readlen < file_size {
			line, _ := reader.ReadString('\n')
			readlen += len(line)
			fmt.Fprintf(f_out, "%d ", count)
			fmt.Fprintf(f_out, "%s", line)
			count++
		}

		// get file information
		fi, _ := os.Stat("whatever.txt")
		fmt.Printf("Upload file size: %d\nOutput file size: %d\n", file_size, fi.Size())
		writer := bufio.NewWriter(conn)
		newline := fmt.Sprintf("Original file size: %d, New file size: %d\n", file_size, fi.Size())
		_, errw := writer.WriteString(newline)
		check(errw)
		writer.Flush()
	}
}
