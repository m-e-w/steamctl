package steam

import (
	"os"
	"testing"
)

func TestGetOwnedGames(t *testing.T) {
	apiKey := os.Getenv("STEAM_API_KEY")
	steamID := os.Getenv("STEAM_ID")

	if apiKey == "" || steamID == "" {
		t.Skip("STEAM_API_KEY or STEAM_ID not set; skipping test")
	}

	client := NewClient(apiKey, false)

	games, count, err := client.GetOwnedGames(steamID)
	if err != nil {
		t.Fatalf("GetOwnedGames failed: %v", err)
	}

	if count <= 0 {
		t.Fatalf("expected at least 1 game, got %d", count)
	}

	if len(games) == 0 {
		t.Fatal("expected non-empty games slice")
	}

	// Validate we have a ID and Name for at least 1 game
	game := games[0]
	if game.ID == 0 {
		t.Error("expected game ID to be set")
	}
	if game.Name == "" {
		t.Error("expected game name to be set")
	}
}
