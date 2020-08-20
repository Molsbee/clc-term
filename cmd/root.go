package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "clc-term",
	Short: "terminal gui for inspecting clc data",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
	},
}

func init() {
	server.PersistentFlags().StringVarP(&accountAlias, "accountAlias", "a", "", "clc account alias")
	rootCmd.AddCommand(configCmd(), dataCenters, server, firewall)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
