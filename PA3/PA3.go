package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fileName := ""

	// prompts the user for the upload filename
	fmt.Printf("Input filename:")
	fmt.Scanf("%s", &fileName)

	// connects to the server that polly implements and runs already on the workstation at port 12000
	conn, errc := net.Dial("tcp", "140.112.42.221:12000")
	check(errc)
	defer conn.Close()
	writer := bufio.NewWriter(conn)

	// Open the file
	f_in, erro := os.Open(fileName)
	check(erro)
	defer f_in.Close()
	scnner := bufio.NewScanner(f_in)

	// get file information
	fi, _ := os.Stat(fileName)

	// sends first the file size
	fmt.Printf("Send the file size first: %d\n", fi.Size())
	fileSize := fmt.Sprintf("%d\n",fi.Size())
	writer.WriteString(fileSize)


	// Read the content and send it to conn
	for scnner.Scan() {
		text := fmt.Sprintf("%s\n", scnner.Text())
		_, errw := writer.WriteString(text)
		check(errw)
	}
	writer.Flush()

	// prints what the server says
	scanner := bufio.NewScanner(conn)
	if scanner.Scan() {
		fmt.Printf("Server says: %s bytes received\n", scanner.Text())
	}
}
