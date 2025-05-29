package configpage

import "github.com/C0dyGary/gift/src/pkg/domain"

func (s *Service) ReadConfigPage(id uint) (*domain.ConfigPage, error) {	
	return s.Repository.SelectConfigPage(id)
}
