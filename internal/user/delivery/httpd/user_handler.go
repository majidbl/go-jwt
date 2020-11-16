package httpd

import (
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/majidzarephysics/go-jwt/internal/domain"
	"github.com/majidzarephysics/go-jwt/pkg/jwt"
)

// UserHandler  represent the httphandler for article
type UserHandler struct {
	AUsecase domain.UserUsecase
}

// NewUserHandler will initialize the articles/ resources endpoint
func NewUserHandler(e *gin.Engine, us domain.UserUsecase) {
	handler := &UserHandler{
		AUsecase: us,
	}
	e.POST("/user/signup", handler.SignUp)
	e.POST("/user/signin", handler.SignIn)
}

// SignUp new User
func (m *UserHandler) SignUp(c *gin.Context) {

	var user domain.User
	c.BindJSON(&user)
	m.AUsecase.SignUp(user)
	c.JSON(200, gin.H{
		"message": "User Created Successfully",
		"User":    user,
		"device":  runtime.GOOS,
	})
}

// SignIn users
func (m *UserHandler) SignIn(c *gin.Context) {
	var loginform domain.LoginForm
	c.BindJSON(&loginform)
	user, err := m.AUsecase.SignIn(loginform.PassWord, loginform.Email)
	if err == nil {
		if user.PassWord == loginform.PassWord {
			jwtsig, errs := jwt.GenerateJWTSigned(user)
			if errs != nil {
				fmt.Println(errs)
			}
			/*fmt.Println(jwtsig)
			var c1 domain.User
			jwt.ParseJSONWebTokenClaims(jwtsig, &c1)
			fmt.Println(c1)
			*/
			c.JSON(200, gin.H{
				"message":   "User founded and logged",
				"User":      user,
				"JWT Token": jwtsig,
			})
		} else {
			c.JSON(200, gin.H{
				"message": "SignIN Failed",
				"err":     "Wrong Password",
			})
		}
	} else {
		c.JSON(200, gin.H{
			"message": "User Not Found",
			"err":     err.Error(),
		})
	}
}
