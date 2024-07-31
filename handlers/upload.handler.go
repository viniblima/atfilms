package handlers

import (
	"fmt"
	"mime/multipart"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gofiber/fiber/v2"
)

func generateFileName(file *multipart.FileHeader, fileType string) (string, string) {
	newName := time.Now().Format("20060102150405")
	split := strings.Split(file.Header.Get("Content-Type"), fmt.Sprintf("%s/", fileType))
	fileName := fmt.Sprintf("%s.%s", newName, split[1])
	folder := fmt.Sprintf("uploads/%s", fileName)
	return folder, fileName
}

func UploadS3(c *fiber.Ctx, file *multipart.FileHeader, fileType string) (*s3manager.UploadOutput, error) {

	folder, fileName := generateFileName(file, fileType)

	errSaveTmp := c.SaveFile(file, folder)

	if errSaveTmp != nil {
		return nil, errSaveTmp
	}

	f, errOpen := os.Open(folder)
	fmt.Println(folder)
	if errOpen != nil {
		return nil, errOpen
	}
	fmt.Println(f)
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	}))
	fmt.Println(sess)
	uploader := s3manager.NewUploader(sess)
	fmt.Println(os.Getenv("AWS_BUCKET"))
	fmt.Println(os.Getenv("AWS_REGION"))
	result, errUploader := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET")),
		Key:    aws.String(fileName),
		Body:   f,
	})

	fmt.Println(errUploader)
	if errUploader != nil {
		return nil, errUploader
	}
	r := os.Remove(folder)

	fmt.Println(r)
	if r != nil {
		return nil, r
	}

	fmt.Println(result)
	return result, nil
}

func RemoveS3(key string) (*s3.DeleteObjectOutput, error) {

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	}))
	svc := s3.New(sess)

	input := &s3.DeleteObjectInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET")),
		Key:    aws.String(key),
	}
	result, err := svc.DeleteObject(input)

	if err != nil {
		return nil, err
	}

	return result, nil
}
