package postgresql

import (
	"fmt"

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
	fmt.Println("postgresql:", email)
	if err := m.Conn.Where("email = ?", email).First(&user).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

/*
//GetAll new User
func (m *mysqlUserRepository) GetAll(c *gin.Context) {
	// Get all records
	var allUser []domain.User
	_ = m.Conn.Find(&allUser)
	fmt.Println(allUser)
	c.JSON(http.StatusOK, gin.H{"message": "hey", "result": allUser})

}

// Delete ByID Will be use for delete specific User
func (m *mysqlUserRepository) Delete(c *gin.Context) {
	// delete specific user by Student ID
	email := c.Param("email")
	var user domain.User
	// Get first matched record
	m.Conn.Where(&domain.User{Email: email}).First(&user)
	m.Conn.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "user deleted succesfully", "status": http.StatusOK})
}

// Update will be ues to update User
func (m *mysqlUserRepository) Update(c *gin.Context) {
	// return specific user by User ID

	username := c.Param("username")
	var oldUser domain.User
	var newUser domain.User
	c.BindJSON(&newUser)
	// Get first matched record
	m.Conn.Where(&domain.User{UserName: username}).First(&oldUser)
	oldUser.UserName = newUser.UserName
	oldUser.Email = newUser.Email
	oldUser.PassWord = newUser.PassWord
	m.Conn.Save(&oldUser)
	c.JSON(http.StatusOK, oldUser)

}
*/
