package steam

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
)

func TestGetOwnedGames(t *testing.T) {
	apiKey := os.Getenv("STEAM_API_KEY")
	steamID := os.Getenv("STEAM_ID")

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
	apiKey = viper.GetString("default.steam_api_key")
	steamID = viper.GetString("default.steam_id")

	// Fail immediately if we have no key or steam id
	if apiKey == "" || steamID == "" {
		t.Fatalf("CONFIG ERROR: STEAM_API_KEY or STEAM_ID not set")
	}

	client := NewClient(apiKey, false)

	items, count, err := client.GetOwnedGames(steamID)
	if err != nil {
		t.Fatalf("GetOwnedGames failed: %v", err)
	}

	// Validate we have at least 1 record returned
	if count <= 0 {
		t.Fatalf("expected at least 1 item, got %d", count)
	}

	// Using the first record, confirm it has a ID and Name present (minimum game data)
	item := items[0]
	if item.ID == 0 {
		t.Error("expected item ID to be set")
	}
	if item.Name == "" {
		t.Error("expected item name to be set")
	}
}
