package main

import (
	"log"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		log.Println("Please provide the path to the typst file")
		os.Exit(-1)
	}

	updateSongsInTypstFile(os.Args[1])
}
