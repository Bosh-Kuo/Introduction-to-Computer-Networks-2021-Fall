package main

import (
	"bufio"
	"fmt"
	"net"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":12015") // listener socket
	conn, _ := ln.Accept()               // data socket
	defer ln.Close()
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	message := ""
	if scanner.Scan() {
		message = scanner.Text() // 不含換行符號
		fmt.Println(message)
	}

	writer := bufio.NewWriter(conn)
	newline := fmt.Sprintf("%d bytes received\n", len(message))
	_, errw := writer.WriteString(newline)
	check(errw)
	writer.Flush()

}
