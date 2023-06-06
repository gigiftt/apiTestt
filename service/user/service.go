package user

import (
	"context"
	"log"
)

type UserServiceInt interface {
	GetAllUserServ(ctx context.Context) ([]UserModel, error)
	CreateUserServ(ctx context.Context, user UserModel) error
	UpdateUserServ(ctx context.Context, user UserModel,name string) error
	DeleteUserServ(ctx context.Context, name string) error
	// GetUserByPathServ(ctx context.Context,name string)([]UserModel, error)
	// GetUserByQueryServ(ctx context.Context,name string)([]UserModel, error)
	// GetUserByRawServ(ctx context.Context,name string)([]UserModel, error)
	// GetUserByFormServ(ctx context.Context,name string)([]UserModel, error)
}

type UserService struct{
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) UserServiceInt {
	return &UserService{
		userRepository: userRepository,
	}
}

func (usr *UserService) GetAllUserServ(ctx context.Context) ([]UserModel, error){
	users,err := usr.userRepository.GetAllUserRepo(ctx)
	if err != nil {
		log.Fatalln(err)
		return nil,err
	}

	return users,nil
}

func (usr *UserService) CreateUserServ(ctx context.Context, user UserModel) error{
	err := usr.userRepository.CreateUserRepo(ctx,user)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}

func (usr *UserService) UpdateUserServ(ctx context.Context, user UserModel,name string) error{
	err := usr.userRepository.UpdateUserRepo(ctx, user, name)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}

func (usr *UserService) DeleteUserServ(ctx context.Context, name string) error{
	err := usr.userRepository.DeleteUserRepo(ctx,name)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}

// func (usr *UserService) GetUserByPathServ(ctx context.Context,name string)([]UserModel, error){

// }

// func (usr *UserService) GetUserByQueryServ(ctx context.Context,name string)([]UserModel, error){

// }

// func (usr *UserService) GetUserByRawServ(ctx context.Context,name string)([]UserModel, error){

// }

// func (usr *UserService) GetUserByFormServ(ctx context.Context,name string)([]UserModel, error){

// }
