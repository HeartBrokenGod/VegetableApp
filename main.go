package main

import (
	"model"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//	constants
const dsn string = "root:secret@tcp(localhost:3306)/vegetableapp?charset=utf8&parseTime=True&loc=Local"

//	set Routes
func setRoutes(router *gin.Engine) {
	router.POST("/signup", signup)
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
		c.JSON(500, map[string]string{"messege": "DB Connection Failed"})
		//panic(err.Error())
	}
	dbSQL, _ := db.DB()
	defer dbSQL.Close()
	//	get db Connection	***

	//	Perform the Operation

	var user model.Users
	err = c.BindJSON(&user)

	if err != nil {
		c.JSON(401, "Request error: Invalid User")
	}

	if dbc := db.Create(&user); dbc.Error != nil {
		c.JSON(401, dbc.Error)
	} else {
		c.JSON(200, map[string]string{"status": "success", "messege": "Sign Up Successful"})
	}

}

func main() {
	router := gin.Default()

	setRoutes(router)

	router.Run(":8080")
}
