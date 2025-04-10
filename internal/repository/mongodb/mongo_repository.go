package mongodb

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"mime"
	"net/http"
	"path/filepath"
	"time"

	"go-gin-image-store/internal/domain/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoImageRepository struct {
	client     *mongo.Client
	bucketName string
}

func compressImage(data []byte, contentType string) ([]byte, error) {
	// Decodificar la imagen original
	img, format, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer

	switch {
	case contentType == "image/jpeg" || format == "jpeg" || format == "jpg":
		// Comprimir JPEG con 75% de calidad (puedes ajustar este valor)
		opts := jpeg.Options{Quality: 75}
		if err := jpeg.Encode(&buf, img, &opts); err != nil {
			return nil, err
		}

	case contentType == "image/png" || format == "png":
		encoder := png.Encoder{
			CompressionLevel: png.BestCompression,
		}
		if err := encoder.Encode(&buf, img); err != nil {
			return nil, err
		}

	default:
		return data, nil
	}

	// Verificar si la compresión tuvo efecto
	if buf.Len() >= len(data) {
		return data, nil
	}

	return buf.Bytes(), nil
}

// NewMongoImageRepository crea un nuevo repositorio para imágenes con MongoDB
func NewMongoImageRepository(client *mongo.Client) *MongoImageRepository {
	return &MongoImageRepository{
		client:     client,
		bucketName: "photos",
	}
}

// Store guarda una imagen en GridFS y retorna metadatos
func (r *MongoImageRepository) Store(ctx context.Context, filename string, content io.Reader) (*models.Image, error) {
	// Leer todo el contenido para determinar tipo y tamaño
	data, err := io.ReadAll(content)
	if err != nil {
		return nil, err
	}

	// Determinar el tipo de contenido
	ext := filepath.Ext(filename)
	contentType := mime.TypeByExtension(ext)
	if contentType == "" {
		contentType = http.DetectContentType(data)
	}

	// comprimir la imagen
	compressedData, err := compressImage(data, contentType)
	if err != nil {
		return nil, fmt.Errorf("compress image: %w", err)
	}

	// Crear un bucket GridFS
	bucket, err := gridfs.NewBucket(
		r.client.Database("image-server"),
		options.GridFSBucket().SetName(r.bucketName),
	)
	if err != nil {
		return nil, err
	}

	// Metadatos para el archivo
	uploadOpts := options.GridFSUpload().
		SetMetadata(map[string]interface{}{
			"contentType": contentType,
			"uploadDate":  time.Now(),
			"filename":    filename,
			"compressed":  true,
		})

	// Subir archivo a GridFS
	uploadStream, err := bucket.OpenUploadStream(filename, uploadOpts)
	if err != nil {
		return nil, err
	}
	defer uploadStream.Close()

	// Escribir datos al stream
	size, err := uploadStream.Write(compressedData)
	if err != nil {
		return nil, err
	}

	// Crear y retornar el modelo
	image := &models.Image{
		ID:          uploadStream.FileID.(primitive.ObjectID).Hex(),
		Name:        filename,
		ContentType: contentType,
		Size:        int64(size),
		CreatedAt:   time.Now(),
	}

	return image, nil
}

// FindByID encuentra una imagen por su ID y retorna sus metadatos y contenido
func (r *MongoImageRepository) FindByID(ctx context.Context, id string) (*models.Image, io.Reader, error) {
	// Convertir string ID a ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, nil, err
	}

	// Crear bucket
	bucket, err := gridfs.NewBucket(
		r.client.Database("image-server"),
		options.GridFSBucket().SetName(r.bucketName),
	)
	if err != nil {
		return nil, nil, err
	}

	// Buscar metadatos del archivo
	var fileInfo gridfs.File
	err = bucket.GetFilesCollection().FindOne(ctx,
		primitive.M{"_id": objectID}).Decode(&fileInfo)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil, errors.New("image not found")
		}
		return nil, nil, err
	}

	// Extraer metadatos
	var fileDoc struct {
		Metadata primitive.M `bson:"metadata"`
	}
	err = bucket.GetFilesCollection().FindOne(ctx,
		primitive.M{"_id": objectID}).Decode(&fileDoc)
	if err != nil {
		return nil, nil, err
	}

	metadata := fileDoc.Metadata
	if metadata == nil {
		metadata = primitive.M{}
	}

	// Crear buffer para descargar el archivo
	var buf bytes.Buffer
	_, err = bucket.DownloadToStream(objectID, &buf)
	if err != nil {
		return nil, nil, err
	}

	// Crear modelo de imagen
	image := &models.Image{
		ID:          id,
		Name:        fileInfo.Name,
		ContentType: metadata["contentType"].(string),
		Size:        fileInfo.Length,
		CreatedAt:   fileInfo.UploadDate,
	}

	return image, &buf, nil
}
