package user

import (
	"context"
	// "golang.org/x/net/context/ctxhttp"
)

type UserRepository interface{
	GetAllUserRepo(ctx context.Context) ([]UserModel, error)
	CreateUserRepo(ctx context.Context,user UserModel) error
	UpdateUserRepo(ctx context.Context,user UserModel,name string) error
	DeleteUserRepo(ctx context.Context,name string) error
	GetUserByFormRepo(ctx context.Context,name string)([]UserModel, error)
	GetFile(ctx context.Context) error
}
