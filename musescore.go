package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/antchfx/xmlquery"
)

func musescoreRemoveTitles(dir string) error {
	files, err := filepath.Glob(dir + "/*/*.mscx")

	if err != nil {
		return err
	}

	for _, f := range files {
		err := musescoreRemoveTitleFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}

func musescoreRemoveTitleFile(fileName string) error {
	// Read XML file
	log.Println("Remove title: ", fileName)
	data, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	newData, err := removeTextWithTitleStyle(data)

	// Write back to file
	err = os.WriteFile(fileName, newData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func removeTextWithTitleStyle(input []byte) ([]byte, error) {
	// Parse the XML document
	doc, err := xmlquery.Parse(bytes.NewReader(input))
	if err != nil {
		return nil, err
	}

	// Find all Text elements with style='title'
	nodes := xmlquery.Find(doc, "//Text[style='title']")

	// Remove each matching node from the tree
	for _, node := range nodes {
		xmlquery.RemoveFromTree(node)
	}

	// Create a buffer to store the output
	var output bytes.Buffer

	// Write the modified XML to the buffer
	doc.WriteWithOptions(&output, xmlquery.WithOutputSelf())

	return output.Bytes(), nil
}

func musescoreUncompress(musescoreDir, musescoreXDir string) error {
	os.RemoveAll(musescoreXDir)
	err := os.Mkdir(musescoreXDir, 0755)
	if err != nil {
		return err
	}

	files, err := os.ReadDir(musescoreDir)
	if err != nil {
		return err
	}

	for _, f := range files {
		if !f.IsDir() && strings.ToLower(filepath.Ext(f.Name())) == ".mscz" {
			var i int
			for i = 0; i < 3; i++ {
				log.Println("Uncompressing: ", f.Name())
				filePath := filepath.Join(musescoreDir, f.Name())
				base := strings.TrimSuffix(f.Name(), filepath.Ext(f.Name()))
				songXDir := filepath.Join(musescoreXDir, base)
				// Create directory for the file
				err = os.MkdirAll(songXDir, 0755)
				if err != nil {
					return fmt.Errorf("Error creating directory: %w", err)
				}

				// Run mscore command
				outputPath := filepath.Join(songXDir, base+".mscx")
				cmd := exec.Command("mscore", filePath, "--export-to", outputPath)
				err = cmd.Run()
				if err == nil {
					break
				}

				log.Printf("Error running mscore: %v, trying again\n", err)
			}

			if i == 3 {
				return fmt.Errorf("Error running mscore: %v, failed 3 times, aborting\n", err)
			}
		}
	}

	return nil
}

func musescoreGenerateSvg(musescoreXDir, svgDir string) error {
	fmt.Println("xdir", musescoreXDir)
	os.RemoveAll(svgDir)
	err := os.Mkdir(svgDir, 0755)
	if err != nil {
		return err
	}

	files, err := filepath.Glob(musescoreXDir + "/*/*.mscx")

	if err != nil {
		return err
	}

	for _, f := range files {
		log.Println("Exporting to SVG: ", f)
		base := strings.TrimSuffix(filepath.Base(f), filepath.Ext(f))
		// Create directory for the file
		if err != nil {
			return fmt.Errorf("Error creating directory: %w", err)
		}

		// Run mscore command
		// mscore "$file" --export-to "svg/${base}.svg" -T 0
		outputPath := filepath.Join(svgDir, base+".svg")
		cmd := exec.Command("mscore", f, "--export-to", outputPath, "-T", "0")
		err = cmd.Run()
		if err == nil {
			break
		}
	}

	return nil
}
