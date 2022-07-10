package main

import (
	"fmt"
	"go_graphql/impls"
	"go_graphql/interfaces"
)

func main() {
	var userRepo interfaces.UserRepo
	userRepoImpl := impls.NewUserRepo()
	userRepo = &userRepoImpl
	fmt.Println(userRepo.Get(0))
}
