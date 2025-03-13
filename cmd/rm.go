package cmd

import (
	"log/slog"

	"github.com/jegj/g3/config"
	"github.com/jegj/g3/handlers"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Delete file from your storage",
	Args:  cobra.ExactArgs(1),
	Long:  `Delete file from your storage`,
	RunE:  rm,
}

func rm(cmd *cobra.Command, args []string) error {
	slog.Debug("rm command...")
	handler := handlers.NewG3BaseHandler(config.Conf.GHToken)
	err := handler.Rm(args[0])
	if err != nil {
		return err
	} else {
		return nil
	}
}

func init() {
	RootCmd.AddCommand(rmCmd)
}
