package utils

import (
	"encoding/base64"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

func EncodeBase64(file *multipart.FileHeader) (string, error) {
	fileData, err := file.Open()
	if err != nil {
		return "", err
	}

	defer fileData.Close()

	bytes, err := io.ReadAll(fileData)
	if err != nil {
		return "", err
	}

	var base64Encoding string
	mimeType := http.DetectContentType(bytes)

	switch mimeType {
	case "image/jpeg":
		base64Encoding = "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding = "data:image/png;base64,"
	case "image/gif":
		base64Encoding = "data:image/gif;base64,"
	case "application/pdf":
		base64Encoding = "data:application/pdf;base64,"
	default:
		base64Encoding = "data:image/png;base64,"
	}

	base64, err := ToBase64(bytes)
	if err != nil {
		return "", err
	}

	return base64Encoding + base64, nil
}

func ToBase64(b []byte) (string, error) {
	encodeBytes := base64.StdEncoding.EncodeToString(b)
	if encodeBytes == "" {
		return "", errors.New("encodeBytes is empty")
	}

	return encodeBytes, nil
}

func DecodeBase64(base64String string) ([]byte, error) {
	parts := strings.SplitN(base64String, ",", 2)
	if len(parts) != 2 {
		return nil, errors.New("invalid base64 string")
	}

	base64Data := parts[1]

	decodeBytes, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return nil, err
	}

	return decodeBytes, nil
}

func SaveImage(base64 string, path string, filename string) error {
	data, err := DecodeBase64(base64)
	if err != nil {
		return err
	}

	err = os.MkdirAll(path, 0666)
	if err != nil {
		return err
	}

	err = os.WriteFile(path + "/" + filename, data, 0666)
	if err != nil {
		return err
	}

	return nil
}

func GetImage(path string, filename string) (string, error) {
	file, err := os.Open(path + "/" + filename)
	if err != nil {
		return "", err
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	base64, err := ToBase64(bytes)
	if err != nil {
		return "", err
	}

	return base64, nil
}

func GenerateFilename(path string, filename string) string {
	return path + "/" + filename
}

func GetExtension(file *multipart.FileHeader) string {
	return file.Filename[strings.LastIndex(file.Filename, ".")+1:]
}