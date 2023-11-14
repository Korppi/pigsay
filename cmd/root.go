/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"strings"

	"github.com/eidolon/wordwrap"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pigsay [flags] {text...}",
	Short: "Make a pig say things!",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
		return nil
	},
	Example: "pigsay Hello World :)\npigsay --eyes @ Now i have different eyes!",
	Version: "1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		eye, _ := cmd.Flags().GetString("eyes")
		wrapper := wordwrap.Wrapper(30, true)
		wrapped := wrapper(strings.Join(args, " "))
		wrappedWithIndent := wordwrap.Indent(wrapped, "  ", true)
		// TODO: calculate longest row of wrappedWithIndent
		// TODO: use repeat function to draw - characters (longest row + 4)
		bubbleAndText := "------------------------------\n" + wrappedWithIndent + "\n------------------------------"
		cmd.Printf(bubbleAndText)
		cmd.Printf("\n          \\|\n\n            _/|________\n           / " + eye[0:1] + "         \\\n          E,            |S\n           \\___________/\n            WW       WW")
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pigsay.yaml)")
	rootCmd.PersistentFlags().StringP("eyes", "e", "o", "give pig different eyes")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
