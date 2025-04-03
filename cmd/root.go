package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/jegj/g3/config"
	"github.com/jegj/g3/fsdata"
	"github.com/jegj/g3/logger"
	"github.com/spf13/cobra"
)

var (
	version    = "dev"
	configFile string
	debug      bool
)

var RootCmd = &cobra.Command{
	Use:   "g3",
	Short: "g3 - Gist Storage Service CLI",
	Long: `g3 - Gist Storage Service CLI. Just like AWS S3, G3 is lightweight and flexible
CLI that provides an easy way to store and access data with the reliability
of GitHub’s infrastructure. G3 is perfect for developers who need a quick,
reliable, and Git-native way to manage small to medium-sized files without
setting up complex storage systems`,
	SilenceErrors: true,
	SilenceUsage:  true,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		slog.Debug(err.Error())
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initG3)
	RootCmd.AddCommand(versionCmd)
	RootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "config file (default is $HOME/.config/g3/config.json)")
	RootCmd.PersistentFlags().BoolVarP(&debug, "verbose", "v", false, "Enable verbose mode")
	RootCmd.CompletionOptions.DisableDefaultCmd = true
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `Print the version number`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("g3 version %s\n", version)
	},
}

func initG3() {
	initLogger(debug)
	initConfig()
	initData()
}

func initLogger(debugMode bool) {
	level := "WARN"
	if debugMode {
		level = "DEBUG"
	}
	logger.SetUpLogger(level)
}

func initConfig() {
	slog.Debug("Init config....")
	if configFile == "" {
		dir, err := config.CreateConfigDirIfRequired()
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
		configFile = filepath.Join(dir, config.DEFAULT_CONFIG_FILENAME)
	}

	if err := config.Conf.Load(configFile); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	if err := config.Conf.Validate(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func initData() {
	slog.Debug("Init data....")
	err := fsdata.CreateG3DataFolderIfRequired()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
