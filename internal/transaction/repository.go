package transaction

import "github.com/jmoiron/sqlx"

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(e *Entity) error {
	insertQuery := `insert into dompetin.transactions(type, amount, category, description, transaction_date, user_id) 
		values(:type, :amount, :category, :description, :transaction_date, :user_id)`
	_, err := r.db.NamedExec(insertQuery, e)
	return err
}

func (r *Repository) FindByTransactionDateAndUserID(startDate, endDate string, userID int) (*[]DetailResponse, error) {
	dataQuery := `select id,
		  type,
		  amount,
		  category,
		  description,
		  transaction_date
		from dompetin.transactions
		where user_id = $1
		and transaction_date between $2 and $3
		order by transaction_date desc`

	var details []DetailResponse
	err := r.db.Select(&details, dataQuery, userID, startDate, endDate)
	return &details, err
}

func (r *Repository) SummaryByTransactionDateAndUserID(startDate, endDate string, userID int) (*SummaryResponse, error) {
	dataQuery := `with summary as (
		  select coalesce(sum(case when type = 'income' then amount else 0 end), 0) as income,
			coalesce(sum(case when type = 'expense' then amount else 0 end), 0) as expense
		  from dompetin.transactions
		  where user_id = $1
		  and transaction_date between $2 and $3
		)
		select income, 
		  expense, 
		  (income - expense) as balance
		from summary`

	var summary SummaryResponse
	err := r.db.Get(&summary, dataQuery, userID, startDate, endDate)
	return &summary, err
}

func (r *Repository) Update(e *Entity) error {
	updateQuery := `update dompetin.transactions
		set type = :type,
		  amount = :amount,
		  category = :category,
		  description = :description,
		  transaction_date = :transaction_date
		where id = :id`
	_, err := r.db.NamedExec(updateQuery, e)
	return err
}

func (r *Repository) Delete(id int) error {
	deleteQuery := `delete from dompetin.transactions where id = $1`
	_, err := r.db.Exec(deleteQuery, id)
	return err
}
