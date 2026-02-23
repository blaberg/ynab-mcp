package tools

import (
	"context"

	"github.com/blaberg/ynab-mcp/internal/ynab"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerBudgetTools(s *server.MCPServer) {
	listBudgets := mcp.NewTool("list_budgets",
		mcp.WithDescription("List all budgets in the YNAB account. Returns budget IDs and names. Use this first to discover available budget IDs."),
	)
	s.AddTool(listBudgets, listBudgetsHandler())

	getBudget := mcp.NewTool("get_budget",
		mcp.WithDescription("Get detailed information about a specific budget including accounts, categories, and category groups."),
		mcp.WithString("budget_id",
			mcp.Required(),
			mcp.Description("The ID of the budget"),
		),
	)
	s.AddTool(getBudget, getBudgetHandler())
}

func listBudgetsHandler() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		client := ynab.FromContext(ctx)
		budgets, err := client.GetBudgets()
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultJSON(map[string]any{"budgets": budgets})
	}
}

func getBudgetHandler() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		client := ynab.FromContext(ctx)
		budgetID := mcp.ParseString(request, "budget_id", "")
		if budgetID == "" {
			return mcp.NewToolResultError("budget_id is required"), nil
		}

		budget, err := client.GetBudget(budgetID)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultJSON(budget)
	}
}
