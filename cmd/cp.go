package cmd

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/jegj/g3/config"
	"github.com/jegj/g3/handlers"
	"github.com/jegj/g3/validator"
	"github.com/spf13/cobra"
)

var description string

// TODO: Present warning of overriding file
var cpCmd = &cobra.Command{
	Use:   "cp [file]",
	Short: "Add a new file into your storage",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("error: missing required argument [file]")
		}
		if !validator.FileExists(args[0]) {
			return errors.New("error: file doesn't exists")
		}
		return nil
	},
	Long: `The "cp" command allows you to add a new file to your storage.

You must provide the file path as an argument, and the command will upload
the specified file to your storage. If a file with the same name already exists, 
it may be overwritten.

Examples:
  # Add a file to storage
  g3 cp /path/to/file.txt

  # Add a file using relative path
  g3 cp ./example.json
`,
	RunE: cp,
}

// TODO: Validate data input
func cp(cmd *cobra.Command, args []string) error {
	slog.Debug("cp command...")
	handler := handlers.NewG3BaseHandler(config.Conf)
	filename := args[0]
	isOverrindingFile := handler.IsOverrindingFile(filename)
	if isOverrindingFile {
		var response string
		fmt.Print("There is file with the same name in your storage. Are you sure you want to override the file? (y/n):")
		_, err := fmt.Scanln(&response)
		if err != nil {
			return err
		}

		if response != "y" && response != "Y" {
			return errors.New("error: canceled by the user")
		}
	}
	err := handler.Cp(filename, description)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func init() {
	RootCmd.AddCommand(cpCmd)
	cpCmd.PersistentFlags().StringVarP(&description, "description", "d", "", "A brief description of your file")
}
