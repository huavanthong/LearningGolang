package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

// store data
func store(data interface{}, filename string) {
	buffer := new(bytes.Buffer)
	// Create a encoder object
	encoder := gob.NewEncoder(buffer)
	// Encoding data
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	// Store in file with binary format.
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
}

// load the data
func load(data interface{}, filename string) {
	// Read file at a bianry format.
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	// Create a buffer to contain data
	buffer := bytes.NewBuffer(raw)
	// Create a Decoder object
	dec := gob.NewDecoder(buffer)
	// Decode data
	err = dec.Decode(data)
	if err != nil {
		panic(err)
	}
}

func main() {
	post := Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"}
	store(post, "post1")
	var postRead Post
	load(&postRead, "post1")
	fmt.Println(postRead)
}
