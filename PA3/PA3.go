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

	// conn, errc := net.Dial("tcp", "140.112.42.221:12000")
	conn, errc := net.Dial("tcp", ":12010")
	// conn, errc := net.Dial("tcp", ":11999")
	check(errc)
	defer conn.Close()

	//promts for input
	fmt.Printf("Input filename: ")
	inputFileName := ""
	fmt.Scanf("%s\n", &inputFileName)

	//create a writer
	writer := bufio.NewWriter(conn)

	//sends first the file size
	f, err := os.Stat(inputFileName)
	check(err)

	//writer.WriteString(string(f.Size()))
	writer.WriteString(fmt.Sprintf("%d\n", f.Size()))
	writer.Flush()

	//sends the file content
	f1, err := os.Open(inputFileName)
	check(err)
	scanner := bufio.NewScanner(f1)
	for scanner.Scan() {
		writer.WriteString(scanner.Text())
		writer.WriteString("\n")
		writer.Flush()
	}
	scanner2 := bufio.NewScanner(conn)
	if scanner2.Scan() {
		fmt.Printf("Server replies: %s\n", scanner2.Text())
	}
}
