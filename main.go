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
	var bytesCountFlag = flag.Bool("c", false, "Bytes count")

	flag.Parse()

	var output string

	var file, err = os.Open(*fileName)
	check(err)
	defer file.Close()

	if *bytesCountFlag {
		var fileInfo, err = file.Stat()
		check(err)
		output += fmt.Sprintf("%d", fileInfo.Size())
	}

	output += " " + *fileName
	fmt.Println(output)
}
