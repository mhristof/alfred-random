package cmd

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/spf13/cobra"
)

var (
	intCmd = &cobra.Command{
		Use:   "int",
		Short: "Generate a random integer",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(randomdata.Number(100000))
		},
	}
)

func init() {
	rootCmd.AddCommand(intCmd)
}
