package cmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

var (
	awsAccountIdCmd = &cobra.Command{
		Use:   "aws-account",
		Short: "Generate a random AWS account ID",
		Run: func(cmd *cobra.Command, args []string) {
			rand.Seed(time.Now().UnixNano())
			max := 429999999999
			min := 420000000000
			fmt.Println(rand.Intn(max-min) + min)
			// do something
		},
	}
)

func init() {
	rootCmd.AddCommand(awsAccountIdCmd)
}
