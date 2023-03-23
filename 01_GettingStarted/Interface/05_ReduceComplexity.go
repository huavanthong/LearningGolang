package main

import (
	"fmt"
	"io"
)

type BookFormat interface {
	Open(path string) (io.ReadCloser, error)
	Close(file io.ReadCloser) error
}

type PDF struct{}

func (p PDF) Open(path string) (io.ReadCloser, error) {
	// Open PDF file and return a reader
	return nil, nil
}

func (p PDF) Close(file io.ReadCloser) error {
	// Close PDF file
	return nil
}

type EPUB struct{}

func (e EPUB) Open(path string) (io.ReadCloser, error) {
	// Open EPUB file and return a reader
	return nil, nil
}

func (e EPUB) Close(file io.ReadCloser) error {
	// Close EPUB file
	return nil
}

type MOBI struct{}

func (m MOBI) Open(path string) (io.ReadCloser, error) {
	// Open MOBI file and return a reader
	return nil, nil
}

func (m MOBI) Close(file io.ReadCloser) error {
	// Close MOBI file
	return nil
}

func ReadBook(path string, format BookFormat) error {
	file, err := format.Open(path)
	if err != nil {
		return err
	}

	// Read book from file

	err = format.Close(file)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	pdf := PDF{}
	epub := EPUB{}
	mobi := MOBI{}

	// Read PDF book
	err := ReadBook("book.pdf", pdf)
	if err != nil {
		fmt.Println("Error reading book:", err)
	}

	// Read EPUB book
	err = ReadBook("book.epub", epub)
	if err != nil {
		fmt.Println("Error reading book:", err)
	}

	// Read MOBI book
	err = ReadBook("book.mobi", mobi)
	if err != nil {
		fmt.Println("Error reading book:", err)
	}
}
