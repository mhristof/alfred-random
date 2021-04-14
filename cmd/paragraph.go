package cmd

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/spf13/cobra"
)

var (
	paragraphCmd = &cobra.Command{
		Use:   "paragraph",
		Short: "Generate a random paragraph",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(randomdata.Paragraph())
		},
	}
)

func init() {
	rootCmd.AddCommand(paragraphCmd)
}
