package postgresql

import (
	"github.com/majidzarephysics/go-jwt/internal/domain"
	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	Conn *gorm.DB
}

// NewMysqlUserRepository will create an object that represent the User.Repository interface
func NewMysqlUserRepository(Conn *gorm.DB) domain.UserRepository {
	return &mysqlUserRepository{Conn: Conn}
}

// SignUp new User
func (m *mysqlUserRepository) SignUp(user domain.User) error {

	if err := m.Conn.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
func (m *mysqlUserRepository) SignIn(password, email string) (domain.User, error) {

	var user domain.User
	//fmt.Println("postgresql:", email)
	if err := m.Conn.Where("email = ?", email).First(&user).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (m *mysqlUserRepository) Account(username string) (domain.User, error) {
	var user domain.User
	if err := m.Conn.Where("user_name = ?", username).First(&user).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}
