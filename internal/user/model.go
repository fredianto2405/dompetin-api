package user

type Entity struct {
	ID       int    `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
}
