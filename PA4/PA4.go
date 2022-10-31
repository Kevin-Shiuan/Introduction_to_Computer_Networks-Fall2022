package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":12010")
	conn, _ := ln.Accept()
	defer ln.Close()
	defer conn.Close()
	// read from the socket first the file size
	reader := bufio.NewReader(conn)
	input, errr := reader.ReadString('\n')
	check(errr)

	// remove '\n' from reading input & convert into int
	value := strings.TrimSpace(input)
	len, err := strconv.Atoi(value)

	// prepare whatever.txt
	output, err := os.Create("whatever.txt")
	check(err)
	defer output.Close()
	writer := bufio.NewWriter(output)

	// reads from the socket one line at a time & store the each line into whatever.txt
	scanner := bufio.NewScanner(conn)
	outputSize := 0
	scannedSize := 0
	counter := 1
	for scanner.Scan() {
		len3, _ := writer.WriteString(fmt.Sprintf("%d ", counter))
		len1, _ := writer.WriteString(scanner.Text())
		len2, _ := writer.WriteString("\n")
		outputSize += len1 + len2 + len3
		scannedSize += len1 + len2
		writer.Flush()
		counter++
		if scannedSize >= len {
			break
		}

	}
	// fmt.Printf("%d bytes is told, %d bytes written, %d bytes file generated\n", len, scannedSize, outputSize)
	fmt.Printf("Upload file size: %d\n", len)
	fmt.Printf("Output file size: %d\n", outputSize)

	writer2 := bufio.NewWriter(conn)
	newline := fmt.Sprintf("%d bytes received, %d bytes file generated", scannedSize, outputSize)
	_, errw := writer2.WriteString(newline)
	check(errw)
	writer2.Flush()
}
