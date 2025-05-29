package configpage

import "github.com/C0dyGary/gift/src/pkg/domain"

func (r *Repository) InsertConfigPage(configpage domain.ConfigPage) (*domain.ConfigPage, error) {	
	if err := r.DB.Create(&configpage).Error; err != nil {
		return nil, err
	}
	var newConfigPage domain.ConfigPage
	if err :=r.DB.First(&newConfigPage, configpage.ID).Error; err != nil {
		return nil, err
	}
	return &newConfigPage, nil
}
