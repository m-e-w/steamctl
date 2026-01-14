# steamctl

[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE) [![Go Version](https://img.shields.io/badge/Go-1.25.5-00ADD8)](https://golang.org/) [![Release](https://img.shields.io/github/v/release/m-e-w/steamctl)](https://github.com/m-e-w/steamctl/releases)

A Go CLI for fetching game data, playtimes, and friends via the Steam Web API. No Steam client required.

---

## Features

- 🚀 **Fast & Lightweight**: Built in Go for quick queries.
- 📊 **Rich Data**: Get owned games, playtimes, last played dates, and friend lists.
- 🔧 **Flexible Output**: Table or JSON formats.
- 🛠️ **Easy Config**: One-command setup with TOML config or env vars.
- 🌐 **Cross-Platform**: Works on Linux, Windows, and more.

---

## Demo

```bash
steamctl games -s playtime -l 3
```

```
#  ID       NAME          PLAYTIME (hrs)  LAST PLAYED
1  892970   Valheim       378.37          2025-03-10
2  578080   PUBG          334.58          2018-10-19
3  1245620  ELDEN RING    274.92          2024-07-14
```

```bash
steamctl friends -s lastlog -l 3
```

```
#  ID                 NAME          LAST LOG        CREATED         PROFILE URL
1  76561198000000000  Alex          2025-01-10      2010-01-01      https://steamcommunity.com/id/alex
2  76561198000000001  Jordan        2025-01-09      2011-02-02      https://steamcommunity.com/id/jordan
3  76561198000000002  Taylor        2025-01-08      2012-03-03      https://steamcommunity.com/id/taylor
```

---

## Quick Start

1. 🔑 Obtain a [Steam Web API Key](docs/prerequisites.md#steam-web-api-key) and find your [Steam Profile URL](docs/prerequisites.md#steam-profile-url).
2. 🛠️ Install steamctl: [Linux One-liner (64-bit)](docs/installation.md#linux-one-liner-64-bit) | [Windows One-liner (64-bit)](docs/installation.md#windows-one-liner-64-bit).
3. ⚙️ [Configure](docs/configuration.md) your API key and profile.
4. 🚀 Use the tool: `steamctl friends` (see [Usage](docs/usage.md) for more).

---

## Table of Contents

- 📋 [Prerequisites](docs/prerequisites.md)
  - [Steam Web API Key](docs/prerequisites.md#steam-web-api-key)
  - [Steam Profile URL](docs/prerequisites.md#steam-profile-url)
- 🛠️ [Installation](docs/installation.md)
  - [Linux One-liner (64-bit)](docs/installation.md#linux-one-liner-64-bit)
  - [Linux Manual Install (64-bit)](docs/installation.md#linux-manual-install-64-bit)
  - [Windows One-liner (64-bit)](docs/installation.md#windows-one-liner-64-bit)
- ⚙️ [Configuration](docs/configuration.md)
  - [Using the configure command](docs/configuration.md#using-the-configure-command)
  - [Environment Variables](docs/configuration.md#environment-variables)
- 📖 [Usage](docs/usage.md)
  - [List owned games](docs/usage.md#list-owned-games)
  - [List friends](docs/usage.md#list-friends)
- ⚠️ [Disclaimer](docs/disclaimer.md)

---

## Contributing

Found a bug or want to add features? Open an [issue](https://github.com/m-e-w/steamctl/issues) or PR!

## License

MIT - See [LICENSE](LICENSE) for details.