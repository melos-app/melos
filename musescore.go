package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func musescoreRemoveTitle(fileName string) {
	// Read XML file
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	// Parse XML into a document
	var doc interface{}
	err = xml.Unmarshal(data, &doc)
	if err != nil {
		panic(err)
	}

	// Remove elements with style='title'
	removeElements(doc)

	// Marshal back to XML
	newData, err := xml.MarshalIndent(doc, "", "  ")
	if err != nil {
		panic(err)
	}

	// Write back to file
	err = os.WriteFile(fileName, newData, 0644)
	if err != nil {
		panic(err)
	}
}

func removeElements(node interface{}) {
	switch v := node.(type) {
	case map[string]interface{}:
		if style, ok := v["style"]; ok && style == "title" {
			delete(v, "Text")
		}
		for _, value := range v {
			removeElements(value)
		}
	case []interface{}:
		for _, item := range v {
			removeElements(item)
		}
	}
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
			if err != nil {
				return fmt.Errorf("Error running mscore: %w\n", err)
			}
		}
	}

	return nil
}
