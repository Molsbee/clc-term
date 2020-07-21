package component

import (
	"github.com/Molsbee/clc-term/clc"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type ServerDetails struct {
	clc   clc.CLC
	app   *tview.Application
	style Style
	text  *tview.TextView
}

func NewServerDetails(app *tview.Application, clc clc.CLC) *ServerDetails {
	text := tview.NewTextView()
	text.SetTextColor(tcell.ColorWhite)
	text.SetBackgroundColor(tcell.ColorBlack)
	text.SetChangedFunc(func() {
		app.Draw()
	})
	return &ServerDetails{
		clc:   clc,
		app:   app,
		style: defaultStyle,
		text:  text,
	}
}

func (s *ServerDetails) Render() *tview.TextView {
	return s.text
}

func (s *ServerDetails) UpdateServer(serverName string) {
	s.text.SetText("Loading server details")
	server := s.clc.GetServer(serverName)
	s.text.SetText(server.String())
}
