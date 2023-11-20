/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"os"
	"strings"

	"github.com/eidolon/wordwrap"
	"github.com/spf13/cobra"
)

var eyes string

var pig string = `
         _/|________
        / %s         \
       E,            |S
        \___________/
         WW       WW
`

var example string = `pigsay
	Display help
pigsay Hello World
	Display pig saying "Hello World"
pigsay -e @ Hello World
	Display pig with custom eye @ saying "Hello World"
pigsay --eyes @ Hello World
	Display pig with custom eye @ saying "Hello World"
`

func buildSpeechBubble(text string) string {
	wrapper := wordwrap.Wrapper(30, true)
	wrapped := wrapper(text)
	splitMessage := strings.Split(wrapped, "\n")
	longest := 0
	for _, v := range splitMessage {
		l := len([]rune(v))
		if l > longest {
			longest = l
		}
	}
	bubbleLeftPadding := 2
	var bubbleSize int
	if bubbleSize = 3; longest > bubbleSize {
		bubbleSize = longest
	}
	//header
	bubble := strings.Repeat(" ", bubbleLeftPadding+1) + strings.Repeat("_", bubbleSize+1) + "\n"
	bubble += strings.Repeat(" ", bubbleLeftPadding) + "/" + strings.Repeat(" ", bubbleSize+1) + "\\\n"
	//body
	size := len(splitMessage)
	for _, v := range splitMessage {
		l := len([]rune(v))
		if size == 1 && (l == 1 || l == 2) {
			bubble += strings.Repeat(" ", bubbleLeftPadding-1) + "|" + strings.Repeat(" ", 2) + v + strings.Repeat(" ", 2) + "|"
		} else {
			bubble += strings.Repeat(" ", bubbleLeftPadding-1) + "|" + "  " + v + strings.Repeat(" ", longest-l+1) + "|"
		}
		bubble += "\n"
	}
	//footer
	bubble += strings.Repeat(" ", bubbleLeftPadding) + "\\" + strings.Repeat("_", 1) + strings.Repeat(" ", 2) + strings.Repeat("_", bubbleSize-2) + "/\n"
	bubble += strings.Repeat(" ", bubbleLeftPadding+2) + "\\|"

	return bubble
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pigsay [-e] <MESSAGE>",
	Short: "Make a pig say things!",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
		return nil
	},
	Example: example,
	Version: "1.0.0",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len([]rune(eyes)) != 1 {
			return errors.New("eyes should be only 1 character, no less no more")
		}
		bubble := buildSpeechBubble(strings.Join(args, " "))
		/*
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
			}*/
		wrapper := wordwrap.Wrapper(30, true)
		wrapped := wrapper(strings.Join(args, " "))
		cmd.Println(wrapped)
		cmd.Printf(bubble)
		cmd.Printf(pig, eyes)
		return nil
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
	rootCmd.PersistentFlags().StringVarP(&eyes, "eyes", "e", "o", "give pig different eyes")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
