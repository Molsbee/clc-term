package component

import (
	"fmt"
	"github.com/Molsbee/clc-term/clc"
	"github.com/Molsbee/clc-term/clc/model"
	"github.com/rivo/tview"
)

type Credential struct {
	clc               clc.CLC
	app               *tview.Application
	style             Style
	title             *tview.TextView
	text              *tview.TextView
	viewedCredentials map[string]model.Credentials
}

func NewCredential(app *tview.Application, clc clc.CLC) *Credential {
	return &Credential{
		clc:               clc,
		app:               app,
		style:             DefaultStyle,
		title:             tview.NewTextView(),
		text:              tview.NewTextView(),
		viewedCredentials: make(map[string]model.Credentials),
	}
}

func (c *Credential) Style(style Style) {
	c.style = style
}

func (c *Credential) Render() tview.Primitive {
	c.text.SetTextColor(c.style.TextColor)
	c.text.SetBackgroundColor(c.style.BGColor)
	c.text.SetChangedFunc(func() {
		c.app.Draw()
	})

	c.title.SetTextColor(c.style.TextColor)
	c.title.SetBackgroundColor(c.style.BGColor)
	c.text.SetChangedFunc(func() {
		c.app.Draw()
	})

	grid := tview.NewGrid().SetRows(1, -1)
	grid.AddItem(c.title, 0, 0, 1, 1, 0, 0, false)
	grid.AddItem(c.text, 1, 0, 1, 1, 0, 0, false)
	return grid
}

func (c *Credential) Update(serverName string) {
	c.title.SetText(fmt.Sprintf("Credentials: %s", serverName))
	c.text.SetText("Loading server credentials")

	credential, ok := c.viewedCredentials[serverName]
	if !ok {
		credential = c.clc.GetServerCredentials(serverName)
	}

	c.text.SetText(fmt.Sprintf(`Username: %s
Password: %s`, credential.Username, credential.Password))
}
