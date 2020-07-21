package cmd

import (
	"github.com/Molsbee/clc-term/clc"
	"github.com/Molsbee/clc-term/clc/model"
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

			clc, err := clc.New(args[0])
			if err != nil {
				log.Fatal("unable to create clc client")
			}
			serverChannel := make(chan model.Server)

			app := tview.NewApplication()
			app.EnableMouse(true)

			grid := tview.NewGrid().SetColumns(30, -1)
			grid.AddItem(component.NewDataCenter(clc, serverChannel).Render(), 0, 0, 1, 1, 0, 0, false)
			serverDetails := component.NewServerDetails()
			view := serverDetails.Render()
			grid.AddItem(view, 0, 1, 1, 1, 0, 0, false)

			go func() {
				for {
					server := <-serverChannel
					serverDetails.UpdateServer(server)
					view.SetChangedFunc(func() {
						app.Draw()
					})
				}
			}()

			if err := app.SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
				panic(err)
			}
		},
	}

	command.AddCommand(dataCenters)
	return command
}
