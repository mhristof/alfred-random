package cmd

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/spf13/cobra"
)

var (
	nameCmd = &cobra.Command{
		Use:   "name",
		Short: "Generate a random name",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(randomdata.FullName(randomdata.RandomGender))
		},
	}
)

func init() {
	rootCmd.AddCommand(nameCmd)
}
