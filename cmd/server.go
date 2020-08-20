package cmd

import (
	"github.com/Molsbee/clc-term/clc"
	"github.com/Molsbee/clc-term/component"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
	"log"
)

var accountAlias string

var server = &cobra.Command{
	Use:     "server",
	Short:   "renders a view of the servers information and credentials based on account alias and server name provided",
	Example: "clc-term server UC1MUC1TEST01 --accountAlias MUC1",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 || len(accountAlias) == 0 {
			log.Fatal("please provide an account alias and server name")
		}

		clc, err := clc.New(accountAlias)
		if err != nil {
			log.Fatal("unable to create clc client")
		}
		serverName := args[0]

		app := tview.NewApplication()
		serverDetails := component.NewServerDetails(app, clc)
		serverDetails.UpdateServer(serverName)

		credential := component.NewCredential(app, clc)
		credential.Update(serverName)

		flex := tview.NewFlex()
		flex.AddItem(serverDetails.Render(), 0, 2, false)
		flex.AddItem(credential.Render(), 0, 1, false)
		if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
			panic(err)
		}
	},
}
