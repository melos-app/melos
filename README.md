# Melos

Melos is a to to create song books from a directory of Musescore files.

## Requirements

- install [Musescore](https://musescore.com/)
- install [Typst](https://typst.app/)
- Musescore files are in a directory where the file name is the title of the
  song.

## Process

This program runs the following steps:

1. uncompress the musescore files into the `musescorex` directory
1. removes the Title from the musescore file
1. Generate SVG file in the `svg` directory
1. create a book using the provided typst file and directory of SVGs

## Using Melos
