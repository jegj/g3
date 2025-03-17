package cmd

import (
	"log/slog"

	"github.com/jegj/g3/config"
	"github.com/jegj/g3/handlers"
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
	handler := handlers.NewG3BaseHandler(config.Conf.GHToken)
	err := handler.Ls()
	if err != nil {
		return err
	} else {
		return nil
	}
}

func init() {
	RootCmd.AddCommand(lsCmd)
}
