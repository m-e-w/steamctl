package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/m-e-w/steamctl/internal/cli"
	"github.com/m-e-w/steamctl/steam"
	"github.com/spf13/cobra"
)

var gamesCmd = &cobra.Command{
	Use:   "games [filter]",
	Args:  cobra.MaximumNArgs(1),
	Short: "List games owned by the Steam account",
	Long: `List all games owned by the configured Steam account.

This command queries the Steam Web API and displays owned games
along with total playtime and last played date.

An optional filter may be provided to match games by name (case-insensitive).

Examples:
	steamctl games
	steamctl games half-life
	steamctl games --sort playtime`,
	RunE: func(cmd *cobra.Command, args []string) error {
		query := ""

		if len(args) == 1 {
			query = strings.ToLower(args[0])
		}

		// Validate we have a Steam ID
		steamID, err := cli.ResolveSteamID(steamIDFlag, steamIDEnv)
		if err != nil {
			return err
		}

		// Get items
		items, _, err := steamClient.GetOwnedGames(steamID)
		if err != nil {
			return err
		}

		// Sort items
		err = cli.SortByKey(items, sortBy, map[string]cli.SortFunc[steam.Game]{
			"name": func(a, b steam.Game) bool {
				return strings.ToLower(a.Name) < strings.ToLower(b.Name)
			},
			"playtime": func(a, b steam.Game) bool {
				return a.PlaytimeForever > b.PlaytimeForever
			},
			"lastplayed": func(a, b steam.Game) bool {
				return a.RtimeLastPlayed > b.RtimeLastPlayed
			},
		})
		if err != nil {
			return err
		}

		// Print items
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		defer w.Flush()

		if !quiet && output == "table" {
			fmt.Fprintln(w, "#\tID\tNAME\tPLAYTIME (hrs)\tLAST PLAYED")
			fmt.Fprintln(w, "-\t--\t----\t--------------\t-----------")
		}

		count := 0
		results := make([]steam.Game, 0, limit)
		for _, item := range items {
			if query != "" && !strings.Contains(strings.ToLower(item.Name), query) {
				continue
			}
			if count >= limit {
				break
			}
			count++

			if output == "table" {
				fmt.Fprintf(
					w,
					"%d\t%d\t%s\t%.2f\t%v\n",
					count,
					item.ID,
					item.Name,
					float64(item.PlaytimeForever)/60.0,
					cli.FormatUnixTime(item.RtimeLastPlayed),
				)
			}
			if output == "json" {
				results = append(results, item)
			}

		}
		if output == "json" {
			enc := json.NewEncoder(os.Stdout)
			enc.SetIndent("", "  ")
			return enc.Encode(results)
		}

		return nil
	},
}

func init() {
	gamesCmd.SilenceUsage = true
	gamesCmd.Flags().StringVarP(
		&sortBy,
		"sort",
		"s",
		"name",
		"Sort by: name, playtime, lastplayed",
	)
	gamesCmd.Flags().StringVarP(
		&steamIDFlag,
		"id",
		"i",
		"",
		"Steam ID of user to retrieve owned games from",
	)
	rootCmd.AddCommand(gamesCmd)
}
