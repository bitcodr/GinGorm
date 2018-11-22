package note

import "github.com/amiralii/gin-gorm/config"


type repository interface{
	All(notes []Note)
}

func All(notes []Note) {
	db := config.GetDB()
	db.Find(&notes)
}