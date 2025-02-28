package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jegj/g3/config"
	"github.com/spf13/cobra"
)

var (
	version    = "dev"
	configFile string
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
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.AddCommand(versionCmd)

	RootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is $HOME/.config/g3/config.json)")
	// RootCmd.PersistentFlags().BoolVarP(&config.Flag.Debug, "debug", "", false, "debug mode")
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `Print the version number`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("g3 version %s\n", version)
	},
}

func initConfig() {
	if configFile == "" {
		dir, err := config.GetDefaultConfigDir()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(1)
		}
		configFile = filepath.Join(dir, "config.json")
	}

	/*
		absPath, err := path.NewAbsolutePath(configFile)
		if err != nil {
			panic(err)
		}

		if err := config.Conf.Load(absPath); err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(1)
		}
	*/
}
