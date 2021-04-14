package cmd

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/spf13/cobra"
)

var (
	emailCmd = &cobra.Command{
		Use:   "email",
		Short: "Generate a random email address",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(randomdata.Email())
		},
	}
)

func init() {
	rootCmd.AddCommand(emailCmd)
}
