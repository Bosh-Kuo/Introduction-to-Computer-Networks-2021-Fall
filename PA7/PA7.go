package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
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
		//	defer conn.Close()
		reader := bufio.NewReader(conn)
		req, err := http.ReadRequest(reader)
		check(err)

		realPath := fmt.Sprintf(".%s", req.URL.Path)
		fmt.Println(realPath)
		if filestat, erro := os.Stat(realPath); erro == nil { //golang 中 if 條件式前可以搭配簡短宣告（要加 ";"）
			// the file requested exists
			fmt.Printf("file size: %d\n", filestat.Size())
		} else if errors.Is(erro, os.ErrNotExist) { // 也可以用os.IsNotExist(erro)，不用errors.Is判斷
			// file does not exist
			fmt.Println(erro)
			fmt.Printf("File not found\n")
		} else {
			// file may or may not exist, but there is something wrong
			panic(erro)
		}
		conn.Close()
	}
}
