package ynab

import (
	"context"
	"fmt"
)

// GetBudgets returns all budgets for the authenticated user.
func (c *Client) GetBudgets(ctx context.Context) ([]BudgetSummary, error) {
	var resp BudgetsResponse
	if err := c.doGet(ctx, "/budgets", &resp); err != nil {
		return nil, fmt.Errorf("getting budgets: %w", err)
	}
	return resp.Data.Budgets, nil
}

// GetBudget returns detailed information for a specific budget.
func (c *Client) GetBudget(ctx context.Context, budgetID string) (*BudgetDetail, error) {
	var resp BudgetDetailResponse
	if err := c.doGet(ctx, "/budgets/"+budgetID, &resp); err != nil {
		return nil, fmt.Errorf("getting budget %s: %w", budgetID, err)
	}
	return &resp.Data.Budget, nil
}
