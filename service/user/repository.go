package user

import "context"

type userRepository interface{
	GetAllUserRepo(ctx context.Context) ([]User, error)
}
