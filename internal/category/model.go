package category

type DTO struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type Request struct {
	Name string `json:"name"`
}
