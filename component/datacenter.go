package component

import (
	"github.com/Molsbee/clc-term/clc"
	"github.com/Molsbee/clc-term/clc/model"
	"github.com/rivo/tview"
	"strings"
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
		style:         defaultStyle,
	}
}

func (d *DataCenter) Style(style Style) {
	d.style = style
}

func (d *DataCenter) Render() tview.Primitive {
	grid := tview.NewGrid().SetColumns(1, -1)
	grid.AddItem(tview.NewBox().SetBackgroundColor(d.style.BGColor), 0, 0, 1, 1, 0, 0, false)
	grid.AddItem(d.dataCenters(), 0, 1, 1, 1, 0, 0, false)
	return grid
}

func (d *DataCenter) dataCenters() tview.Primitive {
	tree := tview.NewTreeView()
	tree.SetBackgroundColor(d.style.BGColor)
	tree.SetSelectedFunc(func(n *tview.TreeNode) {
		n.SetExpanded(!n.IsExpanded())
	})

	root := tview.NewTreeNode("Data Centers")
	root.SetSelectable(false)

	dataCenters := d.clc.GetDataCenters()
	groupChannel := make(chan model.Group, len(dataCenters))
	for _, dataCenter := range dataCenters {
		go func(dc model.DataCenter, channel chan model.Group) {
			dataCenter := d.clc.GetDataCenter(dc.ID)
			hardwareGroup := d.clc.GetGroup(dataCenter.GetHardwareGroupID())
			channel <- hardwareGroup
		}(dataCenter, groupChannel)
	}

	for i := 0; i < len(dataCenters); i++ {
		hardwareGroup := <-groupChannel
		if hardwareGroup.ServersCount != 0 {
			root.AddChild(d.createDataCenter(hardwareGroup.LocationID, hardwareGroup.LocationID, hardwareGroup))
		}
	}

	tree.SetRoot(root)
	return tree
}

func (d *DataCenter) createDataCenter(id, name string, hardwareGroup model.Group) *tview.TreeNode {
	dcNode := tview.NewTreeNode(strings.TrimSpace(name))
	dcNode.SetReference(id)
	dcNode.SetExpanded(false)
	dcNode.SetSelectable(true)
	dcNode.SetSelectedFunc(func() {
		children := children(dcNode.GetChildren())
		for _, g := range hardwareGroup.Groups {
			if !children.Contains(g.ID) {
				dcNode.AddChild(d.createGroup(g))
			}
		}
	})
	return dcNode
}

func (d *DataCenter) createGroup(group model.Group) *tview.TreeNode {
	groupNode := tview.NewTreeNode(group.Name)
	groupNode.SetReference(group.ID)
	groupNode.SetExpanded(false)
	groupNode.SetSelectable(group.ServersCount != 0)
	if group.ServersCount != 0 {
		groupNode.SetColor(d.style.SelectableGroupTextColor)
		groupNode.SetSelectedFunc(func() {
			children := children(groupNode.GetChildren())
			for _, g := range group.Groups {
				if !children.Contains(group.ID) {
					groupNode.AddChild(d.createGroup(g))
				}
			}
			for _, s := range group.GetServers() {
				if !children.Contains(s) {
					groupNode.AddChild(d.createServer(s))
				}
			}
		})
	}
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
