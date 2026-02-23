package tools

import (
	"context"

	"github.com/blaberg/ynab-mcp/internal/ynab"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerCategoryTools(s *server.MCPServer) {
	getCategories := mcp.NewTool("get_categories",
		mcp.WithDescription("List all category groups and categories in a budget with their budgeted amounts, activity, and balances."),
		mcp.WithString("budget_id",
			mcp.Required(),
			mcp.Description("The ID of the budget"),
		),
	)
	s.AddTool(getCategories, getCategoriesHandler())
}

func getCategoriesHandler() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		client := ynab.FromContext(ctx)
		budgetID := mcp.ParseString(request, "budget_id", "")
		if budgetID == "" {
			return mcp.NewToolResultError("budget_id is required"), nil
		}

		categories, err := client.GetCategories(budgetID)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultJSON(map[string]any{"categories": categories})
	}
}
