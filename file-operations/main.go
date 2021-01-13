package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

// A program to remove comments and empty lines from Javascript file (sample.js)

func main() {

	file, err := ioutil.ReadFile("./sample.js")

	if err != nil {
		fmt.Println("Cannot read file", err)
	}

	re, err := regexp.Compile(`\/\*(.|\n)*\*\/|^\/\/.*`)

	if err != nil {
		fmt.Println("Error in Regex", err)
	}

	f_str := string(file)
	f_str = re.ReplaceAllString(f_str, "")
	f_str = regexp.
		MustCompile(`[\n]+`).
		ReplaceAllString(strings.TrimSpace(f_str), "\n")

	f, _ := os.Create("out.js")
	defer f.Close()
	f.WriteString(f_str)

}
