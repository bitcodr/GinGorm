package note

import "github.com/gin-gonic/gin"

//Routes func
func Routes(e *gin.Engine) {

	route := e.Group("/v1")
	{
		route.GET("/notes", getAll)
		route.GET("/notes/:id", getOne)
		route.PUT("/notes/:id", update)
		route.POST("/notes", insert)
		route.DELETE("/notes/:id", delete)
	}

}
