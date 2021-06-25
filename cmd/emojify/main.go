package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	var (
		printListFlag bool
	)
	flag.BoolVar(&printListFlag, "list", false, "print list")
	flag.BoolVar(&printListFlag, "l", false, "print list(short version)")
	flag.Parse()

	if printListFlag {
		PrintList()
		return
	}

	line, readOK := readLine()
	for readOK {
		fmt.Println(ReplaceEmoji(line))
		line, readOK = readLine()
	}
}

// PrintList print emoji list
func PrintList() {
	for key, val := range GetEmojiMap() {
		fmt.Print(key)
		fmt.Print(" ")
		fmt.Println(val)
	}
}

var rdr = bufio.NewReaderSize(os.Stdin, 1000000)

func readLine() (string, bool) {
	buf := make([]byte, 0, 1000000)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil {
			return "", false
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf), true
}
