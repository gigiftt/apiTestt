package user

type UserModel struct {
	FirstName string `json:"firstname" db:"FirstName"`
	LastName  string `json:"lastname" db:"LastName"  query:"name"`
	Age       int    `json:"age" db:"Age"`
	Email     string `json:"email" db:"Email"`
}