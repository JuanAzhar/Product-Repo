package migration

import (
	users "product-rest-api/features/user/model"
	products "product-rest-api/features/product/model"

	"gorm.io/gorm"
)

func InitMigrationMysql(db *gorm.DB) {
	db.AutoMigrate(&users.Users{})
	db.AutoMigrate(&products.Products{})
}
