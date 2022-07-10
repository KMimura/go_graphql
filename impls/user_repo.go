package impls

import (
	"go_graphql/entities"
)

type UserRepoImpl struct{}

func (*UserRepoImpl) Get(id int) *entities.User {
	return &entities.User{
		Id:   001,
		Name: "Steve Herrington",
	}
}

func (*UserRepoImpl) Find(name string) []*entities.User {
	user := &entities.User{
		Id:   002,
		Name: "Nancy Wheeler",
	}
	return []*entities.User{user}
}

func (*UserRepoImpl) Update(user *entities.User) bool {
	return true
}

func (*UserRepoImpl) Create(name string) *entities.User {
	return &entities.User{
		Id:   003,
		Name: "Dustin Henderson",
	}
}

func (*UserRepoImpl) Delete(id int) bool {
	return true
}

func NewUserRepo() UserRepoImpl {
	return UserRepoImpl{}
}
