package queries

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/guardian/gocapimodels/content"
	"io"
	"io/ioutil"
)

type Query interface {
	GetUrl(string) string
	Deserialize(*thrift.TDeserializer, io.ReadCloser) error
}

type ItemQuery struct {
	Params   []Paramable
	Id       string
	Response *content.ItemResponse
}

func NewItemQuery(Id string) *ItemQuery {
	itemQuery := ItemQuery{Id: Id}
	itemQuery.Response = content.NewItemResponse()

	return &itemQuery
}

// TODO: Need to test params to url
func (itemQuery ItemQuery) GetUrl(base string) string {
	paramString := ""

	for i, v := range itemQuery.Params {
		if v != nil {
			paramable := itemQuery.Params[i]
			paramString += paramable.GetParams()
		}
	}

	url := base + itemQuery.Id + paramString

	println(url)

	return url
}

func (itemQuery ItemQuery) Deserialize(deser *thrift.TDeserializer, r io.ReadCloser) error {
	defer r.Close()
	defer deser.Transport.Close()

	rBytes, err := ioutil.ReadAll(r)

	if err != nil {
		return err
	}

	return deser.Read(itemQuery.Response, rBytes)
}
