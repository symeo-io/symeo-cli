package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "symeo",
	Short: "Symeo CLI is used to inject configuration values as environment variables into any process",
	Long:  `Symeo is a simple, end-to-end encrypted service that enables teams to sync and manage their configurations and secrets across their development life cycle`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
