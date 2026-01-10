package steam

import (
	"errors"
	"net/url"
	"strings"
)

type friend struct {
	ID string `json:"steamId"`
}

type friendsResponse struct {
	Friends []friend `json:"friends"`
}

type getfriendsResponse struct {
	Response friendsResponse `json:"friendslist"`
}

type PlayerSummary struct {
	ID          string `json:"steamid"`
	Name        string `json:"personaname"`
	ProfileURL  string `json:"profileurl"`
	LastLogOff  int64  `json:"lastlogoff"`
	TimeCreated int64  `json:"timecreated"`
}

type playerSummariesResponse struct {
	Players []PlayerSummary `json:"players"`
}

type getPlayerSummariesResponse struct {
	Response playerSummariesResponse `json:"response"`
}

func (c *Client) getPlayerSummaries(steamIDs []string) ([]PlayerSummary, error) {
	params := url.Values{}
	params.Set("steamids", strings.Join(steamIDs, ","))

	var resp getPlayerSummariesResponse
	err := c.get(
		"https://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/",
		params,
		&resp,
	)
	if err != nil {
		return nil, err
	}

	return resp.Response.Players, nil
}

// GetFriends returns the list of friends for the configured Steam account
func (c *Client) GetFriends(steamID string) ([]PlayerSummary, error) {
	// If no steamID, return error
	if steamID == "" {
		return nil, errors.New("steamID is required")
	}

	params := url.Values{}
	params.Set("relationship", "all")
	params.Set("steamid", steamID)
	var friendsResp getfriendsResponse
	err := c.get(
		"https://api.steampowered.com/ISteamUser/GetFriendList/v0001/",
		params,
		&friendsResp,
	)
	if err != nil {
		return nil, err
	}

	steamIDs := make([]string, 0, len(friendsResp.Response.Friends))
	for _, friend := range friendsResp.Response.Friends {
		steamIDs = append(steamIDs, friend.ID)
	}

	playerSummaries, err := c.getPlayerSummaries(steamIDs)
	if err != nil {
		return nil, err
	}

	return playerSummaries, nil
}
