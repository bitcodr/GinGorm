package note

import (
	"net/http"

	"github.com/amiralii/gin-gorm/config"
	"github.com/gin-gonic/gin"
)

func getAll(e *gin.Context) {
	var notes []Note
	db := config.GetDB()

	db.Find(&notes)
	e.JSON(http.StatusOK, notes)
}

func getOne(e *gin.Context) {
	id := e.Param("id")
	var note Note
	db := config.GetDB()

	db.Where("id = ?",id).First(&note)
	e.JSON(http.StatusOK, note)
}

func update(e *gin.Context) {
	id := e.Param("id")
	var note Note
	db := config.GetDB()

	if err := db.Where("id = ?", id).First(&note).Error; err != nil {
		e.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}
	e.Bind(&note)
	db.Save(&note)
	e.JSON(http.StatusOK, &note)
}

func insert(e *gin.Context) {
	var note Note
	db := config.GetDB()
	Migration(db)

	if err := e.BindJSON(&note); err != nil {
		e.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}
	db.Create(&note)
	e.JSON(http.StatusOK, &note)
}

func delete(e *gin.Context) {
	id := e.Param("id")
	var note Note
	db := config.GetDB()

	if err := db.Where("id = ?", id).First(&note).Error; err != nil {
		e.AbortWithStatusJSON(http.StatusNoContent, gin.H{
			"error": err.Error(),
		})
		return
	}
	db.Delete(&note)
	e.JSON(http.StatusOK, gin.H{
		"response": "OK",
	})
}
