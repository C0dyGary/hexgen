package configpage

import "github.com/C0dyGary/gift/src/pkg/domain"

func (s *Service) CreateConfigPage(configpage domain.ConfigPage) (*domain.ConfigPage, error) {	
	return s.Repository.InsertConfigPage(configpage)
}
