package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {

	if len(os.Args) < 2 {
		log.Println("Please provide the path to the typst file")
		os.Exit(-1)
	}

	typstFileName := os.Args[1]
	typstFileDir := filepath.Dir(typstFileName)

	musescoreDir := filepath.Join(typstFileDir, "musescore")
	musescoreXDir := filepath.Join(typstFileDir, "musescorex")
	// svgDir := filepath.Join(typstFileDir, "svg")

	err := musescoreUncompress(musescoreDir, musescoreXDir)
	if err != nil {
		log.Println("Error uncompressing musescore files:", err)
		os.Exit(-1)
	}

	err = musescoreRemoveTitles(musescoreXDir)
	if err != nil {
		log.Println("Error removing title:", err)
	}

	// TODO: generate svg
	updateSongsInTypstFile(typstFileName)
	// TODO: make book
}
