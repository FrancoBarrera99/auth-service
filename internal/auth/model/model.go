package model

type User struct {
	ID       string
	Username string
	Password string
	Email    string
}

type Credentials struct {
	Method string
	Data   map[string]interface{}
}

type UserRegister struct {
	Username string
	Password string
	Email    string
}

type UserResponse struct {
	ID       string
	Username string
	Email    string
}
