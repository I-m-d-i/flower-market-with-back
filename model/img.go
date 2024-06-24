package model

import (
	"context"
	database "flover-market/db"
)

type Img struct {
	ImgId     int
	FileName  string
	ImageData []byte
}

func GetImg(ImgName string) (Img, error) {
	db := database.GetDB()
	var img Img
	err := db.QueryRow(context.Background(), "SELECT _id, filename, imagedata FROM files WHERE filename = $1", ImgName).Scan(&img.ImgId, &img.FileName, &img.ImageData)
	if err != nil {
		return img, err
	}
	return img, nil
}

func AddImage(img Img) error {
	db := database.GetDB()
	_, err := db.Exec(context.Background(), "INSERT INTO Files (FileName, ImageData) VALUES ($1, $2)", img.FileName, img.ImageData)
	if err != nil {
		return err
	}
	return nil
}
