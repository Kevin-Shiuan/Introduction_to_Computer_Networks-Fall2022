package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"os"
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
	defer ln.Close()
	// conn, _ := ln.Accept()
	// defer conn.Close()
	for {
		conn, _ := ln.Accept()
		defer conn.Close()

		go handleConnection(conn)
	}
	// handleConnection(conn)
}

func handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	req, err := http.ReadRequest(reader)
	check(err)
	// fmt.Println(req.RequestURI)

	fi, errr := os.Stat(strings.Trim(req.RequestURI, "/"))
	if os.IsNotExist(errr) {
		fmt.Println("File not found")
	} else {
		fmt.Printf("File size: %d\n", fi.Size())
	}

	conn.Close()
}
