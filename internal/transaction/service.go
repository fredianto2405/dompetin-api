package transaction

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(request Request, userID int) error {
	entity := &Entity{
		Type:            request.Type,
		Amount:          request.Amount,
		Category:        request.Category,
		Description:     request.Description,
		TransactionDate: request.TransactionDate,
		UserID:          userID,
	}
	return s.repo.Save(entity)
}

func (s *Service) GetByTransactionDateAndUserID(startDate, endDate string, userID int) (*[]DetailResponse, error) {
	return s.repo.FindByTransactionDateAndUserID(startDate, endDate, userID)
}

func (s *Service) SummaryByTransactionDateAndUserID(startDate, endDate string, userID int) (*SummaryResponse, error) {
	return s.repo.SummaryByTransactionDateAndUserID(startDate, endDate, userID)
}

func (s *Service) Update(request Request, id int) error {
	entity := &Entity{
		ID:              id,
		Type:            request.Type,
		Amount:          request.Amount,
		Category:        request.Category,
		Description:     request.Description,
		TransactionDate: request.TransactionDate,
	}
	return s.repo.Update(entity)
}

func (s *Service) Delete(id int) error {
	return s.repo.Delete(id)
}
