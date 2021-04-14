package cmd

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/spf13/cobra"
)

var (
	femaleCmd = &cobra.Command{
		Use:   "female",
		Short: "Generate a random female name",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(randomdata.FullName(randomdata.Female))
		},
	}
)

func init() {
	rootCmd.AddCommand(femaleCmd)
}
