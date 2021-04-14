package cmd

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/spf13/cobra"
)

var (
	maleCmd = &cobra.Command{
		Use:   "male",
		Short: "Generate a random male name",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(randomdata.FullName(randomdata.Male))
		},
	}
)

func init() {
	rootCmd.AddCommand(maleCmd)
}
