package queries

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/guardian/gocapimodels/content"
	"io"
	"io/ioutil"
)

type ItemQuery struct {
	Params   []Param
	Id       string
	Response *content.ItemResponse
}

func NewItemQuery(Id string) *ItemQuery {
	itemQuery := ItemQuery{Id: Id}
	itemQuery.Response = content.NewItemResponse()

	return &itemQuery
}

func (itemQuery ItemQuery) GetUrl(base string) string {

	paramString := CreateParamString(itemQuery.Params)
	url := base + itemQuery.Id + paramString

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
