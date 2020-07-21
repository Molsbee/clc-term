package cmd

import (
	"github.com/Molsbee/clc-term/component"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func gui() *cobra.Command {
	command := &cobra.Command{
		Use:   "gui",
		Short: "load different gui views of clc information",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(1)
		},
	}

	dataCenters := &cobra.Command{
		Use:   "data-centers",
		Short: "render a view of the groups and servers availabe for an account",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				log.Fatal("please provide a valid account alias")
			}

			grid := tview.NewGrid().SetColumns(30, -1)
			dataCenter := component.NewDataCenter(args[0]).Render()
			grid.AddItem(dataCenter, 0, 0, 1, 1, 0, 0, false)
			grid.AddItem(tview.NewBox(), 0, 1, 1, 1, 0, 0, false)

			app := tview.NewApplication()
			app.EnableMouse(true)
			if err := app.SetRoot(grid, true).SetFocus(dataCenter).Run(); err != nil {
				panic(err)
			}
		},
	}

	command.AddCommand(dataCenters)
	return command
}
