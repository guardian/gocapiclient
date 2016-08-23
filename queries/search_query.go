package queries

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/guardian/gocapimodels/content"
	"io"
	"io/ioutil"
)

type SearchQuery struct {
	Params   []Param
	Response *content.SearchResponse
}

func NewSearchQuery() *SearchQuery {
	searchQuery := SearchQuery{}
	searchQuery.Response = content.NewSearchResponse()

	return &searchQuery
}

func (searchQuery SearchQuery) GetUrl(base string) string {

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
