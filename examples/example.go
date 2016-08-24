package main

import (
	"fmt"
	"github.com/guardian/gocapiclient"
	"github.com/guardian/gocapiclient/queries"
	"log"
	"strconv"
)

func main() {
	client := gocapiclient.NewGuardianContentClient("https://content.guardianapis.com/", "my-api-key")
	searchQuery(client)
	searchQueryPaged(client)
	itemQuery(client)
}

func searchQueryPaged(client *gocapiclient.GuardianContentClient) {
	searchQuery := queries.NewSearchQuery()
	searchQuery.PageOffset = int64(10)

	showParam := queries.StringParam{"q", "sausages"}
	params := []queries.Param{&showParam}

	searchQuery.Params = params

	iterator := client.SearchQueryIterator(searchQuery)

	for response := range iterator {
		fmt.Println("Page: " + strconv.FormatInt(int64(response.CurrentPage), 10))
		for _, v := range response.Results {
			fmt.Println(v.ID)
		}
	}
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
