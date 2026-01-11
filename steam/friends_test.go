package steam

import (
	"os"
	"testing"
)

func TestGetFriends(t *testing.T) {
	apiKey := os.Getenv("STEAM_API_KEY")
	steamID := os.Getenv("STEAM_ID")

	// Fail immediately if we have no key or steam id
	if apiKey == "" || steamID == "" {
		t.Fatalf("CONFIG ERROR: STEAM_API_KEY or STEAM_ID not set")
	}

	client := NewClient(apiKey, false)

	items, err := client.GetFriends(steamID)
	if err != nil {
		t.Fatalf("GetFriends failed: %v", err)
	}

	count := len(items)

	// Validate we have at least 1 record returned
	if count <= 0 {
		t.Fatalf("expected at least 1 item, got %d", count)
	}

	// Using the first record, confirm it has a ID and Name present (minimum friend data)
	item := items[0]
	if item.ID == "" {
		t.Error("expected item ID to be set")
	}
	if item.Name == "" {
		t.Error("expected item name to be set")
	}
}
