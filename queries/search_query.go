package queries

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/guardian/gocapimodels/content"
	"io"
	"io/ioutil"
)

type SearchQuery struct {
	Params      []Param
	Response    *content.SearchResponse
	PageOffset  int64
	currentPage int64
}

func NewSearchQuery() *SearchQuery {
	searchQuery := SearchQuery{}
	searchQuery.Response = content.NewSearchResponse()
	searchQuery.PageOffset = int64(1)
	searchQuery.currentPage = int64(0)

	return &searchQuery
}

func (self *SearchQuery) NextPage() {
	self.currentPage++
}

func (searchQuery SearchQuery) getNextPage() int64 {
	return searchQuery.currentPage + searchQuery.PageOffset
}

func (searchQuery SearchQuery) GetUrl(base string) string {
	pageNumberParam := IntParam{"page", searchQuery.getNextPage()}
	searchQuery.Params = append(searchQuery.Params, pageNumberParam)

	paramString := CreateParamString(searchQuery.Params)

	url := base + "search" + paramString

	return url
}

func (searchQuery SearchQuery) Deserialize(deser *thrift.TDeserializer, r io.ReadCloser) error {
	defer r.Close()
	defer deser.Transport.Close()

	rBytes, err := ioutil.ReadAll(r)

	if err != nil {
		return err
	}

	return deser.Read(searchQuery.Response, rBytes)
}

type SearchPageResponse struct {
	Err            error
	SearchResponse *content.SearchResponse
}
