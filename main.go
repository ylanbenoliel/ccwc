package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var fileName = flag.String("f", "", "File to be readed")
	var bytesCountFlag = flag.Bool("c", false, "Bytes count")
	var linesCountFlag = flag.Bool("l", false, "Lines count")
	var wordsCountFlag = flag.Bool("w", false, "Words count")
	var charactersCountFlag = flag.Bool("m", false, "Characters count")

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

	if *linesCountFlag {
		var lines int
		fileScanner := bufio.NewScanner(file)
		for fileScanner.Scan() {
			lines++
		}
		output += fmt.Sprintf("%d", lines)
	}

	if *wordsCountFlag {
		var wordsCount int
		var text string
		fileScanner := bufio.NewScanner(file)
		for fileScanner.Scan() {
			text = fileScanner.Text()
			text = strings.Trim(text, " ")
			if text == "" {
				continue
			}
			wordsCount += len(strings.Fields(text))
		}
		output += fmt.Sprintf("%d", wordsCount)
	}

	if *charactersCountFlag {
		var charCount int
		reader := bufio.NewReader(file)
		for {
			r, _, err := reader.ReadRune()
			if err != nil {
				break
			}
			if r != utf8.RuneError {
				charCount++
			}
		}
		output += fmt.Sprintf("%d", charCount)
	}

	output += " " + *fileName
	fmt.Println(output)
}
