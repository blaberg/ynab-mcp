package tools

import (
	"context"

	"github.com/blaberg/ynab-mcp/internal/ynab"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerMonthTools(s *server.MCPServer) {
	getMonths := mcp.NewTool("get_budget_months",
		mcp.WithDescription("List all budget months for a budget. Returns monthly summaries with income, budgeted, activity, and to-be-budgeted amounts."),
		mcp.WithString("budget_id",
			mcp.Required(),
			mcp.Description("The ID of the budget"),
		),
	)
	s.AddTool(getMonths, getBudgetMonthsHandler())

	getMonth := mcp.NewTool("get_budget_month",
		mcp.WithDescription("Get a single budget month with category breakdowns. Returns income, budgeted, activity, to-be-budgeted, and per-category details."),
		mcp.WithString("budget_id",
			mcp.Required(),
			mcp.Description("The ID of the budget"),
		),
		mcp.WithString("month",
			mcp.Required(),
			mcp.Description("The budget month in ISO format (e.g. 2024-01-01)"),
		),
	)
	s.AddTool(getMonth, getBudgetMonthHandler())
}

func getBudgetMonthsHandler() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		client := ynab.FromContext(ctx)
		if client == nil {
			return mcp.NewToolResultError("YNAB API token not configured"), nil
		}
		budgetID := mcp.ParseString(request, "budget_id", "")
		if budgetID == "" {
			return mcp.NewToolResultError("budget_id is required"), nil
		}

		months, err := client.GetMonths(ctx, budgetID)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultJSON(map[string]any{"months": months})
	}
}

func getBudgetMonthHandler() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		client := ynab.FromContext(ctx)
		if client == nil {
			return mcp.NewToolResultError("YNAB API token not configured"), nil
		}
		budgetID := mcp.ParseString(request, "budget_id", "")
		if budgetID == "" {
			return mcp.NewToolResultError("budget_id is required"), nil
		}
		month := mcp.ParseString(request, "month", "")
		if month == "" {
			return mcp.NewToolResultError("month is required"), nil
		}

		detail, err := client.GetMonth(ctx, budgetID, month)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultJSON(detail)
	}
}
