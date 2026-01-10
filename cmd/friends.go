package cmd

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/m-e-w/steamctl/internal/cli"
	"github.com/m-e-w/steamctl/steam"
	"github.com/spf13/cobra"
)

var friendsCmd = &cobra.Command{
	Use:   "friends [filter]",
	Args:  cobra.MaximumNArgs(1),
	Short: "List friends associated with the Steam account",
	Long: `List all friends associated with the configured Steam account.
	
This command queries the Steam Web API and displays basic profile
information for each friend.

An optional filter may be provided to match friends by name
(case-insensitive).

Examples:
	steamctl friends
	steamctl friends alex
	steamctl friends --sort name`,
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
		items, err := steamClient.GetFriends(steamID)
		if err != nil {
			return err
		}

		// Sort items
		err = cli.SortByKey(items, sortBy, map[string]cli.SortFunc[steam.PlayerSummary]{
			"name": func(a, b steam.PlayerSummary) bool {
				return strings.ToLower(a.Name) < strings.ToLower(b.Name)
			},
			"lastlog": func(a, b steam.PlayerSummary) bool {
				return a.LastLogOff > b.LastLogOff
			},
			"created": func(a, b steam.PlayerSummary) bool {
				return a.TimeCreated > b.TimeCreated
			},
		})
		if err != nil {
			return err
		}

		// Print items
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		defer w.Flush()

		fmt.Fprintln(w, "\n#\tID\tNAME\tLast LOG\tCREATED\tPROFILE URL")
		fmt.Fprintln(w, "-\t--\t----\t--------\t-------\t-----------")

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
				"%d\t%s\t%s\t%v\t%v\t%s\n",
				count,
				item.ID,
				item.Name,
				cli.FormatUnixTime(item.LastLogOff),
				cli.FormatUnixTime(item.TimeCreated),
				item.ProfileURL,
			)
		}
		fmt.Fprintf(w, "\nShowing %d of %d friends\n", count, len(items))

		return nil
	},
}

func init() {
	friendsCmd.SilenceUsage = true
	friendsCmd.Flags().StringVarP(
		&sortBy,
		"sort",
		"s",
		"name",
		"Sort by: created, lastlog, name",
	)
	friendsCmd.Flags().StringVarP(
		&steamIDFlag,
		"id",
		"i",
		"",
		"Steam ID of user to retrieve friends from",
	)
	rootCmd.AddCommand(friendsCmd)
}
