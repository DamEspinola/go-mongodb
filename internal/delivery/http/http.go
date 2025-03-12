package http

import (
	"go-gin-image-store/internal/repository/mongodb"
	"go-gin-image-store/internal/storage"
	"go-gin-image-store/internal/usercase"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	db := storage.ConnectDB()

	imageRepo := mongodb.NewMongoImageRepository(db)
	imageService := usercase.NewImageService(imageRepo)
	imageHandler := NewImageHandler(imageService)

	// All routes will be added here
	router.POST("/upload", imageHandler.UploadImage())
	router.GET("/image/:id", imageHandler.ServeImage())
}
