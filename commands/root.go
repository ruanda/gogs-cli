package commands

import (
	"fmt"
	"os"

	"github.com/ruanda/gogs-cli/gogs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var GogsClient *gogs.Client

var CmdRoot = &cobra.Command{
	Use:   "gogs-cli",
	Short: "Gogs CLI client",
}

func Execute() {
	if err := CmdRoot.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	CmdRoot.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/gogs-cli.toml)")
	CmdRoot.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigType("toml")
	viper.SetConfigName("gogs-cli")
	viper.AddConfigPath("$HOME/.config")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	GogsClient = gogs.NewClient(
		viper.GetString("gogs.url"),
		viper.GetString("auth.token"),
	)
}
