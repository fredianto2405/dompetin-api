package category

import "github.com/jmoiron/sqlx"

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(name string, userID int) error {
	insertQuery := `insert into dompetin.categories(name, user_id) values($1, $2)`
	_, err := r.db.Exec(insertQuery, name, userID)
	return err
}

func (r *Repository) FindByUserID(userID int) (*[]DTO, error) {
	var categories []DTO
	dataQuery := `select id, name from dompetin.categories where user_id = $1`
	err := r.db.Select(&categories, dataQuery, userID)
	return &categories, err
}

func (r *Repository) Update(id int, name string) error {
	updateQuery := `update dompetin.categories set name = $1 where id = $2`
	_, err := r.db.Exec(updateQuery, name, id)
	return err
}

func (r *Repository) Delete(id int) error {
	deleteQuery := `delete from dompetin.categories where id = $1`
	_, err := r.db.Exec(deleteQuery, id)
	return err
}
