package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func bytesCount(file *os.File) string {
	var fileInfo, err = file.Stat()
	check(err)
	return fmt.Sprintf("%d ", fileInfo.Size())
}

func linesCount(file *os.File) string {
	var lines int
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		lines++
	}
	return fmt.Sprintf("%d ", lines)
}

func wordsCount(file *os.File) string {
	var wordsCount int
	var text string
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		text = string(line[:])
		wordsCount += len(strings.Fields(text))
	}
	return fmt.Sprintf("%d ", wordsCount)
}

func charsCount(file *os.File) string {
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
	return fmt.Sprintf("%d ", charCount)
}

func fileStats(file *os.File) (int, int, int, error) {
	lineCount := 0
	wordCount := 0
	byteCount := 0
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return 0, 0, 0, err
		}
		if err == io.EOF {
			break
		}
		lineCount++
		wordCount += len(strings.Fields(line))
		byteCount += len(line)
	}
	return lineCount, wordCount, byteCount, nil
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
		output += bytesCount(file)
	}

	if *linesCountFlag {
		output += linesCount(file)
	}

	if *wordsCountFlag {
		output += wordsCount(file)
	}

	if *charactersCountFlag {
		output += charsCount(file)
	}

	if !*bytesCountFlag && !*linesCountFlag && !*wordsCountFlag {
		lines, words, bytes, err := fileStats(file)
		check(err)
		output += fmt.Sprintf("%d ", lines)
		output += fmt.Sprintf("%d ", words)
		output += fmt.Sprintf("%d ", bytes)

	}

	output += *fileName
	fmt.Println(output)
}
