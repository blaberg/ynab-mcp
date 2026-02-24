package tools

import (
	"context"

	"github.com/blaberg/ynab-mcp/internal/ynab"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerAccountTools(s *server.MCPServer) {
	getAccounts := mcp.NewTool("get_accounts",
		mcp.WithDescription("List all accounts in a budget with their balances."),
		mcp.WithString("budget_id",
			mcp.Required(),
			mcp.Description("The ID of the budget"),
		),
	)
	s.AddTool(getAccounts, getAccountsHandler())
}

func getAccountsHandler() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		client := ynab.FromContext(ctx)
		if client == nil {
			return mcp.NewToolResultError("YNAB API token not configured"), nil
		}
		budgetID := mcp.ParseString(request, "budget_id", "")
		if budgetID == "" {
			return mcp.NewToolResultError("budget_id is required"), nil
		}

		accounts, err := client.GetAccounts(ctx, budgetID)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultJSON(map[string]any{"accounts": accounts})
	}
}
