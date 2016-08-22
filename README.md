# gocapiclient

A Golang client library for the Guardian's Content API

[![CircleCI](https://circleci.com/gh/guardian/gocapiclient.svg?style=svg)](https://circleci.com/gh/guardian/gocapiclient)

---

This is a work in progress, and currently only supports the `ItemQuery` query type. 

See [theguardian open platorm](http://open-platform.theguardian.com/documentation/) documentation for details of query parameters.

##Example usage

```go
package main

import (
	"github.com/guardian/gocapiclient"
	"github.com/guardian/gocapiclient/queries"
	"log"
)

func main() {
	client := gocapiclient.NewGuardianContentClient("http://content.guardianapis.com/", "yourapikey")
	itemQuery := queries.NewItemQuery("technology/2016/aug/12/no-mans-sky-review-hello-games")

  // Adds query param
	showParam := queries.StringParam{"show-fields", "all"}
	itemQuery.Params = []queries.Param{&showParam}  

	err := client.GetResponse(itemQuery)

	if err != nil {
		log.Fatal(err)
	}

	println(itemQuery.Response.Status)
}
```
