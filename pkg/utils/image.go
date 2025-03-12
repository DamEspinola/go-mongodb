package utils

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

// ValidateImageFormat checks if the provided file has a valid image format.
func ValidateImageFormat(fileName string) bool {
	ext := strings.ToLower(strings.TrimPrefix(fileName, "."))
	return ext == "jpg" || ext == "jpeg" || ext == "png"
}

// ResizeImage resizes the given image to the specified width and height.
func ResizeImage(src image.Image, width, height int) (image.Image, error) {
	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			dst.Set(x, y, src.At(x*src.Bounds().Dx()/width, y*src.Bounds().Dy()/height))
		}
	}
	return dst, nil
}

// SaveImage saves the image to the specified file path.
func SaveImage(img image.Image, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	if strings.HasSuffix(filePath, ".png") {
		return png.Encode(file, img)
	} else if strings.HasSuffix(filePath, ".jpg") || strings.HasSuffix(filePath, ".jpeg") {
		return jpeg.Encode(file, img, nil)
	}
	return nil
}