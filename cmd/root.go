/*
Copyright © 2023 Teppo Kavander
*/
package cmd

import (
	"errors"
	"os"
	"strings"

	"github.com/Korppi/pigsay/version"
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

// Builds speech bubble and text wrapped inside it
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
	bubblePaddingLeft := 2
	textPaddingLeft := 1
	bubbleWidth := longest + 2
	bubbleHeaderUpperPadding := bubblePaddingLeft + 1
	bubbleFooterLinesBeforeArrow := 1
	bubbleFooterArrowWidth := 2
	bubbleFooterLinesAfterArrow := bubbleWidth - 3

	//header
	bubble := strings.Repeat(" ", bubbleHeaderUpperPadding) + strings.Repeat("_", bubbleWidth) + "\n"
	bubble += strings.Repeat(" ", bubblePaddingLeft) + "/" + strings.Repeat(" ", bubbleWidth) + "\\\n"
	//body
	for _, v := range splitMessage {
		textPaddingRight := longest - len([]rune(v)) + 1
		bubble += strings.Repeat(" ", bubblePaddingLeft) + "|" + strings.Repeat(" ", textPaddingLeft) + v + strings.Repeat(" ", textPaddingRight) + "|\n"
	}
	//footer
	bubble += strings.Repeat(" ", bubblePaddingLeft) + "\\" + strings.Repeat("_", bubbleFooterLinesBeforeArrow) +
		strings.Repeat(" ", bubbleFooterArrowWidth) + strings.Repeat("_", bubbleFooterLinesAfterArrow) + "/\n"
	bubble += strings.Repeat(" ", bubblePaddingLeft+2) + "\\|"

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
	Version: version.Version,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len([]rune(eyes)) != 1 {
			return errors.New("eyes should be only 1 character, no less no more")
		}
		bubble := buildSpeechBubble(strings.Join(args, " "))
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
	rootCmd.PersistentFlags().StringVarP(&eyes, "eyes", "e", "o", "give pig different eyes")
}
