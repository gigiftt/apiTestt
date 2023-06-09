package repo

import (
	// "database/sql"

	"context"
	"log"

	user "apiTestt/service/user"

	"github.com/jmoiron/sqlx"
)

const(
	addUser = "INSERT INTO user (FirstName, LastName, Age, Email) VALUES (? , ? , ? , ?)"

	updateUser = "Update user SET FirstName = ?, LastName = ?, Age = ?, Email=? where LastName = ?"

	getUser = "Select * From user"

	deleteUser = "DELETE FROM user  WHERE LastName = ?"
)
type mysqluserRepository struct {
	db *sqlx.DB
}


func NewuserRepository(db *sqlx.DB) user.UserRepository {
	return &mysqluserRepository{
		db:db,
	}
}

func (us *mysqluserRepository) GetAllUserRepo(ctx context.Context) ([]user.UserModel, error) {

	var user []user.UserModel

	err := us.db.Select(&user,getUser)
	if err != nil {
		return nil, err
	}

	return user,nil
}


func (us *mysqluserRepository) CreateUserRepo(ctx context.Context,user user.UserModel) error{
	
	_,err := us.db.Exec(addUser,user.FirstName,user.LastName,user.Age,user.Email)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil

}

func (us *mysqluserRepository) UpdateUserRepo(ctx context.Context,user user.UserModel,name string) error {
	
	_,err := us.db.Exec(updateUser,user.FirstName,user.LastName,user.Age,user.Email, name)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}

func (us *mysqluserRepository) DeleteUserRepo(ctx context.Context,name string) error {
	
	_,err := us.db.Exec(deleteUser,name)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}

func (us *mysqluserRepository) GetUserByFormRepo(ctx context.Context,name string)([]user.UserModel, error){

	var user []user.UserModel
	query := getUser + "Where LastName = ?"

	err := us.db.Get(&user,query,name)
	if err != nil{
		log.Fatalln(err)
		return nil,err
	}

	return user,nil
}


func (us *mysqluserRepository) GetFile(ctx context.Context) error {
	
	return nil
}