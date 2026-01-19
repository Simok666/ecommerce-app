package services

import (
	"errors"
	"image"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
	"github.com/google/uuid"
)

var allowedExt = map[string]bool{
	".jpg": true, ".jpeg": true, ".png": true, ".webp": true,
}

func SaveProductImage(
	file *multipart.FileHeader,
) (string, string, error) {

	ext := filepath.Ext(file.Filename)
	if !allowedExt[ext] {
		return "", "", errors.New("invalid image format")
	}

	src, err := file.Open()
	if err != nil {
		return "", "", err
	}
	defer src.Close()

	img, _, err := image.Decode(src)
	if err != nil {
		return "", "", errors.New("invalid image")
	}

	basePath := "uploads/products"
	os.MkdirAll(basePath, os.ModePerm)

	filename := uuid.New().String()
	mainPath := filepath.Join(basePath, filename+ext)
	thumbPath := filepath.Join(basePath, "thumb_"+filename+ext)

	mainImg := imaging.Fit(img, 1200, 1200, imaging.Lanczos)
	thumbImg := imaging.Thumbnail(img, 300, 300, imaging.Lanczos)

	imaging.Save(mainImg, mainPath)
	imaging.Save(thumbImg, thumbPath)

	return mainPath, thumbPath, nil
}
