package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func main() {
	// Step 1: creating a CSV file
	csvFile, err := os.Create("posts.csv")
	if err != nil {
		panic(err)
	}
	// At the end process, we will close csvFile.
	defer csvFile.Close()

	// Step 2: Create data for writing to CVS file.
	allPosts := []Post{
		Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"},
		Post{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"},
		Post{Id: 3, Content: "Hola Mundo!", Author: "Pedro"},
		Post{Id: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"},
	}

	// Step 3: Write to CSV file.
	// Create a CSV Writer object.
	writer := csv.NewWriter(csvFile)
	// Loop in all posts
	for _, post := range allPosts {
		// Convert struct value member to string.
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		// Write a line to CSV file.
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	// Flush all data
	writer.Flush()

	// Step 4: reading a CSV file
	file, err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}
	// At the end of process, we will close CSV file.
	defer file.Close()

	// Step 5: Create a CSV Reader object
	reader := csv.NewReader(file)
	// Set option for reader ???????????????
	reader.FieldsPerRecord = -1
	// Read all line.
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// Step 6: Convert to array of Post structure.
	var posts []Post
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post{Id: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}
	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)
}
