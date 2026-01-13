package cli

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

func ResolveSteamID(flagID, envID string) (string, error) {
	if flagID != "" {
		return flagID, nil
	}
	if envID != "" {
		return envID, nil
	}
	return "", fmt.Errorf("steamid must be provided via --id flag or STEAM_ID environment variable")
}

func DetectSteamInput(input string) (steamID string, profileURL *url.URL, err error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", nil, errors.New("input is empty")
	}

	// Case 1: SteamID64 (all digits)
	if regexp.MustCompile(`^\d+$`).MatchString(input) {
		return input, nil, nil
	}

	u, err := url.Parse(input)
	if err != nil {
		return "", nil, fmt.Errorf("invalid url: %w", err)
	}

	host := strings.ToLower(u.Host)
	if host != "steamcommunity.com" && host != "www.steamcommunity.com" {
		return "", nil, errors.New("not a steam community url")
	}

	path := strings.Trim(u.Path, "/")
	if !strings.HasPrefix(path, "id/") && !strings.HasPrefix(path, "profiles/") {
		return "", nil, errors.New("not a steam profile url")
	}

	return "", u, nil
}
