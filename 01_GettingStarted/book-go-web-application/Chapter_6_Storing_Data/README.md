We introduced data persistence in chapter 2, briefly touching on how to persist data into a relational database, PostgreSQL.  
In this chapter we’ll delve deeper into data persistence and talk about how you can store data in memory, files, relational databases, and No SQL databases.  
Data persistence is technically not part of web application programming, but it’s often considered the third pillar of any web application—the other two pillars are templates and handlers.  
This is because most web applications need to store data in one form or another.  
I’m generalizing but here are the places where you can store data:
* In memory (while the program is running)
* In files on the filesystem
* In a database, fronted by a server program

To get details contents, please refer to chapter 6 at book go-web-application. 
This tutorial will help you answer question below:
# 6.1 In-memory storage
* [What is In-memory storage?](#In-memory-storage)
* [How to implement a API to get a struct data from a member in Golang?](#In-memory-storage)
* [Suppose that if we fill a data with the same ID, it will happen a error on our program?](Issue-1)



## In-memory-storage
In-memory storage refers not to storing data in in-memory databases but in the running application itself, to be used while the application is running. In-memory data is usually stored in data structures, and for Go, this primarily means with arrays, slices, maps, and most importantly, structs.
1. Design a Post struct that represents a post in a forum application.
```
type Post struct {
	Id      int
	Content string
	Author  string
}
```

2. Design two prototype to getting a post.  
* Access post by Id.
* Access post by Auther.
```
var PostById map[int]*Post
var PostsByAuthor map[string][]*Post
```
You have two variables: 
* PostById maps the unique ID to a pointer to a post
* PostsByAuthor maps the author’s name to a slice of pointers to posts.
More detaisl: [here](https://github.com/huavanthong/MasterGolang/tree/main/01_GettingStarted/book-go-web-application/Chapter_6_Storing_Data/map_store)

### Output
```
&{1 Hello World! Sau Sheong}
&{2 Bonjour Monde! Pierre}
&{1 Hello World! Sau Sheong}
&{4 Greetings Earthlings! Sau Sheong}
&{3 Hola Mundo! Pedro}
```
### Issue-1
What is happened if we fill a two data with the same ID to our program?
```
&{1 Bonjour Monde! Pierre}
<nil>
&{1 Hello World! Sau Sheong}
&{4 Greetings Earthlings! Sau Sheong}
&{3 Hola Mundo! Pedro}
```
How we can depend the issue on the above case?

## File-storage

