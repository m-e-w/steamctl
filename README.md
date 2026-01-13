# steamctl
A Go-based CLI tool for querying Steam account data via the Steam Web API.

## Table of Contents

- [steamctl](#steamctl)
  - [Prerequisites](#prerequisites)
    - [Steam Web API Key](#steam-web-api-key)
    - [Steam Profile URL](#steam-profile-url)
  - [Install](#install)
    - [Linux](#linux)
      - [One-line installer (Linux-64)](#one-line-installer-linux-64)
      - [Manual (Linux-64)](#manual-linux-64)
    - [Windows](#windows)
      - [One-line installer (Windows-64)](#one-line-installer-windows-64)
  - [Configure](#configure)
    - [Using the configure command](#using-the-configure-command)
    - [Environment variables](#environment-variables)
  - [Usage](#usage)
    - [List owned games](#list-owned-games)
      - [Example 1: Show top N games by playtime](#example-1-show-top-n-games-by-playtime)
      - [Example 2: Filter results by name](#example-2-filter-results-by-name)
  - [Disclaimer](#disclaimer)

## Prerequisites

**IMPORTANT (PLEASE READ)**

Before you can use steamctl, you will need to first obtain a Steam Web API key as well as identify your Steam Profile URL.

### Steam Web API Key

To request a Steam Web API key, see: https://partner.steamgames.com/doc/webapi_overview/auth#user-keys

Direct link to registration page: https://steamcommunity.com/dev/apikey

### Steam Profile URL
How to find your Steam Profile URL

In the Steam Client/Browser:
1. Open Steam and click on your Profile Name (top right) or avatar.
2. Select "View Profile".
3. Look at the address bar at the top of the page; this is your Steam Profile URL
4. To copy it: Right-click anywhere on the page (or on your profile name) and select "Copy Page URL"

If you still can't find it, see: 
- https://www.youtube.com/watch?v=MkBT5G8sG00
- https://steamcommunity.com/discussions/forum/1/618458030664854265

## Install

### Linux

#### One-line installer (Linux-64)
Below is a one-line install command you can use to download & install steamctl on Linux platforms.

```bash
curl -fsSL https://raw.githubusercontent.com/m-e-w/steamctl/main/scripts/install.sh | bash
```
If you would like to inspect the script first (always advisable), you can download it first with: 
```bash
curl -fsSLO https://raw.githubusercontent.com/m-e-w/steamctl/main/scripts/install.sh
```
And then open using your text editor of choice e.g. vi, nano, etc
```bash
vi install.sh
```

You can then install after using: 
```bash
cat install.sh | bash
```

Verify installation
```bash
steamctl --help
```
#### Manual (Linux-64)

1. Download the binary and checksums.txt file
    ```bash
    curl -LO "https://github.com/m-e-w/steamctl/releases/latest/download/steamctl-linux-amd64"
    curl -LO "https://github.com/m-e-w/steamctl/releases/latest/download/checksums.txt"
    ```
2. Verify checksum
    ```bash
    grep " steamctl-linux-amd64" checksums.txt | sha256sum -c -
    ```
    The output should match `steamctl-linux-amd64: OK`

3. Rename binary to steamctl
    ```bash
    mv steamctl-linux-amd64 steamctl
    ```
4. Make the binary executable
    ```bash 
    chmod +x steamctl
    ```
5. Ensure your user bin directory exists
    ```bash
    mkdir -p ~/.local/bin
    ```
6. Install to your user bin
    ```bash
    mv steamctl ~/.local/bin
    ```
7. Verify installation
    ```bash
    steamctl --help
    ```

### Windows

#### One-line installer (Windows-64)
Below is a one-line install command you can use to download & install steamctl on Windows platforms.

```powershell
curl.exe -fsSLO https://raw.githubusercontent.com/m-e-w/steamctl/main/scripts/install.ps1; ./install.ps1; Remove-Item "install.ps1"
```
If you would like to inspect the script first (always advisable), you can download it first with: 
```powershell
curl.exe -fsSLO https://raw.githubusercontent.com/m-e-w/steamctl/main/scripts/install.ps1
```
And then open using your text editor of choice e.g. notepad, vscode, etc
```powershell
notepad.exe install.ps1
```

You can then install after using: 
```powershell
./install.ps1
```

Verify installation
```powershell
steamctl --help
```

**Important**

By default, running PowerShell scripts may be disabled on Windows systems. If you try running the script below you may see the following error: 
```powershell
PS C:\Users\devtest\tmp> .\install.ps1
.\install.ps1 : File C:\Users\devtest\tmp\install.ps1 cannot be loaded because running scripts is disabled on this system. For more information, see about_Execution_Policies at https:/go.microsoft.com/fwlink/?LinkID=135170.
At line:1 char:1
+ .\install.ps1
+ ~~~~~~~~~~~~~
    + CategoryInfo          : SecurityError: (:) [], PSSecurityException
    + FullyQualifiedErrorId : UnauthorizedAccess
```
This is normal and expected. 

To get around this, you must temporarily adjust the execution policy for your user. 

To get the effective execution policy for the current PowerShell session, use the `Get-ExecutionPolicy` cmdlet.
```powershell
Get-ExecutionPolicy
```
If you see: `Restricted` you need to change your execution policy as such: 
```powershell
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope Process
```
Note: This will only adjust the execution policy for the current PowerShell session (meaning there is no need to try to revert this command after as it only lasts/affects the PowerShell session you ran it under)

To confirm it took effect, run `Get-ExecutionPolicy` again and you should now see: `RemoteSigned`. You should now be able to run the install script using `./install.ps1`

For more details on working with Execution Policies in PowerShell, see: [about_Execution_Policies](https://learn.microsoft.com/en-us/powershell/module/microsoft.powershell.core/about/about_execution_policies?view=powershell-7.5)

## Configure
### Using the configure command
Use the `steamctl configure` command to setup your local config. You will be prompted for your [Steam Web API Key](#steam-web-api-key) and [Steam Profile URL](#steam-profile-url). 

Note: You can enter either your Steam Profile URL or SteamID64. If you enter your profile URL, your Steam ID will be looked up and stored locally for future requests.

A file called `config` will be saved in the `.steamctl` directory within user's home directory, i.e. `~/.steamctl/config`

It uses [TOML](https://toml.io/en/) as the configuration file format, generating a file that looks like:
```
[default]
steam_api_key = ""
steam_id = ""
steam_profile_url =""
```
Future `steamctl` commands will now attempt to load the Steam Web API Key and Steam ID directly from this file. 

Note: Environment variables (i.e. `STEAM_API_KEY` and `STEAM_ID`) take precedence over the locally stored configuration. This allows you to temporarily override your config for a session if needed.

### Environment Variables
Environment variables can optionally be used in addition to the `steamctl configure` command to manage configuration state.

Examples for setting environment variables are included below. Be sure to replace `your_api_key_here` and `your_steam_id_here` below with your Steam Web API key and Steam ID.  

Linux
```bash
export STEAM_API_KEY="your_api_key_here"
export STEAM_ID="your_steam_id_here"
```

Windows (PowerShell)
```powershell
$env:STEAM_API_KEY = "your_api_key_here"
$env:STEAM_ID = "your_steam_id_here"
```

## Usage
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

### List owned games
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
#### Example 1: Show top N games by playtime
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

#### Example 2: Filter results by name
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

## Disclaimer  

This is an unofficial project and is not affiliated with or endorsed by Valve Corporation.
Steam and the Steam logo are trademarks of Valve Corporation.