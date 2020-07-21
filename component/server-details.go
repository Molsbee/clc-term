package component

import (
	"github.com/Molsbee/clc-term/clc/model"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type ServerDetails struct {
	style Style
	text  *tview.TextView
}

func NewServerDetails() *ServerDetails {
	text := tview.NewTextView()
	text.SetTextColor(tcell.ColorWhite)
	text.SetBackgroundColor(tcell.ColorBlack)
	return &ServerDetails{
		style: defaultStyle,
		text:  text,
	}
}

func (s *ServerDetails) Render() *tview.TextView {
	return s.text
}

func (s *ServerDetails) UpdateServer(server model.Server) {
	s.text.SetText(server.String())
}
