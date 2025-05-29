package configpage

import "github.com/C0dyGary/gift/src/pkg/domain"

func (s *Service) DeleteConfigPage(id uint) error {	
	return s.Repository.DeleteConfigPage(id)
}
