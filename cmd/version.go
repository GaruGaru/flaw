package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

const Version = 0.1

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print current flaw version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s - v%f\n", Name, Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
