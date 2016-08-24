# gocapiclient

A Golang client library for the Guardian's Content API

[![CircleCI](https://circleci.com/gh/guardian/gocapiclient.svg?style=svg)](https://circleci.com/gh/guardian/gocapiclient)

---

This is a work in progress, and currently only supports the `ItemQuery` and `SearchQuery` query type. 

See [theguardian open platorm](http://open-platform.theguardian.com/documentation/) documentation for details of query parameters.

##Example usage

```go
package main

import (
	"github.com/guardian/gocapiclient"
	"github.com/guardian/gocapiclient/queries"
	"log"
	"fmt"
)

func main() {
	client := gocapiclient.NewGuardianContentClient("https://content.guardianapis.com/", "none")
	searchQuery(client)
	itemQuery(client)
	searchQueryPaged(client)
}

func searchQuery(client *gocapiclient.GuardianContentClient) {
	searchQuery := queries.NewSearchQuery()

	showParam := queries.StringParam{"q", "sausages"}
	params := []queries.Param{&showParam}

	searchQuery.Params = params

	err := client.GetResponse(searchQuery)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(searchQuery.Response.Status)
	fmt.Println(searchQuery.Response.Total)

	for _, v := range searchQuery.Response.Results {
		fmt.Println(v.ID)
	}
}

func itemQuery(client *gocapiclient.GuardianContentClient) {
	itemQuery := queries.NewItemQuery("technology/2016/aug/12/no-mans-sky-review-hello-games")

	showParam := queries.StringParam{"show-fields", "all"}
	params := []queries.Param{&showParam}

	itemQuery.Params = params

	err := client.GetResponse(itemQuery)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(itemQuery.Response.Status)
	fmt.Println(itemQuery.Response.Content.WebTitle)
}

func searchQueryPaged(client *gocapiclient.GuardianContentClient) {
	searchQuery := queries.NewSearchQuery()

	showParam := queries.StringParam{"q", "sausages"}
	params := []queries.Param{&showParam}

	searchQuery.Params = params

	iterator := client.SearchQueryIterator(client, searchQuery)

	for results := range iterator {
		fmt.Println("----- New Page -----")

		for _, v := range results {
			fmt.Println(v.ID)
		}
	}
}
```
