package note

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

//Note struct
type Note struct {
	ID       uuid.UUID `json:"id"`
	Title    string    `json:"title" binding:"required"`
	Text     string    `json:"text" binding:"required"`
	UserID   int       `json:"user_id" binding:"required"`
	Status   string    `json:"status" binding:"required" `
	UpdateAt time.Time `json:"updated_at"`
	CreateAt time.Time `json:"created_at"`
}

//BeforeCreated func
func (note *Note) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())
	scope.SetColumn("ID", uuid.NewV4().String())
	return nil
}

//BeforeUpdated func
func (note *Note) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}

//Migration func
func Migration(db *gorm.DB) {
	if !db.HasTable(&Note{}) {
		err := db.CreateTable(&Note{})
		if err != nil {
			log.Fatal("table exist")
		}
	}
	db.AutoMigrate(&Note{})
}
