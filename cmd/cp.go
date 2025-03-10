package cmd

import (
	"log/slog"

	"github.com/jegj/g3/handlers"
	"github.com/spf13/cobra"
)

var cpCmd = &cobra.Command{
	Use:   "cp [filepath]",
	Short: "Add a new file into your storage",
	Args:  cobra.ExactArgs(1),
	Long:  `Add a new file into your storage`,
	RunE:  cp,
}

func cp(cmd *cobra.Command, args []string) error {
	slog.Debug("cp command...")
	err := handlers.Cp(args[0])
	if err != nil {
		return err
	} else {
		return nil
	}
}

func init() {
	RootCmd.AddCommand(cpCmd)
}
