package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	//promts for input
	fmt.Printf("Input filename: ")
	inputFileName := ""
	fmt.Scanf("%s\n", &inputFileName)
	//prompts for output
	fmt.Printf("Output filename: ")
	outputFileName := ""
	fmt.Scanf("%s\n", &outputFileName)
	f1, err := os.Open(inputFileName)
	check(err)
	f2, err := os.Create(outputFileName)
	check(err)
	defer f2.Close()
	scanner := bufio.NewScanner(f1)
	writer := bufio.NewWriter(f2)
	counter := 1
	//reads from the input file one line at a time
	for scanner.Scan() {
		//prepends the line count to each line
		writer.WriteString(fmt.Sprintf("%d ", counter))
		writer.WriteString(scanner.Text())
		writer.WriteString("\n")
		writer.Flush()
		counter++
	}
}
