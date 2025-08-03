package transaction

type Entity struct {
	ID              int    `db:"id"`
	Type            string `db:"type"`
	Amount          int    `db:"amount"`
	Category        string `db:"category"`
	Description     string `db:"description"`
	TransactionDate string `db:"transaction_date"`
	UserID          int    `db:"user_id"`
}

type Request struct {
	Type            string `json:"type"`
	Amount          int    `json:"amount"`
	Category        string `json:"category"`
	Description     string `json:"description"`
	TransactionDate string `json:"transaction_date"`
}

type DetailResponse struct {
	ID              int    `json:"id" db:"id"`
	Type            string `json:"type" db:"type"`
	Amount          int    `json:"amount" db:"amount"`
	Category        string `json:"category" db:"category"`
	Description     string `json:"description" db:"description"`
	TransactionDate string `json:"transaction_date" db:"transaction_date"`
}

type SummaryResponse struct {
	Income  int `json:"income" db:"income"`
	Expense int `json:"expense" db:"expense"`
	Balance int `json:"balance" db:"balance"`
}

type HistoryResponse struct {
	Summary *SummaryResponse  `json:"summary"`
	Details *[]DetailResponse `json:"details"`
}
