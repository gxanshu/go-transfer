/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

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
	Run: func(cmd *cobra.Command, args []string) {
		sender.Send(args[0])
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-transfer.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(sendCommand)
}
