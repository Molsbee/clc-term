package component

import (
	"fmt"
	"github.com/Molsbee/clc-term/clc"
	"github.com/rivo/tview"
	"strings"
)

const (
	crossDCFirewallFormat = "| %-32s | %-6s | %-7v | %-14s | %-8s | %-18s | %-19s | %-8s | %-18s |\n"
	intraDCFirewallFormat = "| %-32s | %-6s | %-7v | %-7s |  %-18s | %-19s | %-18s | %-8s |\n"
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
		builder.WriteString(fmt.Sprintf(crossDCFirewallFormat, "ID", "Status", "Enabled", "Source Account", "Location", "CIDR", "Destination Account", "Location", "CIDR"))
		for _, p := range policies {
			builder.WriteString(fmt.Sprintf(crossDCFirewallFormat, p.ID, p.Status, p.Enabled, p.SourceAccount, p.SourceLocation, p.SourceCidr, p.DestinationAccount, p.DestinationLocation, p.DestinationCidr))
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
		builder.WriteString(fmt.Sprintf(intraDCFirewallFormat, "ID", "Status", "Enabled", "Account", "Source", "Destination Account", "Destination", "Ports"))
		for _, p := range policies {
			builder.WriteString(fmt.Sprintf(intraDCFirewallFormat, p.ID, p.Status, p.Enabled, i.clc.GetAccountAlias(), p.Source[0], p.DestinationAccount, p.Destination[0], p.Ports[0]))
			m := max(max(len(p.Source), len(p.Destination)), len(p.Ports))
			for index := 1; index < m; index++ {
				builder.WriteString(fmt.Sprintf(intraDCFirewallFormat, "", "", "", "", getOrDefault(p.Source, index), "", getOrDefault(p.Destination, index), getOrDefault(p.Ports, index)))
			}
		}
		i.text.SetText(builder.String())
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getOrDefault(list []string, index int) string {
	if len(list) > index {
		return list[index]
	}
	return ""
}
