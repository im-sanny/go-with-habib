package main

import "fmt"

func main() {
	fmt.Println("Into the backend development!")
}

/*
The Static Era — Early 1990s (Web 1.0)
--------------------------------------
- Websites were made of static HTML pages.
- There was no backend or dynamic behavior.
- Users requested a webpage, and the server sent back a fixed HTML file.
- There was no dynamic content or user interaction.

Server
------
A server is a computer that provides data or webpages to users upon request.


Server-Side Rendering — Early 2000s (Web 2.0)
---------------------------------------------
- HTML and CSS started being generated dynamically from the backend.
- The server produced HTML using backend languages like PHP and Java.
- This allowed web pages to become more interactive and dynamic.


The AJAX Revolution — 2005 to 2010
----------------------------------
- AJAX stands for Asynchronous JavaScript and XML.
- It allowed browsers to exchange data with servers without reloading the page.
- This led to the rise of APIs and API endpoints.

API Endpoint
------------
A fixed URL on the server used to receive and return data (for example: /api/users).

- After AJAX, data was commonly sent and received in JSON format.
- From this time, backend development became more important — focusing on:
  - API design
  - Authentication
  - Data processing


REST API and JSON
-----------------
- The concept of REST was introduced by Dr. Roy Fielding in 2000.
- REST stands for Representational State Transfer.
- REST is a software architectural style (a set of design principles) for building APIs.

Representational
----------------
How the state of a resource is represented — commonly in JSON, XML, YAML, or HTML.

Resource
--------
An entity or concept such as a User or a Post.

State
-----
The current data of that resource.

Example:

User:
{
  "id": 42,
  "name": "Fumis",
  "age": 20
}

Post:
{
  "id": 42,
  "title": "Into the backend development",
  "author": "Fumis"
}

When we present this data in JSON format, we call it a representational state.


Transfer
--------
When a client requests a resource, the server sends that resource’s state (in JSON or other formats) back to represent it. That’s what REST means.


REST vs RESTful
---------------
- REST: A design principle or architectural style.
- RESTful: An API or service that follows REST principles.
  A RESTful service is made up of multiple REST APIs.


API
---
- API stands for Application Programming Interface.
- It acts as a bridge that allows two software systems to communicate.
- The frontend and backend communicate with each other via APIs.
*/
