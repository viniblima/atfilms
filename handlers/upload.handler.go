package handlers

import (
	"fmt"
	"mime/multipart"
	"strings"
	"time"
)

func GenerateFileName(file *multipart.FileHeader) (string, string) {
	newName := time.Now().Format("20060102150405")
	split := strings.Split(file.Header.Get("Content-Type"), "image/")
	fileName := fmt.Sprintf("%s.%s", newName, split[1])
	folder := fmt.Sprintf("uploads/%s", fileName)

	return folder, fileName
}
