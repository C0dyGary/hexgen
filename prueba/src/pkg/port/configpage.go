package port

import "github.com/C0dyGary/gift/src/pkg/domain"

type ServiceConfigPage interface {
    CreateConfigPage(domain.ConfigPage) (*domain.ConfigPage,error)
    ReadConfigPage(id uint) (*domain.ConfigPage,error)
    ListConfigPage() (*[]domain.ConfigPage,error)
    UpdateConfigPage(id uint, newData map[string]interface{}) (*domain.ConfigPage,error)
    DeleteConfigPage(id uint) error

}

type RepositoryConfigPage interface {
    InsertConfigPage(domain.ConfigPage) (*domain.ConfigPage,error)
    SelectConfigPage(id uint) (*domain.ConfigPage,error)
    SelectAllConfigPage() (*[]domain.ConfigPage,error)
    UpdateConfigPage(id uint, newData map[string]interface{}) (*domain.ConfigPage,error)
    DeleteConfigPage(id uint) error
}
