package storage

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

const (
	UPLOAD_PATH string = "static/images/"
)

// UploadFile uploads an object
func UploadFile(file multipart.File, objectName string) (fileLocation string, err error) {
	// Destination
	dst, err := os.Create(filepath.Join(UPLOAD_PATH, filepath.Base(objectName)))
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, file); err != nil {
		return "", err
	}

	return UPLOAD_PATH + objectName, nil
}

func DeleteFile(objectName string) error {
	err := os.Remove(objectName)
	if err != nil {
		return err
	}
	
	return nil
}