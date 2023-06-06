package user

type User struct {
	FirstName string `json:"firstname" db:"FirstName"`
	LastName  string `json:"lastname" db:"LastName"`
	Age       int    `json:"age" db:"Age"`
	Email     string `json:"email" db:"Email"`
}