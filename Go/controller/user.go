package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/swaggo/swag/example/GO/model"
	"github.com/swaggo/swag/example/GO/mysql"
	"golang.org/x/crypto/bcrypt"
)

func (c *Controller) Login(ctx *gin.Context){
	db := mysql.Connect()

	email := ctx.Params.ByName("email")
	password := ctx.Params.ByName("password")

	query := "SELECT id, name, email, password from users where email = ?"
	result := db.QueryRow(query, email)
	
	user := &model.User{}

	err := result.Scan(&user.Id, &user.Name, &user.Email, &user.Password)

	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	res := make(map[string]string)

	res["id"] = strconv.Itoa(user.Id)
	res["name"] = user.Name
	res["email"] = user.Email
	res["password"] = user.Password

	response := CheckPasswordHash(password, res["password"])

	if !response {
		ctx.JSON(http.StatusBadRequest, "Email or Password is Incorrect")
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *Controller) RegisterUser(ctx *gin.Context) {
	db := mysql.Connect()

	var user model.User
	err := ctx.ShouldBindJSON(&user);
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	account := model.User{
		Name: user.Name,
		Password: user.Password,
		Email: user.Email,
	}

	hash, err := HashPassword(string(account.Password))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return 
	}
	
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	result, err := db.Exec(query, account.Name, account.Email, hash)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return 
	}
	
	id, err := result.LastInsertId()

	ctx.JSON(http.StatusOK, id)
}

func (c *Controller) ListUsers(ctx *gin.Context) {

	db := mysql.Connect()

	rows, err := db.Query(`select id, name, email from users`)

	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	users := make([]*model.User, 0)

	for rows.Next() {
		user := new(model.User)
		rows.Scan(&user.Id, &user.Name, &user.Email)
		users = append(users, user)
	}
	
	ctx.JSON(http.StatusOK, users)
}

func HashPassword(password string)(string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 9)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}