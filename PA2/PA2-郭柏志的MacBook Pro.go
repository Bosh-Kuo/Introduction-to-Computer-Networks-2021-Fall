package main

import "fmt"
import "os"
import "bufio"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	inputFileName, outputFileName := "", ""
	count := 1
	fmt.Print("Input filename: ")
	fmt.Scanf("%s", &inputFileName)
	fmt.Print("Output filename: ")
	fmt.Scanf("%s", &outputFileName)

	// output.txt
	f_out, err := os.Create(outputFileName)
	check(err)
	defer f_out.Close()

	// input.txt
	f_in, err := os.Open(inputFileName)
	check(err)
	scnner := bufio.NewScanner(f_in)
	defer f_in.Close()

	for scnner.Scan() {
		fmt.Fprintf(f_out, "%d ", count)
		fmt.Fprintf(f_out, "%s\n", scnner.Text())
		count++
	}
}
