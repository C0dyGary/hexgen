package configpage

import "github.com/C0dyGary/gift/src/pkg/domain"

func (r *Repository) UpdateConfigPage(id uint, newData map[string]interface{}) (*domain.ConfigPage, error) {	
	var configpage domain.ConfigPage
	if err := r.DB.First(&configpage, id).Error; err != nil {
		return nil, err
	}
	if err := r.DB.Model(&configpage).Updates(newData).Error; err != nil {
		return nil, err
	}
	return &configpage, nil
}
