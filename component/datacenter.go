package component

import (
	"github.com/Molsbee/clc-term/clc"
	"github.com/Molsbee/clc-term/clc/model"
	"github.com/rivo/tview"
	"sort"
)

type DataCenter struct {
	clc           clc.CLC
	serverChannel chan string
	style         Style
}

func NewDataCenter(clc clc.CLC, serverChannel chan string) *DataCenter {
	return &DataCenter{
		clc:           clc,
		serverChannel: serverChannel,
		style:         DefaultStyle,
	}
}

func (d *DataCenter) Style(style Style) {
	d.style = style
}

func (d *DataCenter) Render() tview.Primitive {
	tree := tview.NewTreeView()
	tree.SetBackgroundColor(d.style.BGColor)
	tree.SetSelectedFunc(func(n *tview.TreeNode) {
		n.SetExpanded(!n.IsExpanded())
	})

	// Load all data centers and iterate them to get the hardware group for the account in each data center.
	dataCenters := d.clc.GetDataCenters()
	groupChannel := make(chan model.Group, len(dataCenters))
	for _, dataCenter := range dataCenters {
		go func(dc model.DataCenter, channel chan model.Group) {
			hardwareGroup := d.clc.GetGroup(d.clc.GetDataCenter(dc.ID).GetHardwareGroupID())
			hardwareGroup.Name = dc.SanitizedName()
			channel <- hardwareGroup
		}(dataCenter, groupChannel)
	}

	// Populate Data Center Tree
	root := tview.NewTreeNode("Data Centers")
	root.SetSelectable(false)
	for i := 0; i < len(dataCenters); i++ {
		hardwareGroup := <-groupChannel
		if hardwareGroup.ServersCount != 0 {
			root.AddChild(d.createGroup(hardwareGroup))
		}
	}

	// Sort Data Center Tree to view them alphabetically
	sort.Slice(root.GetChildren(), func(i, j int) bool {
		return root.GetChildren()[i].GetText() < root.GetChildren()[j].GetText()
	})

	tree.SetRoot(root)
	return tree
}

func (d *DataCenter) createGroup(group model.Group) *tview.TreeNode {
	groupNode := tview.NewTreeNode(group.Name)
	groupNode.SetReference(group.ID)
	groupNode.SetExpanded(false)
	groupNode.SetSelectable(group.ServersCount != 0)
	if group.ServersCount != 0 {
		groupNode.SetColor(d.style.SelectableGroupTextColor)
	}

	groupNode.SetSelectedFunc(func() {
		children := children(groupNode.GetChildren())
		for _, g := range group.Groups {
			if !children.Contains(g.ID) {
				groupNode.AddChild(d.createGroup(g))
			}
		}
		for _, s := range group.GetServers() {
			if !children.Contains(s) {
				groupNode.AddChild(d.createServer(s))
			}
		}
	})
	return groupNode
}

func (d *DataCenter) createServer(serverName string) *tview.TreeNode {
	serverNode := tview.NewTreeNode(serverName)
	serverNode.SetReference(serverName)
	serverNode.SetExpanded(false)
	serverNode.SetSelectable(true)
	serverNode.SetSelectedFunc(func() {
		d.serverChannel <- serverName
	})
	return serverNode
}
