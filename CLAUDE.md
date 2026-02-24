# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Go-based MCP (Model Context Protocol) server that exposes YNAB (You Need A Budget) API tools to LLMs. Uses the `mark3labs/mcp-go` framework. Supports two transport modes: `stdio` (default, for CLI) and `http` (for server deployment).

## Build & Run Commands

```bash
go build ./cmd/ynab-mcp        # Build
go run ./cmd/ynab-mcp/main.go  # Run
go vet ./...                    # Check for issues
```

No tests exist yet. No linter is configured.

## Environment Variables

- `YNAB_API_TOKEN` — Required (for stdio transport). YNAB personal access token.
- `TRANSPORT` — `"stdio"` (default) or `"http"`.
- `PORT` — HTTP server port (default `"8080"`, http transport only).
- HTTP transport reads the token from the `X-YNAB-Token` request header instead.

## Architecture

Three-layer design:

1. **Entry point** (`cmd/ynab-mcp/main.go`) — Creates the MCP server, registers tools, selects transport (stdio/http), and injects the YNAB client into request context.

2. **Tools layer** (`internal/tools/`) — MCP tool definitions and handlers. Each file defines tools for a YNAB resource (budgets, accounts, categories, transactions). `tools.go` registers all tools on the server. Handlers extract the YNAB client from context and delegate to the API client.

3. **YNAB API client** (`internal/ynab/`) — HTTP client wrapping the YNAB REST API (`https://api.ynab.com/v1`). `client.go` has the base HTTP logic. Each resource file adds typed methods. `types.go` holds response structs. `context.go` manages client storage in `context.Context`.

## Key Details

- YNAB amounts are in **milliunits** (1000 = $1.00).
- The `context.go` pattern stores/retrieves the `*ynab.Client` from `context.Context` using a typed key.
- Tool handlers in `internal/tools/` return `mcp.CallToolResult` with JSON-marshaled content or error text.
- When adding a new tool, also add a description of it to the README.
