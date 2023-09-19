package dto

type User struct {
	Username string
}

type CreateUser struct {
	Username string
	Password string
	Question string
	Answer   string
}
