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
	GetUserByFormServ(ctx context.Context,name string)([]UserModel, error)
	GetFile(ctx context.Context) error
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

func (usr *UserService) GetUserByFormServ(ctx context.Context,name string)([]UserModel, error){
	user,err := usr.userRepository.GetUserByFormRepo(ctx,name)

	if err!= nil{
		log.Fatalln(err)
		return nil,err
	}
	return user,nil
}


func (usr *UserService) GetFile(ctx context.Context) error {
	
	return nil
}