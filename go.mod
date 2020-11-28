module VegetableApp

go 1.15

replace model => ./model

require (
	github.com/gin-gonic/gin v1.6.3
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.7
	model v0.0.0-00010101000000-000000000000
)
