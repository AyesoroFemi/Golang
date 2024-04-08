package initializers

import "desktop/go-projectt/models"



func SyncDatabase(){
	DB.AutoMigrate(&models.User{})
}