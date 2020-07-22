package cmd

import (
	"github.com/Molsbee/clc-term/clc"
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
			serverChannel := make(chan string)

			app := tview.NewApplication()
			app.EnableMouse(true)

			flex := tview.NewFlex()
			flex.AddItem(component.NewDataCenter(clc, serverChannel).Render(), 0, 1, false)

			serverDetails := component.NewServerDetails(app, clc)
			flex.AddItem(serverDetails.Render(), 0, 2, false)

			credential := component.NewCredential(app, clc)
			flex.AddItem(credential.Render(), 50, 1, false)

			go func() {
				for {
					serverName := <-serverChannel
					go serverDetails.UpdateServer(serverName)
					go credential.Update(serverName)
				}
			}()

			if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
				panic(err)
			}
		},
	}

	command.AddCommand(dataCenters)
	return command
}
