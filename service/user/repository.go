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
	// GetUserByPathRepo(ctx context.Context,name string)([]UserModel, error)
	// GetUserByQueryRepo(ctx context.Context,name string)([]UserModel, error)
	// GetUserByRawRepo(ctx context.Context,name string)([]UserModel, error)
	// GetUserByFormRepo(ctx context.Context,name string)([]UserModel, error)
}
