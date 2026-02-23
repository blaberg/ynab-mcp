package ynab

import "fmt"

// GetCategories returns all category groups (with nested categories) for a budget.
func (c *Client) GetCategories(budgetID string) ([]CategoryGroup, error) {
	var resp CategoriesResponse
	if err := c.doGet("/budgets/"+budgetID+"/categories", &resp); err != nil {
		return nil, fmt.Errorf("getting categories for budget %s: %w", budgetID, err)
	}
	return resp.Data.CategoryGroups, nil
}
