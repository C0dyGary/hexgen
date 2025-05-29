package configpage

import "github.com/C0dyGary/gift/src/pkg/domain"

func (r *Repository) SelectConfigPage(id uint) (*domain.ConfigPage, error) {	
	var configpage domain.ConfigPage
	if err := r.DB.First(&configpage, id).Error; err != nil {
		return nil, err
	}
	return &configpage, nil
}
