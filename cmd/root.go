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
	Use:   "pigsay [flags] text...",
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
		eye = eye[0:1] // take first character
		wrapper := wordwrap.Wrapper(30, true)
		wrapped := wrapper(strings.Join(args, " "))
		wrappedWithIndent := wordwrap.Indent(wrapped, "  ", true)
		splitMessage := strings.Split(wrappedWithIndent, "\n")
		longest := 0
		finalMessage := ""
		for _, v := range splitMessage {
			l := len(v)
			if l > longest {
				longest = l
			}
			if finalMessage != "" {
				finalMessage += "\n"
			}
			finalMessage += strings.Repeat(" ", 6) + v
		}

		bubbleAndText := strings.Repeat(" ", 6) + strings.Repeat("-", longest+2) + "\n" + finalMessage + "\n" + strings.Repeat(" ", 6) + strings.Repeat("-", longest+2)
		cmd.Printf(bubbleAndText)
		cmd.Printf("\n          \\|\n\n            _/|________\n           / " + eye + "         \\\n          E,            |S\n           \\___________/\n            WW       WW")
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
