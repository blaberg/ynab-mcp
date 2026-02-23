package tools

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/blaberg/ynab-mcp/internal/ynab"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerTransactionTools(s *server.MCPServer) {
	getTransactions := mcp.NewTool("get_transactions",
		mcp.WithDescription("Get transactions for a budget. Optionally filter by date. Amounts are in milliunits (1000 = $1.00)."),
		mcp.WithString("budget_id",
			mcp.Required(),
			mcp.Description("The ID of the budget"),
		),
		mcp.WithString("since_date",
			mcp.Description("Only return transactions on or after this date (ISO format YYYY-MM-DD)"),
		),
	)
	s.AddTool(getTransactions, getTransactionsHandler())

	createTransaction := mcp.NewTool("create_transaction",
		mcp.WithDescription("Create a new transaction in a budget. Amount is in milliunits (1000 = $1.00, use negative for outflows)."),
		mcp.WithString("budget_id",
			mcp.Required(),
			mcp.Description("The ID of the budget"),
		),
		mcp.WithString("account_id",
			mcp.Required(),
			mcp.Description("The ID of the account for this transaction"),
		),
		mcp.WithString("date",
			mcp.Required(),
			mcp.Description("The transaction date in ISO format (YYYY-MM-DD)"),
		),
		mcp.WithNumber("amount",
			mcp.Required(),
			mcp.Description("The transaction amount in milliunits (1000 = $1.00, use negative for outflows)"),
		),
		mcp.WithString("payee_name",
			mcp.Description("The payee name"),
		),
		mcp.WithString("category_id",
			mcp.Description("The category ID for this transaction"),
		),
		mcp.WithString("memo",
			mcp.Description("A memo for the transaction"),
		),
		mcp.WithString("cleared",
			mcp.Description("The cleared status"),
			mcp.Enum("cleared", "uncleared", "reconciled"),
		),
		mcp.WithBoolean("approved",
			mcp.Description("Whether the transaction is approved"),
		),
	)
	s.AddTool(createTransaction, createTransactionHandler())

	createTransactions := mcp.NewTool("create_transactions",
		mcp.WithDescription("Create multiple transactions in a budget in a single request. Each transaction amount is in milliunits (1000 = $1.00, use negative for outflows)."),
		mcp.WithString("budget_id",
			mcp.Required(),
			mcp.Description("The ID of the budget"),
		),
		mcp.WithArray("transactions",
			mcp.Required(),
			mcp.Description("Array of transactions to create"),
			mcp.Items(map[string]any{
				"type": "object",
				"properties": map[string]any{
					"account_id":  map[string]any{"type": "string", "description": "The ID of the account for this transaction"},
					"date":        map[string]any{"type": "string", "description": "The transaction date in ISO format (YYYY-MM-DD)"},
					"amount":      map[string]any{"type": "number", "description": "The transaction amount in milliunits (1000 = $1.00, use negative for outflows)"},
					"payee_name":  map[string]any{"type": "string", "description": "The payee name"},
					"category_id": map[string]any{"type": "string", "description": "The category ID for this transaction"},
					"memo":        map[string]any{"type": "string", "description": "A memo for the transaction"},
					"cleared":     map[string]any{"type": "string", "description": "The cleared status", "enum": []string{"cleared", "uncleared", "reconciled"}},
					"approved":    map[string]any{"type": "boolean", "description": "Whether the transaction is approved"},
				},
				"required": []string{"account_id", "date", "amount"},
			}),
		),
	)
	s.AddTool(createTransactions, createTransactionsHandler())
}

func getTransactionsHandler() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		client := ynab.FromContext(ctx)
		budgetID := mcp.ParseString(request, "budget_id", "")
		if budgetID == "" {
			return mcp.NewToolResultError("budget_id is required"), nil
		}

		sinceDate := mcp.ParseString(request, "since_date", "")

		transactions, err := client.GetTransactions(budgetID, sinceDate)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultJSON(map[string]any{"transactions": transactions})
	}
}

func createTransactionsHandler() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		client := ynab.FromContext(ctx)
		budgetID := mcp.ParseString(request, "budget_id", "")
		if budgetID == "" {
			return mcp.NewToolResultError("budget_id is required"), nil
		}

		rawTxns := mcp.ParseArgument(request, "transactions", nil)
		if rawTxns == nil {
			return mcp.NewToolResultError("transactions is required"), nil
		}

		// Re-marshal and unmarshal to convert the raw interface into typed structs.
		rawJSON, err := json.Marshal(rawTxns)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid transactions data: %v", err)), nil
		}

		var txns []ynab.SaveTransaction
		if err := json.Unmarshal(rawJSON, &txns); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid transactions format: %v", err)), nil
		}

		if len(txns) == 0 {
			return mcp.NewToolResultError("transactions array must not be empty"), nil
		}

		for i, txn := range txns {
			if txn.AccountID == "" {
				return mcp.NewToolResultError(fmt.Sprintf("transaction[%d]: account_id is required", i)), nil
			}
			if txn.Date == "" {
				return mcp.NewToolResultError(fmt.Sprintf("transaction[%d]: date is required", i)), nil
			}
		}

		result, err := client.CreateTransactions(budgetID, txns)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultJSON(map[string]any{"transactions": result})
	}
}

func createTransactionHandler() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		client := ynab.FromContext(ctx)
		budgetID := mcp.ParseString(request, "budget_id", "")
		if budgetID == "" {
			return mcp.NewToolResultError("budget_id is required"), nil
		}

		txn := ynab.SaveTransaction{
			AccountID:  mcp.ParseString(request, "account_id", ""),
			Date:       mcp.ParseString(request, "date", ""),
			Amount:     mcp.ParseInt64(request, "amount", 0),
			PayeeName:  mcp.ParseString(request, "payee_name", ""),
			CategoryID: mcp.ParseString(request, "category_id", ""),
			Memo:       mcp.ParseString(request, "memo", ""),
			Cleared:    mcp.ParseString(request, "cleared", ""),
			Approved:   mcp.ParseBoolean(request, "approved", false),
		}

		if txn.AccountID == "" {
			return mcp.NewToolResultError("account_id is required"), nil
		}
		if txn.Date == "" {
			return mcp.NewToolResultError("date is required"), nil
		}

		result, err := client.CreateTransaction(budgetID, txn)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultJSON(result)
	}
}
