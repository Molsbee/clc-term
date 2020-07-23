package component

import (
	"fmt"
	"github.com/Molsbee/clc-term/clc"
	"github.com/rivo/tview"
	"strings"
)

type CrossDataCenterFirewallPolicies struct {
	clc   clc.CLC
	app   *tview.Application
	text  *tview.TextView
	style Style
}

func NewCrossDataCenterFirewallPolicies(app *tview.Application, clc clc.CLC) *CrossDataCenterFirewallPolicies {
	return &CrossDataCenterFirewallPolicies{
		clc:   clc,
		app:   app,
		text:  tview.NewTextView(),
		style: DefaultStyle,
	}
}

func (cr *CrossDataCenterFirewallPolicies) Style(style Style) {
	cr.style = style
}

func (cr *CrossDataCenterFirewallPolicies) Render() tview.Primitive {
	cr.text.SetTextColor(cr.style.TextColor)
	cr.text.SetBackgroundColor(cr.style.BGColor)
	cr.text.SetChangedFunc(func() {
		cr.app.Draw()
	})

	title := tview.NewTextView()
	title.SetTextColor(cr.style.TextColor)
	title.SetBackgroundColor(cr.style.BGColor)
	title.SetText("Cross Data Center Firewall Policies")

	grid := tview.NewGrid().SetRows(1, -1)
	grid.AddItem(title, 0, 0, 1, 1, 0, 0, false)
	grid.AddItem(cr.text, 1, 0, 1, 1, 0, 0, false)
	return grid
}

func (cr *CrossDataCenterFirewallPolicies) Update(dataCenter string) {
	cr.text.SetText(fmt.Sprintf("Loading policies for data centers (%s)", dataCenter))

	policies := cr.clc.GetCrossDataCenterFirewallPolicies(dataCenter)
	if len(policies) == 0 {
		cr.text.SetText(fmt.Sprintf("No policies found for data center (%s)", dataCenter))
	} else {
		builder := strings.Builder{}
		for _, p := range policies {
			l := fmt.Sprintf("%s   %s   %s\n", p.ID, p.SourceCidr, p.DestinationCidr)
			builder.WriteString(l)
		}
		cr.text.SetText(builder.String())
	}
}

type IntraDataCenterFirewallPolicies struct {
	clc   clc.CLC
	app   *tview.Application
	text  *tview.TextView
	style Style
}

func NewIntraDataCenterFirewallPolicies(app *tview.Application, clc clc.CLC) *IntraDataCenterFirewallPolicies {
	return &IntraDataCenterFirewallPolicies{
		clc:   clc,
		app:   app,
		text:  tview.NewTextView(),
		style: DefaultStyle,
	}
}

func (i *IntraDataCenterFirewallPolicies) Style(style Style) {
	i.style = style
}

func (i *IntraDataCenterFirewallPolicies) Render() tview.Primitive {
	i.text.SetTextColor(i.style.TextColor)
	i.text.SetBackgroundColor(i.style.BGColor)
	i.text.SetChangedFunc(func() {
		i.app.Draw()
	})

	title := tview.NewTextView()
	title.SetTextColor(i.style.TextColor)
	title.SetBackgroundColor(i.style.BGColor)
	title.SetText("Intra Data Center Firewall Policies")

	grid := tview.NewGrid().SetRows(1, -1)
	grid.AddItem(title, 0, 0, 1, 1, 0, 0, false)
	grid.AddItem(i.text, 1, 0, 1, 1, 0, 0, false)
	return grid
}

func (i *IntraDataCenterFirewallPolicies) Update(dataCenter string) {
	i.text.SetText(fmt.Sprintf("Loading policies for data center (%s)", dataCenter))

	policies := i.clc.GetIntraDataCenterFirewallPolicies(dataCenter)
	if len(policies) == 0 {
		i.text.SetText(fmt.Sprintf("No policies found for data center (%s)", dataCenter))
	} else {
		builder := strings.Builder{}
		for _, p := range policies {
			l := fmt.Sprintf("%s   %s   %s\n", p.ID, p.Source, p.Destination)
			builder.WriteString(l)
		}
		i.text.SetText(builder.String())
	}
}
