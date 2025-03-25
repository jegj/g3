package cmd

import (
	"errors"
	"log/slog"

	"github.com/jegj/g3/config"
	"github.com/jegj/g3/handlers"
	"github.com/jegj/g3/validator"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm [file]",
	Short: "Delete file from your storage",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("error: missing required argument [file]")
		}
		if !validator.FileExists(args[0]) {
			return errors.New("error: file doesn't exists")
		}
		return nil
	},
	Long: `The rm command allows you to delete a specific file from your storage. 
  When a file is removed, it is permanently deleted from the Gist, and this action cannot be undone. 
  The command requires the filename to be specified.

  Examples:
  # Delete a file from your storage
  g3 rm /path/to/file.txt

  # Delete a file using relative path
  g3 rm ./example.json
  `,
	RunE: rm,
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
