package ynab

import "fmt"

// GetAccounts returns all accounts for a budget.
func (c *Client) GetAccounts(budgetID string) ([]Account, error) {
	var resp AccountsResponse
	if err := c.doGet("/budgets/"+budgetID+"/accounts", &resp); err != nil {
		return nil, fmt.Errorf("getting accounts for budget %s: %w", budgetID, err)
	}
	return resp.Data.Accounts, nil
}
