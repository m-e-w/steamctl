# steamctl
A Go-based CLI tool for querying Steam account data via the Steam Web API.

## Table of Contents

- [steamctl](#steamctl)
  - [Prerequisites](#prerequisites)
    - [Steam Web API Key](#1-steam-web-api-key)
    - [Steam ID](#2-steam-id)
    - [Environment variables](#3-environment-variables)
  - [Install](#install)
    - [One-line installer (Linux-64)](#one-line-installer-linux-64)
    - [One-line installer (Windows-64)](#one-line-installer-windows-64)
    - [Manual (Linux-64)](#manual-linux-64)
  - [Usage](#usage)
    - [List owned games](#list-owned-games)
      - [Example 1: Show top N games by playtime](#example-1-show-top-n-games-by-playtime)
      - [Example 2: Filter results by name](#example-2-filter-results-by-name)
  - [Disclaimer](#disclaimer)

## Prerequisites

**IMPORTANT (PLEASE READ)**

Before you can use steamctl, you will need to first obtain a Steam Web API key as well as identify your Steam ID.

You will also need to set the values of both of these as environment variables. 

### 1. Steam Web API Key

To request a Steam Web API key, see: https://partner.steamgames.com/doc/webapi_overview/auth#user-keys

## 2. Steam ID

First find your Steam URL if you don't know it: https://steamcommunity.com/discussions/forum/1/618458030664854265

Next, look at the URL  

If you see a numeric ID in the URL, that is your Steam ID.  

If you have a custom profile URL (vanity name), use the method below to find it

```bash
curl your_steam_profile_url | grep steamid
```

Example

```bash
curl https://steamcommunity.com/id/profile_name/ | grep steamid
```

You should see something like: 

```
g_rgProfileData = {"url":"...","steamid":"*****************","personaname":"mara","summary":""};`
```
steamid should be highlighted in the output

## 3. Environment variables

Replace `your_api_key_here` and `your_steam_id_here` below with your Steam Web API key and Steam ID.  

Linux
```bash
export STEAM_API_KEY="your_api_key_here"
export STEAM_ID="your_steam_id_here"
```

Optionally, to avoid pasting your API key in the terminal, you can use a .env file. 

If you clone the repos there is a .env.example file you can rename and then use to store/load configuration variables. 

You can store them there and then load them using `source .env`

Windows (PowerShell)
```powershell
$env:STEAM_API_KEY = "your_api_key_here"
$env:STEAM_ID = "your_steam_id_here"
```

## Install

### One-line installer (Linux-64)
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

### One-line installer (Windows-64)
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

**Important**

By default, running powershell scripts may be disabled on Windows systems. If you try running the script below you may see the following error: 
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
Note: This will only adjust the execution policy for the current PowerShell session (meaning there is no need to try to revert this command after as it only lasts/effects the PowerShell session you ran it under)

To confirm it took effect, run `Get-ExecutionPolicy` again and you should now see: `RemoteSigned`. You should now be able to run the install script using `./install.ps1`

For more details on working with Execution Policies in PowerShell, see: https://learn.microsoft.com/en-us/powershell/module/microsoft.powershell.core/about/about_execution_policies?view=powershell-7.5

### Manual (Linux-64)

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
    ls ~/.local/bin
    ```
    If it does not exist:
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
The above command returns a list of owned games. It sorts the list of list by playtime (descending) and returns the first 5 items. 

#### Example 2: Filter results by name
```bash
steamctl games bio -s playtime
```
```
#  ID      NAME                   PLAYTIME (hrs)  LAST PLAYED
-  --      ----                   --------------  -----------
1  8870    BioShock Infinite      15.13           2015-02-15 00:26:55
2  7670    BioShock               14.83           2020-09-27 11:25:33
3  409710  BioShock Remastered    5.05            2020-09-27 17:50:44
4  8850    BioShock 2             0.75            1970-01-01 19:00:00
5  409720  BioShock 2 Remastered  0.02            2017-09-19 21:13:19

Showing 5 of 348 games
```
The above command displays a list of owned games whose names contain the text "bio". It sorts the list by playtime (descending).

## Disclaimer  

This is an unofficial project and is not affiliated with or endorsed by Valve Corporation.
Steam and the Steam logo are trademarks of Valve Corporation.