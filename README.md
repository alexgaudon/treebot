# treebot

A Discord bot that has a 0.3% chance of reacting to any message with a 🌳 emoji.

## Requirements

- Go 1.25+
- A Discord bot token (see [Discord Developer Portal](https://discord.com/developers/applications))

## Running locally

```sh
DISCORD_TOKEN=your-token-here go run main.go
```

## Running with Docker

```sh
docker build -t treebot .
docker run -d --name treebot -e DISCORD_TOKEN=your-token-here treebot
```

## Configuration

| Variable       | Description               | Required |
|----------------|---------------------------|----------|
| `DISCORD_TOKEN`| Discord bot token         | Yes      |

The reaction probability is defined by `reactChance` in `main.go` (default `0.003` = 0.3%).