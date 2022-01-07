package main

import (
	"fmt"
	"strings"
	"volana/pkg/command"
	"volana/pkg/encryption"
	"volana/pkg/volana"

	"github.com/spf13/cobra"
)

var Secret string

func main() {

	//CMD ENCR

	var cmdEncr = &cobra.Command{ //encrypt command
		Use:   "encr [command to encrypt]",
		Short: "Encrypt a command",
		Long:  `Encrypt a command using AES. Copy output and use decr subcommand to execute it.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			command := strings.Join(args, " ")
			encText, err := encryption.Encrypt(command, Secret)
			if err != nil {
				fmt.Println("error encrypting your classified text: ", err)
			}
			fmt.Println("encr:", encText)
		},
	}

	//CMD DECR
	var cmdDecr = &cobra.Command{
		Use:   "decr [encrypted_command]",
		Short: "decrypt the command and execute it",
		Long:  `decrypt the command previously encrypted with volana, and execute it.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			//Decrypt
			decCmd, err := encryption.Decrypt(args[0], Secret)
			if err != nil {
				fmt.Println("error decrypting your encrypted text: ", err)
			}

			command.Exec(decCmd)
		},
	}
	var rootCmd = &cobra.Command{
		Use:   "volana",
		Short: "Volana enable us to execute command in a stealthy way",
		Run: func(cmd *cobra.Command, args []string) {
			var historic volana.Historic
			for {
				volana.Prefix()
				cmd := volana.GetCommandInteractive(&historic)
				command.ManageSpecialCase(cmd)
				command.Exec(cmd)
			}
		},
	}

	rootCmd.AddCommand(cmdEncr)
	rootCmd.AddCommand(cmdDecr)
	rootCmd.Execute()
}
