package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/m-e-w/steamctl/steam"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var version = "dev"

var steamIDFlag string
var steamIDEnv string
var debug bool
var steamClient *steam.Client

var sortBy string
var limit int
var quiet bool
var output string

var rootCmd = &cobra.Command{
	Use:   "steamctl",
	Short: "Query Steam account and library data",
	Long: `steamctl queries data from the Steam Web API.

Use steamctl to retrieve information about Steam accounts, friends,
and game libraries, including owned games, playtime statistics,
and player summaries.`,
	Version: version,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().IntVarP(
		&limit,
		"limit",
		"l",
		500,
		"maximum number of results to display on screen",
	)
	rootCmd.PersistentFlags().BoolVarP(
		&quiet,
		"quiet",
		"q",
		false,
		"suppress non-essential output",
	)
	rootCmd.PersistentFlags().StringVarP(
		&output,
		"output",
		"o",
		"table",
		"output format (table, json)",
	)
	rootCmd.PersistentFlags().BoolVarP(
		&debug,
		"debug",
		"d",
		false,
		"show debug messages",
	)
}

func initConfig() {
	viper.AutomaticEnv()

	steamAPIKey := viper.GetString("STEAM_API_KEY")
	steamIDEnv = viper.GetString("STEAM_ID")
	if steamAPIKey == "" || steamIDEnv == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		configDir := filepath.Join(home, ".steamctl")

		viper.SetConfigName("config")
		viper.SetConfigType("toml")
		viper.AddConfigPath(configDir)

		err = viper.ReadInConfig()
		if err != nil {
			// Do nothing
		}
		steamAPIKey = viper.GetString("default.steam_api_key")
		steamIDEnv = viper.GetString("default.steam_id")
	}

	steamClient = steam.NewClient(
		steamAPIKey,
		debug,
	)
}
