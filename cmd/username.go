package cmd

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/spf13/cobra"
)

var (
	usernameCmd = &cobra.Command{
		Use:   "username",
		Short: "Generate a random username",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(fmt.Sprintf("%s-%d", randomdata.SillyName(), randomdata.Number(1000)))
		},
	}
)

func init() {
	rootCmd.AddCommand(usernameCmd)
}
