package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "go-calendar [command]",
}

// Execute Apply CLI commands
func Execute() {
	rootCmd.AddCommand(apiServerCmd)
	rootCmd.AddCommand(grpcServerCmd)
	rootCmd.AddCommand(versionCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
