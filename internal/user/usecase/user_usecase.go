package usecase

import "github.com/majidzarephysics/go-jwt/internal/domain"

type userUsecase struct {
	UserRepo domain.UserRepository
}

// NewUserUsecase will create new an userUsecase object representation of domain.UserUsecase interface
func NewUserUsecase(a domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		UserRepo: a,
	}
}

// SignUp ....
func (a *userUsecase) SignUp(user domain.User) error {
	err := a.UserRepo.SignUp(user)
	if err != nil {
		return err
	}
	return nil
}

// SignIn ....
func (a *userUsecase) SignIn(password, email string) (domain.User, error) {
	u, err := a.UserRepo.SignIn(password, email)
	if err != nil {
		return domain.User{}, err
	}
	return u, nil
}

//Account ....
func (a *userUsecase) Account(username string) (domain.User, error) {
	user, err := a.UserRepo.Account(username)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
