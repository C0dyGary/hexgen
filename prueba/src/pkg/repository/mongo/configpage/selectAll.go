package configpage

import "github.com/C0dyGary/gift/src/pkg/domain"

func (r *Repository) SelectAllConfigPage() (*[]domain.ConfigPage, error) {	
	var configpage []domain.ConfigPage
	if err := r.DB.Find(&configpage).Error; err != nil {
		return nil, err
	}
	return &configpage, nil
}
