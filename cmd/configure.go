package cmd

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"

	"github.com/m-e-w/steamctl/internal/cli"
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
		steamInput := ""
		fmt.Print("Steam profile URL (or SteamID64): ")
		fmt.Scan(&steamInput)

		configContent := ""

		// Take input from user and check if we get a steam id or profile url
		steamID, steamProfileURL, err := cli.DetectSteamInput(steamInput)
		if err != nil {
			return err
		}

		if steamProfileURL != nil {
			// Make the request to get the steam profile page content
			resp, err := http.Get(steamProfileURL.String())
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			// Try to find the steamid in the response
			re := regexp.MustCompile(`"steamid":"([0-9]+)"`)
			match := re.FindSubmatch(body)
			if len(match) > 1 {
				steamID = string(match[1])
			} else {
				return errors.New("no steamid found in steam profile page")
			}

			configContent = fmt.Sprintf(
				"[default]\nsteam_api_key = \"%s\"\nsteam_id = \"%s\"\nsteam_profile_url = \"%s\"\n",
				steamAPIKey,
				steamID,
				steamProfileURL.String(),
			)
		} else {
			configContent = fmt.Sprintf(
				"[default]\nsteam_api_key = \"%s\"\nsteam_id = \"%s\"\n",
				steamAPIKey,
				steamID,
			)
		}

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
	configureCmd.SilenceUsage = true
	rootCmd.AddCommand(configureCmd)
}
