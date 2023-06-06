package user

import "context"

type UserRepository interface{
	GetAllUserRepo(ctx context.Context) ([]User, error)
}
