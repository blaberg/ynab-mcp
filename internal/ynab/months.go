package ynab

import (
	"context"
	"fmt"
)

// GetMonths returns all budget months for a budget.
func (c *Client) GetMonths(ctx context.Context, budgetID string) ([]MonthSummary, error) {
	var resp MonthsResponse
	if err := c.doGet(ctx, "/budgets/"+budgetID+"/months", &resp); err != nil {
		return nil, fmt.Errorf("getting months for budget %s: %w", budgetID, err)
	}
	return resp.Data.Months, nil
}

// GetMonth returns a single budget month with category details.
func (c *Client) GetMonth(ctx context.Context, budgetID, month string) (*MonthDetail, error) {
	var resp MonthDetailResponse
	if err := c.doGet(ctx, "/budgets/"+budgetID+"/months/"+month, &resp); err != nil {
		return nil, fmt.Errorf("getting month %s for budget %s: %w", month, budgetID, err)
	}
	return &resp.Data.Month, nil
}
