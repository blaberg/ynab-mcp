package ynab

import (
	"context"
	"fmt"
)

// GetCategories returns all category groups (with nested categories) for a budget.
func (c *Client) GetCategories(ctx context.Context, budgetID string) ([]CategoryGroup, error) {
	var resp CategoriesResponse
	if err := c.doGet(ctx, "/budgets/"+budgetID+"/categories", &resp); err != nil {
		return nil, fmt.Errorf("getting categories for budget %s: %w", budgetID, err)
	}
	return resp.Data.CategoryGroups, nil
}
