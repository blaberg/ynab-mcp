# ynab-mcp

A [Model Context Protocol](https://modelcontextprotocol.io/) (MCP) server that exposes [YNAB](https://www.ynab.com/) (You Need A Budget) API tools to LLMs. Built with Go using the [mcp-go](https://github.com/mark3labs/mcp-go) framework.

## Features

- List and inspect budgets
- View accounts and balances
- Browse category groups and categories with budgeted amounts, activity, and balances
- Query transactions with optional date filtering
- Create single or bulk transactions
- Two transport modes: **stdio** (for CLI/desktop) and **HTTP** (for server deployment)

## YNAB API Token

To use this server you need a YNAB personal access token:

1. Log in to your YNAB account at https://app.ynab.com
2. Go to **Account Settings** → **Developer Settings**
3. Click **New Token** and follow the prompts
4. Copy the token — you'll need it for the configuration below

More details in the [YNAB API docs](https://api.ynab.com/#personal-access-tokens).

## Installation

Pick one of the three methods below.

### Container package (ghcr.io)

```bash
docker pull ghcr.io/blaberg/ynab-mcp:latest
```

Claude Desktop config (`claude_desktop_config.json`):

```json
{
  "mcpServers": {
    "ynab": {
      "command": "docker",
      "args": ["run", "-i", "--rm", "-e", "YNAB_API_TOKEN", "ghcr.io/blaberg/ynab-mcp:latest"],
      "env": {
        "YNAB_API_TOKEN": "your-token-here"
      }
    }
  }
}
```

Claude Code CLI:

```bash
claude mcp add ynab -e YNAB_API_TOKEN=your-token-here -- docker run -i --rm -e YNAB_API_TOKEN ghcr.io/blaberg/ynab-mcp:latest
```

### Go install

Requires Go 1.24+.

```bash
go install github.com/blaberg/ynab-mcp/cmd/ynab-mcp@latest
```

Claude Desktop config (`claude_desktop_config.json`):

```json
{
  "mcpServers": {
    "ynab": {
      "command": "ynab-mcp",
      "env": {
        "YNAB_API_TOKEN": "your-token-here"
      }
    }
  }
}
```

Claude Code CLI:

```bash
claude mcp add ynab -e YNAB_API_TOKEN=your-token-here -- ynab-mcp
```

### Build from source

Requires Go 1.24+.

```bash
git clone https://github.com/blaberg/ynab-mcp.git
cd ynab-mcp
go build ./cmd/ynab-mcp
```

Claude Desktop config (`claude_desktop_config.json`):

```json
{
  "mcpServers": {
    "ynab": {
      "command": "/path/to/ynab-mcp",
      "env": {
        "YNAB_API_TOKEN": "your-token-here"
      }
    }
  }
}
```

Claude Code CLI:

```bash
claude mcp add ynab -e YNAB_API_TOKEN=your-token-here -- /path/to/ynab-mcp
```

## Configuration

| Variable | Description | Default |
|---|---|---|
| `YNAB_API_TOKEN` | YNAB personal access token (required for stdio transport) | — |
| `TRANSPORT` | Transport mode: `stdio` or `http` | `stdio` |
| `PORT` | HTTP server port (http transport only) | `8080` |

In HTTP transport mode, the YNAB token is read from the `X-YNAB-Token` request header instead of an environment variable.

## Usage

### Stdio transport

```bash
export YNAB_API_TOKEN="your-token-here"
ynab-mcp
```

### HTTP transport

```bash
export TRANSPORT=http
export PORT=8080
ynab-mcp
```

Requests must include the `X-YNAB-Token` header with a valid YNAB API token.

## Available Tools

| Tool | Description |
|---|---|
| `list_budgets` | List all budgets in the YNAB account |
| `get_budget` | Get detailed budget info including accounts, categories, and category groups |
| `get_accounts` | List all accounts in a budget with balances |
| `get_categories` | List all category groups and categories with budgeted amounts, activity, and balances |
| `get_transactions` | Get transactions for a budget, optionally filtered by date |
| `create_transaction` | Create a single transaction |
| `create_transactions` | Create multiple transactions in a single request |
| `update_transaction` | Update an existing transaction in a budget |

YNAB amounts are in **milliunits** (1000 = $1.00). Use negative values for outflows.

## Architecture

```
cmd/ynab-mcp/main.go       Entry point — server setup, transport selection
internal/tools/             MCP tool definitions and handlers
internal/ynab/              YNAB REST API client (https://api.ynab.com/v1)
```

## License

MIT
