# steamctl
A Go-based CLI tool for querying Steam account data via the Steam Web API.

## Setup

1. Obtain a Steam Web API key. See: https://partner.steamgames.com/doc/webapi_overview/auth#user-keys
2. Determine your Steam ID
    - First find your Steam URL if you don't know it: https://steamcommunity.com/discussions/forum/1/618458030664854265
    - If you see a random string, in the URL that is your Steam ID. If you have a custom steam profile URL, use the below trick to find it
        ```
        curl your_steam_profile_url | grep steamid
        ```
        - Example: `curl https://steamcommunity.com/id/profile_name/ | grep steamid`
        - You should see something like: `g_rgProfileData = {"url":"https:\/\/steamcommunity.com\/id\/profile_name\/","steamid":"*****************","personaname":"mara","summary":""};`
3. Configure environment variables for your Steam API Key & Steam ID
    ```
    export STEAM_API_KEY=your_api_key_here
    export STEAM_ID=your_steam_id_here
    ```