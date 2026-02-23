package ynab

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// GetTransactions returns transactions for a budget, optionally filtered by since_date.
func (c *Client) GetTransactions(budgetID string, sinceDate string) ([]TransactionDetail, error) {
	path := "/budgets/" + budgetID + "/transactions"
	if sinceDate != "" {
		path += "?since_date=" + sinceDate
	}

	var resp TransactionsResponse
	if err := c.doGet(path, &resp); err != nil {
		return nil, fmt.Errorf("getting transactions for budget %s: %w", budgetID, err)
	}
	return resp.Data.Transactions, nil
}

// CreateTransaction creates a new transaction in a budget.
func (c *Client) CreateTransaction(budgetID string, txn SaveTransaction) (*TransactionDetail, error) {
	wrapper := SaveTransactionWrapper{Transaction: txn}
	body, err := json.Marshal(wrapper)
	if err != nil {
		return nil, fmt.Errorf("marshaling transaction: %w", err)
	}

	var resp SaveTransactionsResponse
	if err := c.doPost("/budgets/"+budgetID+"/transactions", bytes.NewReader(body), &resp); err != nil {
		return nil, fmt.Errorf("creating transaction in budget %s: %w", budgetID, err)
	}
	return &resp.Data.Transaction, nil
}

// CreateTransactions creates multiple transactions in a budget in a single request.
func (c *Client) CreateTransactions(budgetID string, txns []SaveTransaction) ([]TransactionDetail, error) {
	wrapper := SaveTransactionsArrayWrapper{Transactions: txns}
	body, err := json.Marshal(wrapper)
	if err != nil {
		return nil, fmt.Errorf("marshaling transactions: %w", err)
	}

	var resp BulkSaveTransactionsResponse
	if err := c.doPost("/budgets/"+budgetID+"/transactions", bytes.NewReader(body), &resp); err != nil {
		return nil, fmt.Errorf("creating transactions in budget %s: %w", budgetID, err)
	}
	return resp.Data.Transactions, nil
}
