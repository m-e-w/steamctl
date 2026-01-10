package steam

import (
	"errors"
	"net/url"
)

type Game struct {
	ID                       int    `json:"appid"`
	Name                     string `json:"name"`
	ImgIconURL               string `json:"img_icon_url"`
	HasCommunityVisibleStats bool   `json:"has_community_visible_stats"`
	PlaytimeForever          int    `json:"playtime_forever"`
	RtimeLastPlayed          int64  `json:"rtime_last_played"`
	PlaytimeDisconnected     int    `json:"playtime_disconnected"`
	ContentDescriptorIDs     []int  `json:"content_descriptorids"`
}

type ownedGamesResponse struct {
	GameCount int    `json:"game_count"`
	Games     []Game `json:"games"`
}

type getOwnedGamesResponse struct {
	Response ownedGamesResponse `json:"response"`
}

// GetOwnedGames returns the list of games owned by the configured Steam account along with the total game count.
func (c *Client) GetOwnedGames(steamID string) ([]Game, int, error) {
	if steamID == "" {
		return nil, 0, errors.New("steamID is required")
	}
	params := url.Values{}
	params.Set("include_appinfo", "true")
	params.Set("include_played_free_games", "true")
	params.Set("steamid", steamID)

	var resp getOwnedGamesResponse
	err := c.get(
		"https://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/",
		params,
		&resp,
	)
	if err != nil {
		return nil, 0, err
	}

	return resp.Response.Games, resp.Response.GameCount, nil
}
