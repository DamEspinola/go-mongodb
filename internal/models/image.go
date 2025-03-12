package models

import (
	"io"
	"mime/multipart"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Image struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Filename string             `bson:"filename" json:"filename"`
	Data     []byte             `bson:"data" json:"data"`
	Uploaded primitive.DateTime `bson:"uploaded" json:"uploaded"`
}

func NewImage(file *multipart.FileHeader) (*Image, error) {
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	data, err := io.ReadAll(src)
	if err != nil {
		return nil, err
	}

	return &Image{
		ID:       primitive.NewObjectID(),
		Filename: file.Filename,
		Data:     data,
	}, nil
}
