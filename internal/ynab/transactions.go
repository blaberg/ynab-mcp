package ynab

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

// GetTransactions returns transactions for a budget, optionally filtered by since_date.
func (c *Client) GetTransactions(ctx context.Context, budgetID string, sinceDate string) ([]TransactionDetail, error) {
	path := "/budgets/" + budgetID + "/transactions"
	if sinceDate != "" {
		path += "?since_date=" + sinceDate
	}

	var resp TransactionsResponse
	if err := c.doGet(ctx, path, &resp); err != nil {
		return nil, fmt.Errorf("getting transactions for budget %s: %w", budgetID, err)
	}
	return resp.Data.Transactions, nil
}

// CreateTransaction creates a new transaction in a budget.
func (c *Client) CreateTransaction(ctx context.Context, budgetID string, txn SaveTransaction) (*TransactionDetail, error) {
	wrapper := SaveTransactionWrapper{Transaction: txn}
	body, err := json.Marshal(wrapper)
	if err != nil {
		return nil, fmt.Errorf("marshaling transaction: %w", err)
	}

	var resp SaveTransactionsResponse
	if err := c.doPost(ctx, "/budgets/"+budgetID+"/transactions", bytes.NewReader(body), &resp); err != nil {
		return nil, fmt.Errorf("creating transaction in budget %s: %w", budgetID, err)
	}
	return &resp.Data.Transaction, nil
}

// UpdateTransaction updates an existing transaction in a budget.
func (c *Client) UpdateTransaction(ctx context.Context, budgetID string, transactionID string, txn SaveTransaction) (*TransactionDetail, error) {
	wrapper := SaveTransactionWrapper{Transaction: txn}
	body, err := json.Marshal(wrapper)
	if err != nil {
		return nil, fmt.Errorf("marshaling transaction: %w", err)
	}

	var resp SaveTransactionsResponse
	if err := c.doPut(ctx, "/budgets/"+budgetID+"/transactions/"+transactionID, bytes.NewReader(body), &resp); err != nil {
		return nil, fmt.Errorf("updating transaction %s in budget %s: %w", transactionID, budgetID, err)
	}
	return &resp.Data.Transaction, nil
}

// CreateTransactions creates multiple transactions in a budget in a single request.
func (c *Client) CreateTransactions(ctx context.Context, budgetID string, txns []SaveTransaction) ([]TransactionDetail, error) {
	wrapper := SaveTransactionsArrayWrapper{Transactions: txns}
	body, err := json.Marshal(wrapper)
	if err != nil {
		return nil, fmt.Errorf("marshaling transactions: %w", err)
	}

	var resp BulkSaveTransactionsResponse
	if err := c.doPost(ctx, "/budgets/"+budgetID+"/transactions", bytes.NewReader(body), &resp); err != nil {
		return nil, fmt.Errorf("creating transactions in budget %s: %w", budgetID, err)
	}
	return resp.Data.Transactions, nil
}
