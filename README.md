# steamctl
A Go-based CLI tool for querying Steam account data via the Steam Web API.

## Setup

1. Obtain a Steam Web API key.  
    See: https://partner.steamgames.com/doc/webapi_overview/auth#user-keys
2. Determine your Steam ID
    - First find your Steam URL if you don't know it: https://steamcommunity.com/discussions/forum/1/618458030664854265
    - If you see a numeric ID in the URL, that is your Steam ID. If you have a custom profile URL (vanity name), use the method below to find it:
        ```bash
        curl your_steam_profile_url | grep steamid
        ```
        - Example: `curl https://steamcommunity.com/id/profile_name/ | grep steamid`
        - You should see something like: `g_rgProfileData = {"url":"https:\/\/steamcommunity.com\/id\/profile_name\/","steamid":"*****************","personaname":"mara","summary":""};`
3. Configure environment variables for your Steam API key and Steam ID
    ```bash
    export STEAM_API_KEY=your_api_key_here
    export STEAM_ID=your_steam_id_here
    ```
4. Download the binary  
    ```bash
    curl -L "https://github.com/m-e-w/steamctl/releases/download/v0.1.3/steamctl-linux-amd64" -o steamctl
    ```
    Release: https://github.com/m-e-w/steamctl/releases/tag/v0.1.3
5. Verify checksum
    ```bash
    sha256sum steamctl
    ```
    The output should match the SHA256 checksum shown in the GitHub Releases UI.
    
    Example: 

    ![Github Releases UI Image](docs/screenshots/github_releases_sha256.png?raw=true "Github Releases UI Image")

6. Make the binary executable
    ```bash 
    chmod +x steamctl
    ```
7. Ensure your user bin directory exists
    ```bash
    ls ~/.local/bin
    ```
    If it does not exist:
    ```bash
    mkdir -p ~/.local/bin
    ```
8. Install to your user bin
    ```bash
    mv steamctl ~/.local/bin
    ```
9. Reload your bash profile
    ```bash
    source ~/.profile
    ```
    Note:  
    - ~/.profile is typically where `PATH="$HOME/.local/bin:$PATH"` gets added to PATH. 
    - ~/.profile is only loaded on login (meaning new terminal windows will not have ~/.local/bin in PATH)
    - If you are using SSH, simply exiting and reinitiating the connection should suffice
    - If using WSL, you may need to shutdown/reload to avoid having to source ~/.profile every time you open a new terminal going forwards
        - i.e. `wsl --shutdown`
10. Verify installation
    ```bash
    steamctl --help
    ```

## Usage
```
Usage:
  steamctl [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  friends     List friends associated with the Steam account
  games       List games owned by the Steam account
  help        Help about any command

Flags:
  -f, --format string   output format (table, json) (default "table")
  -h, --help            help for steamctl
  -l, --limit int       maximum number of results to display on screen (default 500)
  -q, --quiet           suppress non-essential output
  -v, --version         version for steamctl

Use "steamctl [command] --help" for more information about a command.
```

### List owned games

```bash
steamctl games -s playtime -l 5
```
```
#  ID       NAME                 PLAYTIME (hrs)  LAST PLAYED
-  --       ----                 --------------  -----------
1  892970   Valheim              378.37          2025-03-10 17:26:20
2  578080   PUBG: BATTLEGROUNDS  334.58          2018-10-19 01:56:48
3  1245620  ELDEN RING           274.92          2024-07-14 10:03:15
4  440      Team Fortress 2      261.33          2022-06-26 23:24:53
5  1086940  Baldur's Gate 3      220.83          2023-10-09 03:09:38

Showing 5 of 348 games
```
The above command returns a list of owned games using a Steam ID loaded from environment variables.  
It sorts the games by playtime (descending) and returns the first 5 items. 

```
Usage:
  steamctl games [filter] [flags]

Flags:
  -h, --help          help for games
  -i, --id string     Steam ID of user to retrieve owned games from
  -s, --sort string   Sort by: name, playtime, lastplayed (default "name")

Global Flags:
  -f, --format string   output format (table, json) (default "table")
  -l, --limit int       maximum number of results to display on screen (default 500)
  -q, --quiet           suppress non-essential output
```

## Disclaimer  

This is an unofficial project and is not affiliated with or endorsed by Valve Corporation.
Steam and the Steam logo are trademarks of Valve Corporation.