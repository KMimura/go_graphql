package interfaces

import (
	"go_graphql/entities"
)

type UserRepo interface {
	Get(int) *entities.User
	Find(string) []*entities.User
	Update(*entities.User) bool
	Create(string) *entities.User
	Delete(int) bool
}
