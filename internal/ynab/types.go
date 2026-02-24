package ynab

// Amounts in the YNAB API are in "milliunits": 1000 milliunits = $1.00.
// For example, -10000 means -$10.00.

// BudgetSummary represents a budget overview.
type BudgetSummary struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	LastModifiedOn string `json:"last_modified_on"`
}

// BudgetDetail represents detailed budget information.
type BudgetDetail struct {
	ID             string          `json:"id"`
	Name           string          `json:"name"`
	LastModifiedOn string          `json:"last_modified_on"`
	Accounts       []Account       `json:"accounts"`
	CategoryGroups []CategoryGroup `json:"category_groups"`
}

// Account represents a YNAB account.
type Account struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Type             string `json:"type"`
	OnBudget         bool   `json:"on_budget"`
	Closed           bool   `json:"closed"`
	Balance          int64  `json:"balance"`
	ClearedBalance   int64  `json:"cleared_balance"`
	UnclearedBalance int64  `json:"uncleared_balance"`
	Deleted          bool   `json:"deleted"`
}

// CategoryGroup represents a group of categories.
type CategoryGroup struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Hidden     bool       `json:"hidden"`
	Deleted    bool       `json:"deleted"`
	Categories []Category `json:"categories"`
}

// Category represents a budget category.
type Category struct {
	ID              string `json:"id"`
	CategoryGroupID string `json:"category_group_id"`
	Name            string `json:"name"`
	Hidden          bool   `json:"hidden"`
	Budgeted        int64  `json:"budgeted"`
	Activity        int64  `json:"activity"`
	Balance         int64  `json:"balance"`
	Deleted         bool   `json:"deleted"`
}

// TransactionDetail represents a transaction.
type TransactionDetail struct {
	ID           string `json:"id"`
	Date         string `json:"date"`
	Amount       int64  `json:"amount"`
	Memo         string `json:"memo"`
	Cleared      string `json:"cleared"`
	Approved     bool   `json:"approved"`
	AccountID    string `json:"account_id"`
	AccountName  string `json:"account_name"`
	PayeeID      string `json:"payee_id"`
	PayeeName    string `json:"payee_name"`
	CategoryID   string `json:"category_id"`
	CategoryName string `json:"category_name"`
	Deleted      bool   `json:"deleted"`
}

// SaveTransaction represents the data needed to create a transaction.
type SaveTransaction struct {
	AccountID  string `json:"account_id"`
	Date       string `json:"date"`
	Amount     int64  `json:"amount"`
	PayeeName  string `json:"payee_name,omitempty"`
	CategoryID string `json:"category_id,omitempty"`
	Memo       string `json:"memo,omitempty"`
	Cleared    string `json:"cleared,omitempty"`
	Approved   bool   `json:"approved,omitempty"`
}

// Response wrappers matching YNAB API response structure.

type BudgetsResponse struct {
	Data struct {
		Budgets []BudgetSummary `json:"budgets"`
	} `json:"data"`
}

type BudgetDetailResponse struct {
	Data struct {
		Budget BudgetDetail `json:"budget"`
	} `json:"data"`
}

type AccountsResponse struct {
	Data struct {
		Accounts []Account `json:"accounts"`
	} `json:"data"`
}

type CategoriesResponse struct {
	Data struct {
		CategoryGroups []CategoryGroup `json:"category_groups"`
	} `json:"data"`
}

type TransactionsResponse struct {
	Data struct {
		Transactions []TransactionDetail `json:"transactions"`
	} `json:"data"`
}

type SaveTransactionWrapper struct {
	Transaction SaveTransaction `json:"transaction"`
}

type SaveTransactionsResponse struct {
	Data struct {
		Transaction TransactionDetail `json:"transaction"`
	} `json:"data"`
}

type SaveTransactionsArrayWrapper struct {
	Transactions []SaveTransaction `json:"transactions"`
}

type BulkSaveTransactionsResponse struct {
	Data struct {
		TransactionIDs     []string            `json:"transaction_ids"`
		Transactions       []TransactionDetail `json:"transactions"`
		DuplicateImportIDs []string            `json:"duplicate_import_ids"`
	} `json:"data"`
}

// MonthSummary represents a budget month overview.
type MonthSummary struct {
	Month      string `json:"month"`
	Income     int64  `json:"income"`
	Budgeted   int64  `json:"budgeted"`
	Activity   int64  `json:"activity"`
	ToBeBudget int64  `json:"to_be_budgeted"`
	Deleted    bool   `json:"deleted"`
}

// MonthDetail represents a single budget month with category breakdowns.
type MonthDetail struct {
	Month      string     `json:"month"`
	Income     int64      `json:"income"`
	Budgeted   int64      `json:"budgeted"`
	Activity   int64      `json:"activity"`
	ToBeBudget int64      `json:"to_be_budgeted"`
	Categories []Category `json:"categories"`
	Deleted    bool       `json:"deleted"`
}

type MonthsResponse struct {
	Data struct {
		Months []MonthSummary `json:"months"`
	} `json:"data"`
}

type MonthDetailResponse struct {
	Data struct {
		Month MonthDetail `json:"month"`
	} `json:"data"`
}
