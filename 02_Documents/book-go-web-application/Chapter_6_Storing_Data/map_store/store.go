package main

import (
	"fmt"
)

// Step 1: Declare a Post structure to contain the data
type Post struct {
	Id      int
	Content string
	Author  string
}

// Step 2: Declare two variable with data-type is map.
// By default, map will be nil
var PostById map[int]*Post
var PostsByAuthor map[string][]*Post

// Step 4: Implement a methods to store a post struct into our maps
// 1. Get the post ID to store the post.
// 2. Get the post author to append the post.
func store(post Post) {
	PostById[post.Id] = &post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)
}

func main() {

	//Step 3: Use make() to allocate memory for our maps.
	PostById = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)

	// Step 5: Create a post data
	post1 := Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"}
	post2 := Post{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"}
	post3 := Post{Id: 3, Content: "Hola Mundo!", Author: "Pedro"}
	post4 := Post{Id: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"}

	// Step 6: Use store() function to store my data.
	store(post1)
	store(post2)
	store(post3)
	store(post4)

	// Step 7: Print the output
	fmt.Println(PostById[1])
	fmt.Println(PostById[2])

	for _, post := range PostsByAuthor["Sau Sheong"] {
		fmt.Println(post)
	}
	for _, post := range PostsByAuthor["Pedro"] {
		fmt.Println(post)
	}
}
