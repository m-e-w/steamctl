# Configuration

## Using the configure command

Use the `steamctl configure` command to setup your local config. You will be prompted for your [Steam Web API Key](prerequisites.md#steam-web-api-key) and [Steam Profile URL](prerequisites.md#steam-profile-url).

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

## Environment Variables

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