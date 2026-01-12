package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Create or update steamctl configuration",
	Long: `Create or update the steamctl configuration file.

This command helps you set up steamctl by generating a configuration file 
in your home directory (~/.steamctl/config).
You will be prompted for required values such as your Steam API key
and Steam ID.

Configuration precedence:
  1. Environment variables
  2. Configuration file
  3. Command-line flags

Examples:
  steamctl configure
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		steamAPIKey := ""
		fmt.Print("Steam Web API Key: ")
		fmt.Scan(&steamAPIKey)

		steamID := ""
		fmt.Print("Steam ID: ")
		fmt.Scan(&steamID)

		configContent := fmt.Sprintf(
			"[default]\nsteam_api_key = \"%s\"\nsteam_id = \"%s\"\n",
			steamAPIKey,
			steamID,
		)
		configData := []byte(configContent)

		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		directoryPath := filepath.Join(home, ".steamctl")
		err = os.MkdirAll(directoryPath, 0700)
		if err != nil {
			return err
		}

		err = os.WriteFile(filepath.Join(directoryPath, "config"), configData, 0600)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
}
