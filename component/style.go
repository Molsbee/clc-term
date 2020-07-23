package component

import "github.com/gdamore/tcell"

type Style struct {
	TextColor                tcell.Color
	BGColor                  tcell.Color
	DataCenterListTextColor  tcell.Color
	SelectableGroupTextColor tcell.Color
}

var DefaultStyle = Style{
	TextColor:                tcell.ColorWhite,
	BGColor:                  tcell.NewRGBColor(0, 0, 0),
	DataCenterListTextColor:  tcell.NewRGBColor(230, 230, 230),
	SelectableGroupTextColor: tcell.NewRGBColor(0, 98, 255),
}
