package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version   = "dev"
	commit    = "unknown"
	buildDate = "unknown"
)

var rootCmd = &cobra.Command{
	Use:   "divine",
	Short: "A general-purpose divination CLI",
	Long:  "Draw cards from tarot, I Ching, creative prompts, and more.",
	Version: version,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetVersionTemplate("divine {{.Version}}\n")
	rootCmd.AddCommand(versionCmd)
}
