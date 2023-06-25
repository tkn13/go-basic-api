package adduser

import "github.com/ThaksinCK/go-basic-api.git/common"

type UserService struct {
	repo *UserRepository
}

func NewUserService() *UserService {
	repo := UserRepository{}
	return &UserService{
		repo: repo.NewUserRepository(),
	}
}

func (u *UserService) AddUser(request AddUserRequest) (err error) {
	if err = common.CheckEmptyInput(request.ID); err != nil {
		return err
	}
	if err = common.CheckEmptyInput(request.FirstName); err != nil {
		return err
	}
	if err = common.CheckEmptyInput(request.LastName); err != nil {
		return err
	}

	if err = u.repo.AddUser(request); err != nil {
		return err
	}
	return nil
}
