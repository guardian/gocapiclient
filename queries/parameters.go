package queries

import (
	"strconv"
)

type Paramable interface {
	GetParams() string
}

// TODO: This is a nieve approach and won't scale - perhaps
// ShowParams should have an array of params that each know how
// to be coerced to a string value (i.e. Param structs adhering
// to "Stringify" interface?
type ShowParams struct {
	ShowFields   string
	ShowTags     string
	ShowElements string
	ShowBlocks   string
	ShowSection  bool
}

func (showParams ShowParams) GetParams() string {
	paramString := "?" +
		"show-fields=" + showParams.ShowFields +
		"&show-tags=" + showParams.ShowTags +
		"&show-elements=" + showParams.ShowElements +
		"&show-blocks=" + showParams.ShowBlocks +
		"&show-section=" + strconv.FormatBool(showParams.ShowSection)

	return paramString
}
