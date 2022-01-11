package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var flagCONFIG string

var rootCmd = &cobra.Command{
	Use:   "cross-check",
	Short: "A brief description of Cross Check",
	Long: `This tool helps to making health check for pgsql and mysql servers.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}