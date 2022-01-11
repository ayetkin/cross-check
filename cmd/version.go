package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Cross Check",
	Long:  `All software has versions. This is  Cross Check's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Cross Check - v1.1.0\nDeveloped by Ali Yetkin 2021")
	},
}
