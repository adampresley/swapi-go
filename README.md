# SWAPI for Go

SWAPI, or the [Star Wars API](http://swapi.co/) is a RESTful based API for Star Wars data. If you want information about the canon Star Wars universe, and wish to get this data via an API, then the SWAPI is for you! This is a Go-based package for consuming this library. This wrapper does not make use of generic interfaces, but instead provides concrete structures for each resource type and collection.

It is still early in development, and pull requests are welcome!

### Getting Started

To begin using the SWAPI Go wrapper add the following line to your import statements.

```go
import "github.com/adampresley/swapi-go/swapi"
```

Next up, the basics.

### Basics

Basically this package works by creating an instance of a `SWAPIClient` and calling a method to return the appropriate Star Wars data. Each API method returns three pieces of information: a result structure, an HTTP status code, and an error object. The result structure will vary based on the method called. For example calling **GetAllPeople()** will return a `PeopleCollection` structure, while **GetPersonById(1)** will return a `People` structure.

If there is a hard error, such as communicating with the server, the error structure will be non-nil. If the error is one returned by the API, the error structure will be nil, however the *ErrorMessage* key in the result structure will contain a message, and the HTTP status will reflect something other than **200**.

For more information [visit the wiki](https://github.com/adampresley/swapi-go/wiki).

### Examples

Below are a few examples.

#### Getting People

```go
package main

import (
	"log"

	"github.com/adampresley/swapi-go/swapi"
)

func main() {
	client := swapi.NewClient()
	result, status, err := client.GetAllPeople()

	if err != nil {
		log.Fatalf("Cannot get people: %s", err.Error())
	}

	if status != 200 {
		log.Println("HTTP error with message", result.ErrorMessage)
	}

	log.Println(result)
}
```

#### Get a Person

```go
package main

import (
	"log"

	"github.com/adampresley/swapi-go/swapi"
)

func main() {
	client := swapi.NewClient()
	result, status, err := client.GetPersonById(1)

	if err != nil {
		log.Fatalf("Cannot get person: %s", err.Error())
	}

	if status != 200 {
		log.Println("HTTP error with message", result.ErrorMessage)
	}

	log.Println(result)
}
```

### License
The MIT License (MIT)

Copyright (c) 2014 Adam Presley

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.