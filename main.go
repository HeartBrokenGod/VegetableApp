package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//	constants
const dsn string = "root:secret@tcp(localhost:3306)/vegetableapp?charset=utf8&parseTime=True&loc=Local"

//	set Routes
func setRoutes(router *gin.Engine) {
	router.POST("/signup", signup)
	router.POST("/login", login)
}

// GenerateToken returns a unique token based on the provided email string
func GenerateToken(email string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(email), bcrypt.DefaultCost)
	if err != nil {
		//log.Fatal(err)
	}
	//fmt.Println("Hash to store:", string(hash))

	hasher := md5.New()
	hasher.Write(hash)
	return hex.EncodeToString(hasher.Sum(nil))
}

//	get db Connection
func getNewDbConnection() (*gorm.DB, error) {

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil

}

//	user sign up
func signup(c *gin.Context) {
	//	get db Connection
	db, err := getNewDbConnection()
	if err != nil {
		c.JSON(500, map[string]string{"message": "DB Connection Failed"})
		//panic(err.Error())
	}
	dbSQL, _ := db.DB()
	defer dbSQL.Close()
	//	get db Connection	***

	//	Perform the Operation

	var user model.Users
	err = c.BindJSON(&user)

	if err != nil {
		c.JSON(402, map[string]string{"message": "Request error: Invalid User"})
	}

	if dbc := db.Create(&user); dbc.Error != nil {
		c.JSON(402, dbc.Error)
	} else {
		c.JSON(200, map[string]string{"status": "success", "message": "Sign Up Successful"})
	}

}

//	user sign up
func login(c *gin.Context) {
	//	get db Connection
	db, err := getNewDbConnection()
	if err != nil {
		c.JSON(500, map[string]string{"message": "DB Connection Failed"})
		//panic(err.Error())
	}
	dbSQL, _ := db.DB()
	defer dbSQL.Close()
	//	get db Connection	***

	//	Perform the Operation

	type loginCredentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var loginDetails loginCredentials
	err = c.BindJSON(&loginDetails)

	if err != nil {
		fmt.Println(err)
		c.JSON(402, map[string]string{"message": "Request error: Invalid Request Body"})
	}

	usermgr := model.UsersMgr(db)

	user, err := usermgr.GetFromEmail(loginDetails.Email)

	if err != nil {
		c.JSON(402, map[string]string{"message": "Request error: Invalid Email"})
	} else {
		if loginDetails.Password == user.Password {
			//	Generate Token
			token := GenerateToken(loginDetails.Email)

			fmt.Println("token Generated")

			var tokenObj model.Tokens

			tokenObj = model.Tokens{Token: token, UserID: user.ID, StatusID: 1}

			fmt.Println(tokenObj.UserID, " , ", tokenObj.StatusID, " , ", tokenObj.Token)

			if dbc := db.Create(tokenObj); dbc.Error != nil {
				c.JSON(402, dbc.Error)
			} else {
				c.JSON(200, map[string]string{"status": "success", "message": "Login Successful", "token": token})
			}

		} else {
			c.JSON(402, map[string]string{"message": "Request error: Invalid Password"})
		}
	}

	// if dbc := db.Create(&user); dbc.Error != nil {
	// 	c.JSON(402, dbc.Error)
	// } else {
	// 	c.JSON(200, map[string]string{"status": "success", "message": "Sign Up Successful"})
	// }

}

func main() {
	router := gin.Default()

	setRoutes(router)

	router.Run(":8080")
}
