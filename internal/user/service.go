package user

import passwordutil "dompetin-api/pkg/password"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetByEmail(email string) (*Entity, error) {
	return s.repo.FindByEmail(email)
}

func (s *Service) IsEmailExist(email string) (bool, error) {
	return s.repo.IsEmailExist(email)
}

func (s *Service) Create(email, password string) error {
	hashedPassword, err := passwordutil.HashPassword(password)
	if err != nil {
		return err
	}
	return s.repo.Save(email, hashedPassword)
}
