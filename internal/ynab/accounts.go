package ynab

import (
	"context"
	"fmt"
)

// GetAccounts returns all accounts for a budget.
func (c *Client) GetAccounts(ctx context.Context, budgetID string) ([]Account, error) {
	var resp AccountsResponse
	if err := c.doGet(ctx, "/budgets/"+budgetID+"/accounts", &resp); err != nil {
		return nil, fmt.Errorf("getting accounts for budget %s: %w", budgetID, err)
	}
	return resp.Data.Accounts, nil
}
