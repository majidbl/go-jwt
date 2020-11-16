package domain

import "github.com/jinzhu/gorm"

// User Struct
type User struct {
	gorm.Model
	UserName string `json:"username"`
	PassWord string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Name     string `json:"name"`
}

// LoginForm ....
type LoginForm struct {
	PassWord string `json:"password"`
	Email    string `json:"email"`
}

// UserRepository represent the User's Repository
type UserRepository interface {
	SignUp(newuser User) error
	SignIn(username, email string) (User, error)
	Account(username string) (User, error)
}

// UserUsecase represent the User's UseCase
type UserUsecase interface {
	SignUp(newuser User) error
	SignIn(username, email string) (User, error)
	Account(username string) (User, error)
}
