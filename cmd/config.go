package cmd

import (
	"bufio"
	"fmt"
	"github.com/Molsbee/clc-term/config"
	"github.com/howeyc/gopass"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

func configCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "config",
		Short: "write/read config file for use with application",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(1)
		},
	}

	read := &cobra.Command{
		Use:   "read",
		Short: "read user config file",
		Run: func(cmd *cobra.Command, args []string) {
			config, err := config.Read()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(config)
		},
	}

	write := &cobra.Command{
		Use:   "write",
		Short: "write user config file",
		Run: func(cmd *cobra.Command, args []string) {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("CLC Username: ")
			username, _ := reader.ReadString('\n')
			username = strings.Replace(username, "\r", "", -1)
			username = strings.Replace(username, "\n", "", -1)

			fmt.Print("CLC Password: ")
			passBytes, _ := gopass.GetPasswd()

			if err := config.Write(config.Config{
				Username: username,
				Password: string(passBytes),
			}); err != nil {
				log.Fatal(err)
			}
		},
	}
	command.AddCommand(read, write)

	return command
}
