package cmd

import (
	"log/slog"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Show all the files in your storage",
	Long:  `Show all the files in your storage`,
	RunE:  ls,
}

func ls(cmd *cobra.Command, args []string) error {
	slog.Info("ls command...")
	return nil
}

func init() {
	RootCmd.AddCommand(lsCmd)
}
