# Introduction
To get details contents, please refer to book go-web-application.  
This tutorial will help you answer question below:

# Table of contents
1. [Request Response](#request-response)
2. [HTML Form](#html-form)
3. [Enctype](#enctype)
4. [ParseForm](#parseform)
5. [PostForm](#postform)
6. [ResponseWriter](#responsewriter)

# Questions
## About Request
* [You are ready to understand about Requests and Response?](#Request-Response)
* [What is format for request-header in Golang?](#Format-request-header)
* [How many methods we can use in Header Golang?](#methods-in-header)
* [What is HTML form?](#HTML-Form)
* [The important things when we use enctype ? And how to insert to HTML form?](#Enctype)
* [What is ParseForm?](#ParseForm)
* [What is PostForm?](#PostForm)
* [What is MultipartForm?](#MultipartForm)
* [How to upload a file to brower, and handle on Golang](#Files)
## About Response
* [What is ResponseWriter in Golang?](#ResponseWriter)
* [How to write a example response to client](#Write)
## About Cookies
* [What is Cookie? How cookie was created? Life cycle for cookie?](#Cookie)
* [Could understand about struct Cookie in Go?](#Struct-cookie)
* [How to send a cookie to the brower?](#Send-cookie)
* [How to get a cookie from brower?](#Get-cookie)
* [Why do we need to use cookie? What is Flash message? Get a example using cookie for flash message?](#Flash-message)

-------------------------------------------------------------------------------------------------------------
## Request-Response
Both requests and responses have basically the same structure:
```
1. Request or response line
2. Zero or more headers
3. An empty line, followed by …
4. … an optional message body 
```
Example format 
```
GET /Protocols/rfc2616/rfc2616.html HTTP/1.1
Host: www.w3.org
User-Agent: Mozilla/5.0 
(empty line)
```

## Format request header
A header is a map, with the key a string and the value a slice of strings.  
Note: map is used as same as dictionary with key-value.  
### Get all of member in Header
```
func headers(w http.ResponseWriter, r *http.Request) {
   h := r.Header
   fmt.Fprintln(w, h)
}
```
Output 
```
map[
  Accept:[text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8] 
  Accept-Encoding:[gzip, deflate] 
  Accept-Language:[en-US,en;q=0.5] 
  Connection:[keep-alive] 
  Sec-Fetch-Dest:[document] 
  Sec-Fetch-Mode:[navigate] 
  Sec-Fetch-Site:[none] 
  Sec-Fetch-User:[?1] 
  Upgrade-Insecure-Requests:[1] 
  User-Agent:[Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:96.0) Gecko/20100101 Firefox/96.0]
]
```
### Get a specific member of Header
```
h := r.Header["Accept-Encoding"]
```
Output
```
[gzip, deflate]
```

## Methods in Header
There are four basic methods on Header, which allow you to add, delete, get, and set values, given a key.
### Get by method
```
h := r.Header.Get("Accept-Encoding")
```

Output
```
gzip, deflate
```
### Set new member to Header
```
	r.Header.Set("Host", "domain.tld")
	h := r.Header.Get("Host")
	fmt.Fprintln(w, h)
```

Output
```
domain.tld
```
### Delete a member in Header
```
   r.Header.Del("Host")
	h2 := r.Header.Get("Host")
	fmt.Fprintln(w, h2, "Nothing we can get")
```

Output
```
 Nothing we can get
```

## Format request body
Body field consists of a Reader interface and a Close interface.  
* A Reader is an interface that has a Read method that takes in a slice of bytes and returns the number of bytes read and an optional error.  
* A Closer is an interface that has a Close method, which takes in nothing and returns an optional error.  
Get a example body
```
func body(w http.ResponseWriter, r *http.Request) {
   len := r.ContentLength
   body := make([]byte, len)
   r.Body.Read(body)
   fmt.Fprintln(w, string(body))
}
```

To run this example
```
curl -id "first_name=sausheong&last_name=chang" 127.0.0.1:8080/body
```
Output
```
HTTP/1.1 200 OK
Date: Tue, 13 Jan 2015 16:11:58 GMT
Content-Length: 37
Content-Type: text/plain; charset=utf-8
first_name=sausheong&last_name=chang
```
More details: [here](https://github.com/huavanthong/MasterGolang/tree/main/01_GettingStarted/book-go-web-application/Chapter_4_Processing_Requests/body)
## HTML-Form
HTML form often look like this:
```
<form action="/process" method="post">
   <input type="text" name="first_name"/>
   <input type="text" name="last_name"/>
   <input type="submit"/>
</form>
```
We can specific POST request with content type of HTML form by using Enctype.
## Enctype
Using enctype for content type request looks like this:
```
<form action="/process" method="post" enctype="application/x-www-
➥ form-urlencoded">
   <input type="text" name="first_name"/>
   <input type="text" name="last_name"/>
   <input type="submit"/>
</form>
```

## ParseForm
Have you ever put a question about element input in HTML form. How we can get this data above, such as: first_name-value? 
That exactly we need ParseForm to do it.
```
func process(w http.ResponseWriter, r *http.Request) {
   r.ParseForm()
   fmt.Fprintln(w, r.Form)
}
```
Suppose we implement html form as below:
```
<form action=http://127.0.0.1:8080/process?hello=world&thread=123 ➥ method="post" enctype="application/x-www-form-urlencoded">
      <input type="text" name="hello" value="sau sheong"/>
      <input type="text" name="post" value="456"/>
      <input type="submit"/>
</form>
```
**Case 1:** In folder form, If only run on brower
```
    http://localhost:8080/process?hello=world&thread=123
Output:
    map[hello:[world] thread:[123]]
```
**Case 2:** In folder form, if run client.html, and submit post request, we will see
```
    map[hello:[sau sheong world] post:[456] thread:[123]]
```
**What is the difference between Case 1 and Case 2?**
- Hello is a key, if you only run case 1, key: hello will have a value: world
- In case 2, key: hello must get a value from input firstly ( name="hello" value="sau sheong"), if it finishs, it will insert more value "world" from action request. That is a reason why you can see value: [sau sheong world]. And have more keys when you parse form request.
## PostForm
Our action, we submit a POST request with:
* first_name = sau sheong
* post = 456
How we can get only post value ==> We can get value by key "post"
```
func process(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    fmt.Fprintln(w, r.Form["post"])
}
```

Output
```
[456]
```
This is a reason we use PostForm
```
func process(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    fmt.Fprintln(w, r.PostForm)
}
```
Output
```
map[hello:[sau sheong] post:[456]]
```
## MultipartForm
```
Summary:
    fmt.Fprintln(w, "(1)", r.FormValue("hello"))
    fmt.Fprintln(w, "(2)", r.PostFormValue("hello"))
    fmt.Fprintln(w, "(3)", r.PostForm)
    fmt.Fprintln(w, "(4)", r.MultipartForm)
Output:
    (1) world
    (2) 
    (3) map[]
    (4) &{map[hello:[sau sheong] post:[456]] map[]}
```
## Files
If you want client can upload file, firstly, you need update HTML Form
```
<form action="http://localhost:8080/process?hello=world&thread=123" method="post" enctype="multipart/form-data">
   <input type="text" name="hello" value="sau sheong"/>
   <input type="text" name="post" value="456"/>
   <input type="file" name="uploaded"> <<<<<<<<<===================
   <input type="submit">
</form>
```
To retrive uploaded files, we can use FormFile.
```
func process(w http.ResponseWriter, r *http.Request) {
   file, _, err := r.FormFile("uploaded")
   if err == nil {
       data, err := ioutil.ReadAll(file)
       if err == nil {
          fmt.Fprintln(w, string(data))
       }
   }
}
```
## ResponseWriter
The ResponseWriter interface has three methods:  
 ■ Write  
 ■ WriteHeader  
 ■ Header  

### Write
Write method takes in an array of bytes, and this gets written into the body of the HTTP response. 
If the header doesn’t have a content type by the time Write is called,the first 512 bytes of the data are used to detect the content type. 
```
func writeExample(w http.ResponseWriter, r *http.Request) {
    str := `
      <html>
         <head><title>Go Web Programming</title></head>
         <body><h1>Hello World</h1></body>
      </html>`
    w.Write([]byte(str))
}
```
### WriteHeader
The WriteHeader method is pretty useful if you want to return error codes. Let’s say you’re writing an API and though you defined the interface, you haven’t fleshed it out, so you want to return a 501 Not Implemented status code.  
```
func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(501)
    fmt.Fprintln(w, "No such service, try next door")
}
```
Output
```
> curl -i 127.0.0.1:8080/writeheader
HTTP/1.1 501 Not Implemented
Date: Tue, 13 Jan 2015 16:20:29 GMT
Content-Length: 31
Content-Type: text/plain; charset=utf-8
No such service, try next door
```
### Header
Finally the Header method returns a map of headers that you can modify.  
The modified headers will be in the HTTP response that’s sent to the client. 
* Use to redirect page.
* Use to forward page.
```
func headerExample(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Location", "http://google.com")
    w.WriteHeader(302)
}
```
Output:
```
> curl -i 127.0.0.1:8080/redirect
HTTP/1.1 302 Found
Location: http://google.com
Date: Tue, 13 Jan 2015 16:22:16 GMT
Content-Length: 0
Content-Type: text/plain; charset=utf-8
```

## Cookie
Please remember some feature below to understand about Cookie: 
1. Cookie là một phần thông tin nhỏ được lưu tại client - brower.
2. Nguồn gốc của cookie được tạo ra, là vì cookie được gửi từ server thông qua HTTP response message.
3. Và mỗi lần client sends một HTTP request đến server, the cookies cũng sẽ được sent along with HTTP request.
4. Cookies được thiết kế để khắc phục (overcome) tình trạng không trạng thái của HTTP.  
More details: [here](https://www.meisternote.com/app/note/JStJbiZPebag/4-4-cookies)
### Struct-cookie
Please refer package cookie: [here](https://go.dev/src/net/http/cookie.go)  

### Send-cookie
We use setCookie() methods to send cookie from server to client.
```
func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "Go Web Programming",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "Manning Publications Co",
		HttpOnly: true,
	}
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}
```
Run serve.go and check result 
```
http://127.0.0.1:8080/set_cookie
=> Select: Storage in Inspect page on brower.
```
### Get-cookie
To get all names-value in cookie:
```
func getCookie(w http.ResponseWriter, r *http.Request) {
   h := r.Header["Cookie"]
   fmt.Fprintln(w, h)
}
```
To get a specific name-value:
```
func getCookie(w http.ResponseWriter, r *http.Request) {
   c1, err := r.Cookie("first_cookie")
   if err != nil {
      fmt.Fprintln(w, "Cannot get the first cookie")
   }
   cs := r.Cookies()
   fmt.Fprintln(w, c1)
   fmt.Fprintln(w, cs)
}
```
To check result cookie:
- Please run API set_cookie first. (If not, you only run get_cookie, we wil can't see data in cookie, 
- Then run API get_cookie to receive data in cookie.
```
http://127.0.0.1:8080/get_cookie 
Then: 
http://127.0.0.1:8080/get_cookie 
```
### Flash-message
These transient messages are commonly known as **flash messages**  
**More detail:** [here](https://www.meisternote.com/app/note/k4YdJh40r5E0/4-4-4-using-cookies-for-flash-messages)  

As you know the problem if you run API to get the cookie firstly, you can't find the cookie, and you want to show message to client “No message found.”. You have to do two thing:
1. Create a cookie with the same name, but with MaxAge set to a negative number and an Expires value that’s in the past.
2. Send the cookie to the browser with SetCookie.  
To send message to client, and encode message to hide info.
```
func setMessage(w http.ResponseWriter, r *http.Request) {
	msg := []byte("Hello World!")
	c := http.Cookie{
		Name:  "flash",
		Value: base64.URLEncoding.EncodeToString(msg),
	}
	http.SetCookie(w, &c)
}
```
Then, we receive messsage: 
- is null, notifiy error
- not null, decode message from client.
```
func showMessage(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("flash")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "No message found")
		}
	} else {
		rc := http.Cookie{
			Name:    "flash",
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		}
		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w, string(val))
	}
}
```


