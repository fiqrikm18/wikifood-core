package user

import "WikfoodCore/internal/repositories/user"

type UserService struct {
	UserRepository user.IUserRepository
}

func New(repository user.IUserRepository) *UserService {
	return &UserService{
		UserRepository: repository,
	}
}
