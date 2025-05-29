package configpage

import "github.com/C0dyGary/gift/src/pkg/domain"

func (s *Service) UpdateConfigPage(id uint, newData map[string]interface{}) (*domain.ConfigPage, error) {	
	return s.Repository.UpdateConfigPage(id, newData)
}
