package queries

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"io"
)

type Query interface {
	GetUrl(string) string
	Deserialize(*thrift.TDeserializer, io.ReadCloser) error
}

func CreateParamString(params []Param) string {
	if len(params) == 0 {
		return ""
	}

	paramString := "?"

	for _, v := range params {
		paramString += v.ToParamString() + "&"
	}

	return paramString[:len(paramString)-1]
}
