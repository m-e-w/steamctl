package cli

import "fmt"

func ResolveSteamID(flagID, envID string) (string, error) {
	if flagID != "" {
		return flagID, nil
	}
	if envID != "" {
		return envID, nil
	}
	return "", fmt.Errorf("steamid must be provided via --id flag or STEAM_ID environment variable")
}
