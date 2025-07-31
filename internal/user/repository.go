package user

import "github.com/jmoiron/sqlx"

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindByEmail(email string) (*Entity, error) {
	var entity Entity
	dataQuery := `select id, email, password from dompetin.users where email = $1`
	err := r.db.Get(&entity, dataQuery, email)
	return &entity, err
}

func (r *Repository) IsEmailExist(email string) (bool, error) {
	var isExist bool
	dataQuery := `select exists (select 1 from dompetin.users where email = $1)`
	err := r.db.Get(&isExist, dataQuery, email)
	if err != nil {
		return false, err
	}
	return isExist, nil
}

func (r *Repository) Save(email, password string) error {
	insertQuery := `insert into dompetin.users (email, password) values ($1, $2)`
	_, err := r.db.Exec(insertQuery, email, password)
	return err
}
