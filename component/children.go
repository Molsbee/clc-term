package component

import "github.com/rivo/tview"

type children []*tview.TreeNode

func (c children) Contains(id string) bool {
	for _, child := range c {
		if child.GetReference().(string) == id {
			return true
		}
	}
	return false
}
