package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

func main() {
	file := flag.String("file", "", "Typst file to process")
	uncompress := flag.Bool("uncompress", false, "uncompress musescore files")
	removeTitles := flag.Bool("removeTitles", false, "remove titles")
	generateSvg := flag.Bool("generateSvg", false, "generate svg files")
	updateTypst := flag.Bool("updateTypst", false, "update list of songs in typst file")
	makeBook := flag.Bool("makeBook", false, "Generate PDF book file")

	if !*uncompress && !*removeTitles && !*generateSvg && !*updateTypst && !*makeBook {
		*uncompress = true
		*removeTitles = true
		*generateSvg = true
		*updateTypst = true
		*makeBook = true
	}

	if *file == "" {
		log.Println("Please provide the path to the typst file")
		os.Exit(-1)
	}

	typstFileName := os.Args[1]
	typstFileDir := filepath.Dir(typstFileName)

	musescoreDir := filepath.Join(typstFileDir, "musescore")
	musescoreXDir := filepath.Join(typstFileDir, "musescorex")
	svgDir := filepath.Join(typstFileDir, "svg")

	if *uncompress {
		err := musescoreUncompress(musescoreDir, musescoreXDir)
		if err != nil {
			log.Println("Error uncompressing musescore files:", err)
			os.Exit(-1)
		}
	}

	if *removeTitles {
		err := musescoreRemoveTitles(musescoreXDir)
		if err != nil {
			log.Println("Error removing title:", err)
		}
	}

	if *generateSvg {
		// TODO: need to finish this
		err := musescoreGenerateSvg(musescoreXDir, svgDir)
		if err != nil {
			log.Println("Error generating svg:", err)
		}
	}

	if *updateTypst {
		updateSongsInTypstFile(typstFileName)
	}

	if *makeBook {
		// TODO: make book
	}
}
