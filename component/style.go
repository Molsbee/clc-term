package component

import "github.com/gdamore/tcell"

type Style struct {
	DataCenterListBGColor    tcell.Color
	DataCenterListTextColor  tcell.Color
	SelectableGroupTextColor tcell.Color
}

var defaultStyle = Style{
	DataCenterListBGColor:    tcell.NewRGBColor(0, 0, 0),
	DataCenterListTextColor:  tcell.NewRGBColor(230, 230, 230),
	SelectableGroupTextColor: tcell.NewRGBColor(0, 98, 255),
}
