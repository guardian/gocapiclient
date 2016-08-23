package queries

import (
	"strconv"
)

type Param interface {
	ToParamString() string
}

type StringParam struct {
	Key   string
	Value string
}

func (p StringParam) ToParamString() string {
	return p.Key + "=" + p.Value
}

type BoolParam struct {
	Key   string
	Value bool
}

func (p BoolParam) ToParamString() string {
	return p.Key + "=" + strconv.FormatBool(p.Value)
}

type IntParam struct {
	Key   string
	Value int64
}

func (p IntParam) ToParamString() string {
	return p.Key + "=" + strconv.FormatInt(p.Value, 10)
}
