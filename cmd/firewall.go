package cmd

import (
	"github.com/Molsbee/clc-term/clc"
	"github.com/Molsbee/clc-term/component"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
	"log"
)

var firewall = &cobra.Command{
	Use:   "firewall",
	Short: "renders a view of the cross & intra data center firewall polies for an account",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("please provide a valid account alias")
		}

		clc, err := clc.New(args[0])
		if err != nil {
			log.Fatal("unable to create clc client")
		}

		app := tview.NewApplication()
		cross := component.NewCrossDataCenterFirewallPolicies(app, clc)
		intra := component.NewIntraDataCenterFirewallPolicies(app, clc)

		root := tview.NewTreeNode("Data Centers")
		root.SetColor(component.DefaultStyle.TextColor)
		root.SetSelectable(false)
		for _, dc := range clc.GetDataCenters() {
			node := tview.NewTreeNode(dc.SanitizedName())
			node.SetReference(dc.ID)
			node.SetExpanded(false)
			node.SetSelectable(true)
			node.SetColor(component.DefaultStyle.TextColor)
			root.AddChild(node)
		}

		tree := tview.NewTreeView()
		tree.SetBackgroundColor(component.DefaultStyle.BGColor)
		tree.SetSelectedFunc(func(n *tview.TreeNode) {
			dataCenter := n.GetReference().(string)
			go cross.Update(dataCenter)
			go intra.Update(dataCenter)
		})
		tree.SetRoot(root)

		flex := tview.NewFlex()
		flex.AddItem(tree, 0, 1, false)
		flex.AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(cross.Render(), 0, 1, false).
			AddItem(intra.Render(), 0, 1, false), 0, 3, false)

		app.EnableMouse(true)
		if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
			log.Fatal(err)
		}
	},
}
