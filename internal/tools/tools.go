package tools

import (
	"github.com/mark3labs/mcp-go/server"
)

// RegisterTools registers all YNAB MCP tools on the server.
func RegisterTools(s *server.MCPServer) {
	registerBudgetTools(s)
	registerAccountTools(s)
	registerCategoryTools(s)
	registerTransactionTools(s)
}
