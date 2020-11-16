package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	httpdelivery "github.com/majidzarephysics/go-jwt/internal/user/delivery/httpd"
	"github.com/majidzarephysics/go-jwt/internal/user/repository/postgresql"
	"github.com/majidzarephysics/go-jwt/internal/user/usecase"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	/*user := domain.User{
		UserName: "majid72bl",
		PassWord: "123456",
		Email:    "test",
		Role:     "admin",
		Name:     "Majid",
	}*/
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbUser := os.Getenv("db_user")
	dbPassword := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	//dbDriver := os.Getenv("db_driver")
	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=5432 sslmode=disable TimeZone=Asia/Tehran"
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	ur := postgresql.NewMysqlUserRepository(db)

	uu := usecase.NewUserUsecase(ur)
	httpdelivery.NewUserHandler(r, uu)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
