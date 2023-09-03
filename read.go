package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {
	fmt.Println("Enter the path of file: ")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	namesfile := scan.Text()

	file, _ := os.Open(namesfile)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())

	}

	file.Close()

	nameSlice := make([]name, 0, len(txtlines))

	for _, v := range txtlines {
		words := strings.Fields(v)
		fmt.Println(words[0], words[1])
		nameSlice = append(nameSlice, name{words[0], words[1]})
	}

	fmt.Println()
	for _, j := range nameSlice {
		fmt.Println(j.fname, j.lname)
	}

}
