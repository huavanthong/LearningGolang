# Introduction 
To get details contents, please refer to book go-web-application.  

# Table of Contents
1. [Serving Go](#serving-go)
2. [Handlers and handler functions](#handlers-and-handler-functions)
3. [Using HTTP/2](#using-http2)
# Questions
This tutorial will help you answer question below:
## About Serving Go
* [Firstlu, when you start to investigate feature in this chapter? Are you understanding about Architecture for Handling requests with Go server or not?](https://www.meisternote.com/app/note/B6NG-U69TSGK/3-2-serving-go)
## About Server configuration
* [In a simple word, http.Server in Go is just a struct Configuration? Do you know this struct?](#the-server-struct-configuration)
* [Could you get a example to configure your server on Golang?](#setting-configure-for-server)
* [How do you set timeout read/write on server using configuration?](#timeout-with-server-configuration)
* [Could you create a issue to waiting from server forever?](#issue-waiting-forever)


## About HTTPs
* [What method to protect the communication between client and server?](#serving-through-https)
* [What process to handle a security communication?](#process-handle)
* [What is SSL, TLS, and HTTPS?](#ssl-tls-and-https)
* [What is a SSL/TLS certificate?](#ssltls-certificate)
* [You put a SSL certificate on server or client?](#ssltls-certificate)
* [How client can use SSL certificate?](#how-client-use-it)
* [How do you generate a SSL certificate?](#generate-ssl-certificates)
* [How do you generate a private key? Difference between the private key and SSL certificate?](#generate-private-key)

## About Handlers and handler functions
* [What is a handler in Golang?](#handler)
* [What is a handler function in Golang?](#handler-functions)
* [How you use chain handler or chain function in Golang](#chain-handler-and-handler-function)
* [What purpose to use chain handler function?](#what-purpose-chain-handler)


## About Using HTTP/2
* [Why we need to use HTTP/2](#using-http2)
* [What is cURL?Why we need it]()

# Getting Started
## Serving Go
### The Server struct configuration
Configuration intergrate many features in this struct. It includes:
* Setting the timeout for reading the request.
* Setting the timeout for writing the response.
* Setting an error logger for the Server struct.
* ....

```
type Server struct {
    Addr string
    Handler Handler
    ReadTimeout time.Duration
    WriteTimeout time.Duration
    MaxHeaderBytes int
    TLSConfig *tls.Config
    TLSNextProto map[string]func(*Server, *tls.Conn, Handler)
    ConnState func(net.Conn, ConnState)
    ErrorLog *log.Logger
}
```

#### Setting configure for server
```
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}
```
More details: [here](https://github.com/huavanthong/MasterGolang/tree/main/01_GettingStarted/book-go-web-application/Chapter_3_Handling_Requests/configurable)

#### Timeout with server configuration
Server configuration have some parameter help you set timeout on server.  
You can set it following sample code below.
```
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: WriteTimeout + 10*time.Millisecond, //10ms Redundant time
		IdleTimeout:  15 * time.Second,
	}
```
However, the question that you need to answer is how REST API, or handler function can use that timeout on server.  
To understand deeply about it, you need to learn about Context. More details: [here](https://www.meisternote.com/app/note/20mo3LrNH_Oz/context)  

So, we can implement context in handler function following sample code.
```
	ctx, _ := context.WithTimeout(context.Background(), WriteTimeout)
	worker, cancel := context.WithCancel(context.Background())
```
More details: [here](https://github.com/huavanthong/MasterGolang/tree/feature/chapter3/01_GettingStarted/book-go-web-application/Chapter_3_Handling_Requests/config_timeout)
### Issue waiting forever
To understand about this issue.  

Please refer here: [server_timeout_issue](https://github.com/huavanthong/MasterGolang/blob/feature/chapter3-timeout/01_GettingStarted/book-go-web-application/Chapter_3_Handling_Requests/config_timeout/server_timeout_issue.go)

### Resolve issue waiting forever
To understand about this solution.  

Please refer here: [server_timeout_solution](https://github.com/huavanthong/MasterGolang/blob/feature/chapter3-timeout/01_GettingStarted/book-go-web-application/Chapter_3_Handling_Requests/config_timeout/server_timeout_solution.go)


### Serving through HTTPS
Most major websites use HTTPS to encrypt and protect the communications between the client and the server when confidential information like passwords and credit card information is shared.
You need to make sure 
#### Process handle
Let take a overrall about process to get more security. Please look this example.  
Client log in a page in brower, request a confidential information like passwords, and credit card information.
    - This request need to be compliant with Payment Card Industry(PCI) Data Security Standard.
    - You need to encrypt the communications between the client and the server.
 

### SSL, TLS, and HTTPS
#### SSL or TLS
* SSL (Secure Socket Layer) is a protocol that provides data encryption and authenti-cation between two parties, usually a client and a server, using Public Key Infrastructure (PKI).
* SSL was originally developed by Netscape and was later taken over by the Internet Engineering Task Force (IETF), which renamed it TLS. 

#### HTTPS
* HTTPS, or HTTP over SSL, is essentially just that—HTTP layered over an SSL/TLS connection.

#### SSL/TLS certificate
* An SSL/TLS certificate (I’ll use the term SSL certificate as it’s more widely known) is used to provide data encryption and authentication.
* An SSL certificate is an X.509 formatted piece of data that contains some information, as well as a public key, stored at a web server.
* SSL certificates are usually signed by a certificate authority (CA), which assures the authenticity of the certificate. 

#### How client use it
* When the client makes a request to the server, it returns with the certificate.
* If the client is satisfied that the certificate is authentic, it will generate a random key and use the certificate (or more specifically the public key in the certificate) to encrypt it.
* This symmetric key is the actual key used to encrypt the data between the client and the server.

### Generate SSL certificates


### Generate private key
To generate a private key. We use the crypto/rsa library and call the GenerateKey function to create an RSA private key
```
pk, _ := rsa.GenerateKey(rand.Reader, 2048)
```


The RSA private key struct that’s created has a public key that we can access, useful when we use the x509.CreateCertificate function to create our SSL certificate.
```
derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template,
➥ &pk.PublicKey, pk)
```


The CreateCertificate function takes a number of parameters, including the Certificate struct and the public and private keys, to create a slice of DER formatted bytes. The rest is relatively straightforward: we use the encoding/pem library to encode the certificate into the cert.pem file:
```
certOut, _ := os.Create("cert.pem")
pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
certOut.Close()
```


We also PEM encode and save the key we generated earlier into the key.pem file
```
keyOut, _ := os.Create("key.pem")
pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes:
➥ x509.MarshalPKCS1PrivateKey(pk)})
keyOut.Close()
```
**Note:** 
    - The certificate is signed by a CA, the certificate file should be the concatenation of the server’s certificate followed by the CA’s certificate.
More details: [here](https://github.com/huavanthong/MasterGolang/tree/main/01_GettingStarted/book-go-web-application/Chapter_3_Handling_Requests/gencert)

## Handlers and handler functions
### Handler
To create a handler in Golang we follow 3 steps.  
**Step 1:** Create a type for handler.
```
type MyHandler struct{}

```


**Step 2:** Create a pointer receiver to ServeHTTP.  
To know about pointer receiver. [here](https://www.meisternote.com/app/note/abVdOWLumtxW/methods-with-pointer-receiver)
```
func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
```

**Step 3:** At main(), initialize a handler, then put to your Server configuration.
More details about Server struct configurtion. [here](#the-server-struct-configuration)
```
func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler,
	}
	server.ListenAndServe()
}
```
More details:
- [a single handler](https://github.com/huavanthong/MasterGolang/tree/main/01_GettingStarted/book-go-web-application/Chapter_3_Handling_Requests/handler) 
- [multiple handler](https://github.com/huavanthong/MasterGolang/tree/main/01_GettingStarted/book-go-web-application/Chapter_3_Handling_Requests/multihandler)


### Handler functions
To create a handler function we follow 2 steps.
**Step 1:** Create a handle function with 2 argument as below
```
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
```

**Step 2:** At main(), use HandleFunc and put URL and handler function to it.
```
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/hello", hello)

	server.ListenAndServe()
}
```
### Chain handler and handler function
* To know chain handler: [chain_handler](https://github.com/huavanthong/MasterGolang/blob/main/01_GettingStarted/book-go-web-application/Chapter_3_Handling_Requests/chain_handler/server.go)
* To know chain function: [chain_handlerfunc](https://github.com/huavanthong/MasterGolang/blob/main/01_GettingStarted/book-go-web-application/Chapter_3_Handling_Requests/chain_handlerfunc/server.go)

### What purpose chain handler
Handlers and handler functions can be chained to allow modular processing of requests through separation of concerns.
Through that, we can do many things with it:
* Catch log before you start to validate a request.
* Validate/protect a request before writing a response.
* Finally, return response message for client.
More details: [chain_handler](https://github.com/huavanthong/MasterGolang/tree/main/01_GettingStarted/book-go-web-application/Chapter_3_Handling_Requests/chain_handler)


## Using HTTP/2


## cURL
* cURL is a command-line tool that allows users to get or send files through a URL. It supports a large number of common internet protocols, including HTTP and HTTPS.  
* cURL is installed by default in many variants of Unix, including OS X, but is also available in Windows.
Starting from version 7.43.0, cURL supports HTTP/2. You can perform a request using the HTTP/2 protocol passing the --http2 flag.

### To download it
> Refer: http://curl.haxx.se/download.html

### To run it
To run curl with https
```
curl -I --http2 --insecure https://localhost:8080/
```
**Note:** Remember, you need to run it against HTTPS. Because you created your own certificate and private key, by default c URL will not proceed as it will try to verify the certificate. To force cURL to accept your certificate, you need to set the insecure flag.  

Output
```
HTTP/2.0 200
content-type:text/plain; charset=utf-8
content-length:12
date:Mon, 15 Feb 2016 05:33:01 GMT
```
More details: [here](https://github.com/huavanthong/MasterGolang/tree/main/01_GettingStarted/book-go-web-application/Chapter_3_Handling_Requests/https)