package cmd

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"

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
		items, total, err := steamClient.GetOwnedGames(steamID)
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

		fmt.Fprintln(w, "\n#\tID\tNAME\tPLAYTIME (hrs)\tLAST PLAYED")
		fmt.Fprintln(w, "-\t--\t----\t--------------\t-----------")

		count := 0
		for _, item := range items {
			if query != "" && !strings.Contains(strings.ToLower(item.Name), query) {
				continue
			}
			if count >= limit {
				break
			}
			count++
			fmt.Fprintf(
				w,
				"%d\t%d\t%s\t%.2f\t%v\n",
				count,
				item.ID,
				item.Name,
				float64(item.PlaytimeForever)/60.0,
				time.Unix(item.RtimeLastPlayed, 0).Local().Format("2006-01-02 15:04:05"),
			)
		}
		fmt.Fprintf(w, "\nShowing %d of %d games\n", count, total)

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
