package http

import (
	"go-gin-image-store/internal/usercase"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ImageHandler struct {
	imageService *usercase.ImageService
}

func NewImageHandler(imageService *usercase.ImageService) *ImageHandler {
	return &ImageHandler{imageService: imageService}
}

func (h *ImageHandler) UploadImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, header, err := c.Request.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		defer file.Close()

		image, err := h.imageService.UploadImage(c.Request.Context(), header.Filename, file)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"fileId":   image.ID,
			"filename": header.Filename,
			"size":     image.Size,
		})
	}
}
func (h *ImageHandler) ServeImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		image, data, err := h.imageService.GetImage(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		dataBytes, err := io.ReadAll(data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read image data"})
			return
		}

		c.Data(http.StatusOK, image.ContentType, dataBytes)
	}
}
