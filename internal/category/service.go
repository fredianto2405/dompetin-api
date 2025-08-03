package category

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(name string, userID int) error {
	return s.repo.Save(name, userID)
}

func (s *Service) GetByUserID(userID int) (*[]DTO, error) {
	return s.repo.FindByUserID(userID)
}

func (s *Service) Update(id int, name string) error {
	return s.repo.Update(id, name)
}

func (s *Service) Delete(id int) error {
	return s.repo.Delete(id)
}
