# Introduction
We introduced data persistence in chapter 2, briefly touching on how to persist data into a relational database, PostgreSQL.  
In this chapter 6, we’ll delve deeper into data persistence and talk about how you can store data in memory, files, relational databases, and No SQL databases.  
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

# 6.2 File storage
* [What is disadvantage for using in-memory? Why we need File storage?](#File-storage)
* [What is CVS? What things can we do with it?](#CVS)
* [What is Gob package? What things can we do with it?](#Gob)
* [How many ways to read/write a file in Golang?](#Ways-to-read/write-a-file)
* [Demo a example to read and write to a file](#Reading-and-writing-to-a-file)
* [What is the difference between CSV file and using Gob package?](#Gob-package)

# 6.3 Go and SQL
* [What is the commands PostgreSQL for setting Database?](#Setting-up-the-database)
* [Are you understand about work-flow to access database on PostgresSQL yet?](#work-flow)
* [How do you get the database driver?](#Register-the-database-driver)
* [What is work flow running a database driver?](#work-flow-1)

# 6.4 Go and SQL relationships
* [How many basic relationships in Database?](#Go-and-SQL-relationships)
* [How are you setting SQL with realtions](#Setting-up-the-database-with-realtionships)

# 6.5 Go relational mappers

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
* PostById maps the unique ID to a pointer to a post.
* PostsByAuthor maps the author’s name to a slice of pointers to posts.  
More detaisl: [map_store](https://github.com/huavanthong/MasterGolang/tree/main/01_GettingStarted/book-go-web-application/Chapter_6_Storing_Data/map_store)

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
Storing in memory is fast and immediate because there’s no retrieval from disk. But there’s one very important drawback: in-memory data isn’t actually persistent.  
There are a number of ways data can be persisted, but the most common method is to store it to some sort of nonvolatile storage such as a hard disk or flash memory.  
Specifically we’ll explore two ways of storing data to files in Go:
1. Text format file.
2. CVS(comma-seperated values): 
### CSV
CSV is a common file format that’s used for transferring data from the user to the system.  
* It can be quite useful when you need to ask your users to provide you with a large amount of data and it’s not feasible to ask them to enter the data into your [forms](#https://www.meisternote.com/app/note/OjXGfZF7AK_C/4-2-html-forms-and-go).  
* You can ask your users to use their favorite spreadsheet, enter all their data, and then save it as CSV and upload it to your web application.  
* Once you have the file,File storage you can decode the data for your purposes.
* Similarly, you can allow your users to get their data by creating a CSV file out of their data and sending it to them from your web application.
### Gob
Gob is a binary format that can be saved in a file, providing a quick and effective means of serializing in-memory data to one or more files.

### Ways-to-read/write-a-file
#### Using WriteFile and ReadFile
To write to file 
```
    err := ioutil.WriteFile("data1", data, 0644)
```

To read from file
```
    read1, _ := ioutil.ReadFile("data1")
```
#### Using File struct
To create a file
```
    file1, _ := os.Create("data2")
    defer file1.Close()
```

To write to file
```
    bytes, _ := file1.Write(data)
```

To open a file
```
    file2, _ := os.Open("data2")
	defer file2.Close()
```

To read from file
```
	read2 := make([]byte, len(data))
	bytes, _ = file2.Read(read2)
```
More detaisl: [read_write_files](https://github.com/huavanthong/MasterGolang/tree/main/01_GettingStarted/book-go-web-application/Chapter_6_Storing_Data/read_write_files)
### Reading-and-writing-to-a-file
To create a CSV file.
```
    csvFile, err := os.Create("posts.csv")	
    if err != nil {
		panic(err)
	}
	defer csvFile.Close()
```

To write to a CSV file. Using [encoding/csv](https://pkg.go.dev/encoding/csv) package
```
	writer := csv.NewWriter(csvFile)
	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()
```

To open a CSV file.
```
    file, err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
```

To read a CSV file.
```
    reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
```

To contain the content CSV file to the buffer.
```
	var posts []Post
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post{Id: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}
```
More details: [csv_store](https://github.com/huavanthong/MasterGolang/tree/main/01_GettingStarted/book-go-web-application/Chapter_6_Storing_Data/csv_store)
### Gob-package
The encoding/gob package manages streams of gobs, which are binary data, exchanged between an encoder and a decoder. It’s designed for serialization and transporting data but it can also be used for persisting data.  
Encoders and decoders wrap around writers and readers, which conveniently allows you to use them to write to and read from files.  
  
To store data by encoding using Gob.
```
func store(data interface{}, filename string) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
}
```

To load data by decoding using Gob.
```
func load(data interface{}, filename string) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(data)
	if err != nil {
		panic(err)
	}
}
```
More details: [gob_store](https://github.com/huavanthong/MasterGolang/tree/main/01_GettingStarted/book-go-web-application/Chapter_6_Storing_Data/gob_store)


## Go-and-SQL
More details: [sql_store1](https://github.com/huavanthong/MasterGolang/tree/main/01_GettingStarted/book-go-web-application/Chapter_6_Storing_Data/sql_store1)

### Setting-up-the-database:
Once you’ve created the database, you’ll follow these steps:
1. Create the database user.
2. Create the database for the user.
3. Run the setup script that’ll create the table that you need.
#### To create a Postgres database user called gwp
- option -P: tells the createuser program to prompt you for a password for gwp.
- option -d: tells the program to allow gwp to create databases.
```
    createuser -P -d gwp
```
#### To create a database named gwp
```
    createdb gwp
```
#### To run script sql
- Create a file setup.sql
```
    psql -U gwp -f setup.sql -d gwp
```
### Connecting to the database
To connect to the database
```
var Db *sql.DB

// connect to the Db
func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}
```
#### Work-flow
1. The sql.DB struct is a handle to the database and represents a pool of zero or database connections that’s maintained by the **sql** package.  
2. Setting up the connection to the database is a one-liner using the **Open** function.  
    - Passing in the database driver name (in our case, it’s postgres).
    - And a data source name.
        - The data source name is a string that’s specific to the database drive and tells the driver how to connect to the database. 
3. The **Open** function then returns a pointer to a **sql.DB struct**.  
**Note:** 
    - the Open function doesn’t connect to the database or even validate the parameters yet—it simply sets up the necessary structs for connection to the database late.
    - The connection will be set up lazily when it’s needed.
    - **sql.DB** doesn’t needed to be closed (you can do so if you like); **it’s simply a handle and not the actual connection**. Remember that this abstraction contains a pool of database connections and will maintain them.
### Register-the-database-driver 
To register database driver
```
    sql.Register("postgres", &drv{})
```
#### Work-flow

### Usage on Database
#### Create a post
Implement a Create() function. 
```
func (post *Post) Create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}
```
To use it
```
	post := Post{Content: "Hello World!", Author: "Sau Sheong"}
    post.Create()
```
#### Retrieving a post
Implement a GetPost() function.
```
// Get a single post
func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}
```
To use it
```
	readPost, _ := GetPost(post.Id)
```
#### Updating a post
Implement a Update() function.
```
// Update a post
func (post *Post) Update() (err error) {
	_, err = Db.Exec("update posts set content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author)
	return
}
```
To use it
```
	// Update the post
	readPost.Content = "Bonjour Monde!"
	readPost.Author = "Pierre"
	readPost.Update()
```
#### Deleting a post
Implement a Update() function.
```
// Delete a post
func (post *Post) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}
```
To use it
```
	// Delete the post
	readPost.Delete()
```
#### Getting all posts
Implement a Update() function.
```
// get all posts
func Posts(limit int) (posts []Post, err error) {
	rows, err := Db.Query("select id, content, author from posts limit $1", limit)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}
```
To use it
```
	// Get all posts
	posts, _ := Posts(10)
```
## Go-and-SQL-relationships
One of the reasons relational databases are so popular for storing data is because tables can be related. This allows pieces of data to be related to each other in a consistent and easy-to-understand way.  
There are essentially four ways of relating a record to other records.
* **One to one** (has one)—A user has one profile, for example.
* **One to many** (has many)—A user has many forum posts, for example.
* **Many to one** (belongs to)—Many forum posts belong to one user, for example.
* **Many to many** —A user can participate in many forum threads, while a forum thread can also have many users contributing to it, for example.\

### Setting-up-the-database-with-realtionships

### One-to-many relationships