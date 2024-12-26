package melos

import (
	"encoding/xml"
	"io/ioutil"
)

func MusescoreRemoveTitle(fileName string) {
	// Read XML file
	data, err := ioutil.ReadFile(fileName)
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
	err = ioutil.WriteFile(fileName, newData, 0644)
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
