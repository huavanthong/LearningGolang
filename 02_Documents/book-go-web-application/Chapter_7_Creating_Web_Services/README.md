# Introduction
Web services, as you’ll recall from our brief discussion in chapter 1, provide a service to other software programs. This chapter expands on this and shows how you can use Go to write or consume web services. You’ll learn how to create and parse XML and JSON first, because these are the most frequently used data formats with web services. We’ll also discuss SOAP and RESTful services before going through the steps for creating a simple web service in JSON.

To get details contents, please refer to chapter 7 at book go-web-application. 

# Table of content
1. [Introducing web services](#introducing-web-services)

# Questions
This tutorial will help you answer question below:
## About Introducing web services
* [What is definition of web service in a Web Service Architecture?](#Definition)
* [How many types of web services?](#Types-of-web-services)
* [When do you use SOAP-based? or REST-based?](Purpose)
## About Introducing SOAP-based web services
* [How do you use SOAP-based in a simple project?](#Introducing-SOAP-based-web-services)
## About Introducing REST-based web services
## About Parsing and creating XML with Go
## About Parsing and creating JSON with Go
## About Creating Go web services

-------------------------------------------------------------------------------------------------------------
## Introducing web services
### Definition
```
A Web service is a software system designed to support interoperable machine-to-machine interaction over a network. It has an interface described in a machine-processable format (specifically WSDL). Other systems interact with the Web service in a manner prescribed by its description using SOAP messages, typically conveyed using HTTP with an XML serialization in conjunction with other Web-related standards.
```
### Types-of-web-services
There are different types of web services, including SOAP-based, REST-based, and XML-RPC based.  
The two most popular types are REST-based and SOAP-based.  
* SOAP-based web services are mostly being used in enterprise systems.
* REST-based web services are more popular in publicly availableweb services.
### Purpose
In these cases:
* SOAP is used in internalapplications for enterprise integration 
* REST is used for external, third-party developers. 

## Introducing-SOAP-based-web-services
