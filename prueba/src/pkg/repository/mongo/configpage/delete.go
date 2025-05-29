package configpage

import "github.com/C0dyGary/gift/src/pkg/domain"

func (r *Repository) DeleteConfigPage(id uint) error {	
	if err := r.DB.Delete(&domain.ConfigPage{}, id).Error; err != nil {
		return err
	}
	return nil
}
