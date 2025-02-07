package port

import "github.com/C0dyGary/prueba/src/pkg/domain"

type ServiceCategory interface {
    CreateCategory(domain.Category) (*domain.Category,error)
    ReadCategory(id uint) (*domain.Category,error)
    ListCategory() (*[]domain.Category,error)
    UpdateCategory(id uint, newData map[string]interface{}) (*domain.Category,error)
    DeleteCategory(id uint) error

}

type RepositoryCategory interface {
    InsertCategory(domain.Category) (*domain.Category,error)
    SelectCategory(id uint) (*domain.Category,error)
    SelectAllCategory() (*[]domain.Category,error)
    UpdateCategory(id uint, newData map[string]interface{}) (*domain.Category,error)
    DeleteCategory(id uint) error
}
