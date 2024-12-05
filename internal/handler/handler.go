package handler

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	_ "music-service/docs"
	"music-service/internal/service"
)

type Handler struct {
	service service.SongService
}

func NewHandler(service service.SongService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	songs := router.Group("/api")
	{
		songs.GET("/songs", h.GetSongs)
		songs.GET("/songs/:id/text", h.GetSongText)
		songs.POST("/songs", h.CreateSong)
		songs.DELETE("songs/:id", h.DeleteSong)
		songs.PUT("songs/:id", h.UpdateSong)
	}
	return router
}
