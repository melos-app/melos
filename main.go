package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

type song struct {
	name   string
	images []string
}

func (s song) String() string {
	r := s.name + ": "
	for _, i := range s.images {
		r += i + " "
	}
	return r
}

type songs map[string]*song

func (s songs) String() string {
	r := ""
	for _, v := range s {
		r += v.String() + "\n"
	}
	return r
}

func newSongs() songs {
	return songs(make(map[string]*song))
}

var reSong = regexp.MustCompile(`^(.+)-(\d+)\.svg$`)

func (s *songs) addFile(f string) error {
	matches := reSong.FindStringSubmatch(f)

	if len(matches) >= 3 {
		basename := matches[1]
		//indexS := matches[2]
		basename = strings.Replace(basename, "_", " ", -1)
		// TODO: add file to Songs
		s_, ok := (*s)[basename]
		if !ok {
			(*s)[basename] = &song{name: basename, images: []string{f}}
		} else {
			s_.images = append(s_.images, f)
		}
	} else {
		return fmt.Errorf("No match found")
	}

	return nil
}

type typ struct {
	fileName string
	songs    songs
}

func newTyp(fname string) *typ {
	return &typ{
		fileName: fname,
		songs:    songs(make(map[string]*song)),
	}
}

func (t *typ) scan() error {
	file, err := os.Open(t.fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	titleRegex := regexp.MustCompile(`#song\("(.+?)"`)
	imageRegex := regexp.MustCompile(`image\("(.+?)"`)

	var currentSong song
	var inSong bool

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if titleMatch := titleRegex.FindStringSubmatch(line); len(titleMatch) > 1 {
			title := titleMatch[1]
			if inSong {
				// need to make a copy here, or all will reference one instance
				cs := currentSong
				t.songs[title] = &cs
			}
			currentSong = song{name: title, images: []string{}}
			inSong = true
		}

		if imageMatches := imageRegex.FindAllStringSubmatch(line, -1); len(imageMatches) > 0 {
			for _, match := range imageMatches {
				if len(match) > 1 {
					currentSong.images = append(currentSong.images, match[1])
				}
			}
		}

		if line == ")" && inSong {
			t.songs[currentSong.name] = &currentSong
			inSong = false
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func (t *typ) update(s songs) error {
	f, err := os.OpenFile(t.fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	// loop through songs and add any that are not already in typ file
	for k, v := range s {
		if _, ok := t.songs[k]; !ok {
			log.Println("Adding song: ", k)
			_, err = f.WriteString(fmt.Sprintf("#song(\"%s\",\n", k))
			if err != nil {
				return err
			}
			for i, image := range v.images {
				is := fmt.Sprintf("image(\"svg/%s\"),\n", image)
				if i == 0 {
					is = "  (" + is
				} else {
					is = "   " + is
				}

				if i == len(v.images)-1 {
					is += ")"
				}

				_, err = f.WriteString(is)
				if err != nil {
					return err
				}
			}
			_, err = f.WriteString(")\n")
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func updateSongsInTypstFile(fName string) {

	log.Println("Update list of songs")
	files, err := os.ReadDir("svg")
	if err != nil {
		log.Fatal(err)
	}

	s := newSongs()

	for _, file := range files {
		err := s.addFile(file.Name())
		if err != nil {
			log.Printf("Error adding file %f: %f", file.Name(), err)
			os.Exit(-1)
		}
	}

	t := newTyp(fName)

	err = t.scan()
	if err != nil {
		log.Printf("Error scanning typ file: %v", err)
		os.Exit(-1)
	}

	err = t.update(s)

	if err != nil {
		log.Printf("Error updating typ file: %v", err)
	}
}

func removeTitle(fileName string) {
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

func main() {
	if len(os.Args) < 2 {
		log.Println("Usage: ", os.Args[0], "typst_file")
		os.Exit(-1)
	}

	typstFileName := os.Args[1]

	updateSongsInTypstFile(typstFileName)
}
