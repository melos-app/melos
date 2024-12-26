package main

import (
	"log"
	"os"

	"github.com/melos-app/melos"
)

func main() {

	if len(os.Args) < 2 {
		log.Println("Please provide the path to the typst file")
		os.Exit(-1)
	}

	melos.UpdateSongsInTypstFile(os.Args[1])
}
