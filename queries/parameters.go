package queries

import (
	"strconv"
)

type Paramable interface {
	GetParams() string
}

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
		"show-tags=" + showParams.ShowTags +
		"show-elements=" + showParams.ShowElements +
		"show-blocks=" + showParams.ShowBlocks +
		"show-section=" + strconv.FormatBool(showParams.ShowSection)

	return paramString
}
