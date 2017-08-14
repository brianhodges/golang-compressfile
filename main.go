package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Println("Error: ", err)
	}
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("You must enter path parameter.")
		fmt.Println("Ex. go run main.go C:\\path\\to\\directory")
		os.Exit(1)
	}
	err := Zip(os.Args[1])
	check(err)
}

func Zip(filepath string) error {
	pathPieces := strings.Split(filepath, "\\")
	filename := pathPieces[len(pathPieces)-1]
	filePieces := strings.Split(filename, ".")
	output := filePieces[0] + ".zip"

	newfile, err := os.Create(output)
	check(err)
	defer newfile.Close()

	zipWriter := zip.NewWriter(newfile)
	defer zipWriter.Close()

	zipfile, err := os.Open(filepath)
	check(err)
	defer zipfile.Close()

	info, err := zipfile.Stat()
	check(err)

	header, err := zip.FileInfoHeader(info)
	check(err)
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	check(err)

	_, err = io.Copy(writer, zipfile)
	check(err)

	fmt.Println("Compressed: " + output)
	return nil
}
