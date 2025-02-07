package port

import "github.com/C0dyGary/prueba/src/pkg/domain"

type ServiceUser interface {
    CreateUser(domain.User) (*domain.User,error)
    ReadUser(id uint) (*domain.User,error)
    ListUser() (*[]domain.User,error)
    UpdateUser(id uint, newData map[string]interface{}) (*domain.User,error)
    DeleteUser(id uint) error

}

type RepositoryUser interface {
    InsertUser(domain.User) (*domain.User,error)
    SelectUser(id uint) (*domain.User,error)
    SelectAllUser() (*[]domain.User,error)
    UpdateUser(id uint, newData map[string]interface{}) (*domain.User,error)
    DeleteUser(id uint) error
}
