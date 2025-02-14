package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/Musa-Labs/Spark/cmd"
)


var (
	verbose bool
)

var rootCmd = &cobra.Command{
	Use:   "spark",
	Short: "N8N Development Toolkit",
	Long:  `N8N Development Toolkit`,
}

func init() {
	rootCmd.AddCommand(cmd.NewCmd())
	rootCmd.AddCommand(cmd.DeployCmd())
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
