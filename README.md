# ynab-mcp

A [Model Context Protocol](https://modelcontextprotocol.io/) (MCP) server that exposes [YNAB](https://www.ynab.com/) (You Need A Budget) API tools to LLMs. Built with Go using the [mcp-go](https://github.com/mark3labs/mcp-go) framework.

## Features

- List and inspect budgets
- View accounts and balances
- Browse category groups and categories with budgeted amounts, activity, and balances
- Query transactions with optional date filtering
- Create single or bulk transactions
- Two transport modes: **stdio** (for CLI/desktop) and **HTTP** (for server deployment)

## Prerequisites

- Go 1.24+
- A [YNAB personal access token](https://api.ynab.com/#personal-access-tokens)

## Installation

```bash
go install github.com/blaberg/ynab-mcp/cmd/ynab-mcp@latest
```

Or build from source:

```bash
git clone https://github.com/blaberg/ynab-mcp.git
cd ynab-mcp
go build ./cmd/ynab-mcp
```

## Configuration

| Variable | Description | Default |
|---|---|---|
| `YNAB_API_TOKEN` | YNAB personal access token (required for stdio transport) | — |
| `TRANSPORT` | Transport mode: `stdio` or `http` | `stdio` |
| `PORT` | HTTP server port (http transport only) | `8080` |

In HTTP transport mode, the YNAB token is read from the `X-YNAB-Token` request header instead of an environment variable.

## Usage

### With Claude Desktop

Add the following to your Claude Desktop MCP configuration:

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

YNAB amounts are in **milliunits** (1000 = $1.00). Use negative values for outflows.

## Architecture

```
cmd/ynab-mcp/main.go       Entry point — server setup, transport selection
internal/tools/             MCP tool definitions and handlers
internal/ynab/              YNAB REST API client (https://api.ynab.com/v1)
```

## License

MIT
