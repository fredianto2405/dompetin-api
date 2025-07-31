package response

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

type Meta struct {
	Total     int  `json:"total"`
	Page      int  `json:"page"`
	PageCount int  `json:"page_count"`
	HasNext   bool `json:"has_next"`
	HasPrev   bool `json:"has_prev"`
}

type Pagination struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}
