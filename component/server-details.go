package component

import (
	"github.com/Molsbee/clc-term/clc"
	"github.com/Molsbee/clc-term/clc/model"
	"github.com/rivo/tview"
)

type ServerDetails struct {
	clc           clc.CLC
	app           *tview.Application
	style         Style
	text          *tview.TextView
	viewedServers map[string]model.Server
}

func NewServerDetails(app *tview.Application, clc clc.CLC) *ServerDetails {
	return &ServerDetails{
		clc:           clc,
		app:           app,
		style:         defaultStyle,
		text:          tview.NewTextView(),
		viewedServers: make(map[string]model.Server),
	}
}

func (s *ServerDetails) Style(style Style) {
	s.style = style
}

func (s *ServerDetails) Render() *tview.TextView {
	s.text.SetTextColor(s.style.TextColor)
	s.text.SetBackgroundColor(s.style.BGColor)
	s.text.SetChangedFunc(func() {
		s.app.Draw()
	})
	return s.text
}

func (s *ServerDetails) UpdateServer(serverName string) {
	s.text.SetText("Loading server details")
	server, ok := s.viewedServers[serverName]
	if !ok {
		server = s.clc.GetServer(serverName)
		s.viewedServers[serverName] = server
	}

	s.text.SetText(server.String())
}
