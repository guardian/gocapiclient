package queries_test

import (
	"bytes"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/guardian/gocapiclient/queries"
	"github.com/guardian/gocapimodels/content"
	"io/ioutil"
	"testing"
)

func TestItemQueryGetUrl(t *testing.T) {
	t.Log("Constructing query url")
	itemQuery := queries.NewItemQuery("id")
	url := itemQuery.GetUrl("http://www.example.com/")

	if url != "http://www.example.com/id" {
		t.Error("Incorrect query url for itemQuery.GetUrl")
	}
}

func TestItemQueryDeserialize(t *testing.T) {
	t.Log("Deserialize thrift for ItemQuery")

	itemQuery := queries.NewItemQuery("id")
	itemResponse := content.ItemResponse{Status: "testing"}

	// Matched pair de/serializers
	serial := thrift.NewTSerializer()
	deser := thrift.NewTDeserializer()

	serialItemResponse, err := serial.Write(&itemResponse)
	itemResponseReadCloser := ioutil.NopCloser(bytes.NewReader(serialItemResponse))

	if err != nil {
		t.Error(err)
	}

	err = itemQuery.Deserialize(deser, itemResponseReadCloser)
	if err != nil {
		t.Error(err)
	}

	if itemQuery.Response.Status != "testing" {
		t.Error("Incorrect data found in deserialized ItemResponse")
	}

}

func TestItemQueryGetUrlParams(t *testing.T) {
	t.Log("Constructing query url with params")

	itemQuery := queries.NewItemQuery("id")
	stringParam := queries.StringParam{
		Key:   "show-example",
		Value: "value",
	}

	itemQuery.Params = []queries.Param{&stringParam}
	url := itemQuery.GetUrl("http://www.example.com/")

	if url != "http://www.example.com/id?show-example=value" {
		t.Error("Incorrect query url for itemQuery.GetUrl")
	}
}
