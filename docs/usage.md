# Usage

```
Usage:
  steamctl [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  configure   Create or update steamctl configuration
  friends     List friends associated with the Steam account
  games       List games owned by the Steam account
  help        Help about any command

Flags:
  -d, --debug           show debug messages
  -h, --help            help for steamctl
  -l, --limit int       maximum number of results to display on screen (default 500)
  -o, --output string   output format (table, json) (default "table")
  -q, --quiet           suppress non-essential output
  -v, --version         version for steamctl

Use "steamctl [command] --help" for more information about a command.
```

## List owned games

```
Usage:
  steamctl games [filter] [flags]

Flags:
  -h, --help          help for games
  -i, --id string     Steam ID of user to retrieve owned games from
  -s, --sort string   Sort by: name, playtime, lastplayed (default "name")

Global Flags:
  -d, --debug           show debug messages
  -l, --limit int       maximum number of results to display on screen (default 500)
  -o, --output string   output format (table, json) (default "table")
  -q, --quiet           suppress non-essential output
```

### Example 1: Show top N games by playtime

`steamctl games -s playtime -l 5`

```
#  ID       NAME                 PLAYTIME (hrs)  LAST PLAYED
-  --       ----                 --------------  -----------
1  892970   Valheim              378.37          2025-03-10 17:26:20
2  578080   PUBG: BATTLEGROUNDS  334.58          2018-10-19 01:56:48
3  1245620  ELDEN RING           274.92          2024-07-14 10:03:15
4  440      Team Fortress 2      261.33          2022-06-26 23:24:53
5  1086940  Baldur's Gate 3      220.83          2023-10-09 03:09:38
```

The above command returns a list of owned games. It sorts the list of games by playtime (descending) and returns the first 5 items.

### Example 2: Filter results by name

`steamctl games bio -s playtime`

```
#  ID      NAME                   PLAYTIME (hrs)  LAST PLAYED
-  --      ----                   --------------  -----------
1  8870    BioShock Infinite      15.13           2015-02-15 00:26:55
2  7670    BioShock               14.83           2020-09-27 11:25:33
3  409710  BioShock Remastered    5.05            2020-09-27 17:50:44
4  8850    BioShock 2             0.75            1970-01-01 19:00:00
5  409720  BioShock 2 Remastered  0.02            2017-09-19 21:13:19
```

The above command displays a list of owned games whose names contain the text "bio". It sorts the list by playtime (descending).