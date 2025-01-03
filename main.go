package main

import (
	"flag"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var fileName = flag.String("f", "", "File to be readed")

	var countBytesFlag = flag.Bool("c", false, "Bytes count")
	var output string

	flag.Parse()

	var file, err = os.ReadFile(*fileName)
	check(err)

	if *countBytesFlag {
		output += fmt.Sprintf("%d", len(file))
	}

	output += " " + *fileName
	fmt.Println(output)
}
