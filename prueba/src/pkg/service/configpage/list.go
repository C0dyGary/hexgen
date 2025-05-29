package configpage

import "github.com/C0dyGary/gift/src/pkg/domain"

func (s *Service) ListConfigPage() (*[]domain.ConfigPage, error) {	
	return s.Repository.SelectAllConfigPage()
}
