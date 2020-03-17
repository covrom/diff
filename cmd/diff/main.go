package main

import (
	"fmt"
	"log"
	"os"

	"github.com/covrom/diff"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: diff [src] [dst]")
		os.Exit(1)
	}

	src, err := diff.GetFileLines(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	dst, err := diff.GetFileLines(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	diff.PrintDiffSlices(src, dst)
}
