package cmd

import (
	"errors"
	"log/slog"

	"github.com/jegj/g3/config"
	"github.com/jegj/g3/handlers"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [file]",
	Short: "Get file from your storage",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("error: missing required argument [file]")
		}
		return nil
	},
	Long: `The get command allows you to download a specific file from your storage.
  If the remote file exists in your current location, file will be overriden 

  Examples:
  # Get a file from your storage
  g3 get file.txt
  `,
	RunE: get,
}

func get(cmd *cobra.Command, args []string) error {
	slog.Debug("get command...")
	handler := handlers.NewG3BaseHandler(config.Conf)
	err := handler.Rm(args[0])
	if err != nil {
		return err
	} else {
		return nil
	}
}

func init() {
	RootCmd.AddCommand(getCmd)
}
