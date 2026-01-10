package cmd

import (
	"fmt"
	"os"

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
}

func initConfig() {
	viper.AutomaticEnv()
	steamAPIKey := viper.GetString("STEAM_API_KEY")

	if steamAPIKey == "" {
		fmt.Fprintln(os.Stderr, "STEAM_API_KEY is not set")
		os.Exit(1)
	}

	steamIDEnv = viper.GetString("STEAM_ID")
	debug = false

	steamClient = steam.NewClient(
		steamAPIKey,
		debug,
	)
}
