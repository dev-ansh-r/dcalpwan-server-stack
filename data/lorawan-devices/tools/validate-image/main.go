package main

import (
	"bytes"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func testPNG(path string) bool {
	// Read image
	finput, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	input, err := ioutil.ReadAll(finput)
	if err != nil {
		panic(err)
	}
	in := bytes.NewReader(input)
	_, err = png.Decode(in)
	finput.Close()
	if err != nil {
		return false
	}

	return true
}

func testJPEG(path string) bool {
	// Read image
	finput, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	input, err := ioutil.ReadAll(finput)
	if err != nil {
		panic(err)
	}
	in := bytes.NewReader(input)
	_, err = jpeg.Decode(in)
	finput.Close()
	if err != nil {
		return false
	}

	return true
}

func main() {
	args := os.Args[1:]
	file := args[0]

	if _, err := os.Stat(file); os.IsNotExist(err) {
		log.Fatalf("file %v does not exist", file)
	}

	ext := strings.ToLower(filepath.Ext(file))
	if ext == ".png" {
		success := testPNG(file)
		if !success {
			log.Fatalf("png.Decode failed: Unknown file format %v", file)
		}
	}
	if ext == ".jpg" || ext == ".jpeg" {
		success := testJPEG(file)
		if !success {
			log.Fatalf("jpeg.Decode failed: Unknown file format %v", file)
		}
	}
}
