package cmd

import (
	"log/slog"

	"github.com/jegj/g3/config"
	"github.com/jegj/g3/handlers"
	"github.com/spf13/cobra"
)

var description string

var cpCmd = &cobra.Command{
	Use:   "cp [filepath]",
	Short: "Add a new file into your storage",
	Args:  cobra.ExactArgs(1),
	Long:  `Add a new file into your storage`,
	RunE:  cp,
}

func cp(cmd *cobra.Command, args []string) error {
	slog.Debug("cp command...")
	handler := handlers.NewG3BaseHandler(config.Conf.GHToken)
	err := handler.Cp(args[0], description)
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
