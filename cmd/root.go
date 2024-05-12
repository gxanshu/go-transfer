package cmd

import (
	"fmt"
	"os"

	"github.com/gxanshu/go-transfer/receiver"
	"github.com/gxanshu/go-transfer/sender"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-transfer",
	Short: "Share file from one computer to another from terminal",
	Long: `
    go-transfer is file transfer software written in Go-lang. its super simple to use
    and allow to share file from terminal.
  `,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "display go-transfer version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("go-transfer v1")
	},
}

var sendCommand = &cobra.Command{
	Use:   "send",
	Short: "send a file to another computer",
	Long:  "Provide filename and it will return IP to receive from",
	Run: func(cmd *cobra.Command, args []string) {
		sender.Send(args[0])
	},
}

var receiveCommand = &cobra.Command{
	Use:   "receive",
	Short: "receive a file from another computer",
	Long:  "Fetch file from sender computer",
	Run: func(cmd *cobra.Command, args []string) {
		receiver.Receive(args[0], args[1])
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(sendCommand)
	rootCmd.AddCommand(receiveCommand)
}
