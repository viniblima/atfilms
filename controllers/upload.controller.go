package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/atfilms/handlers"
	"github.com/viniblima/atfilms/repository"
)

type UploadController interface {
	UploadItem(c *fiber.Ctx) error
}

type uploadController struct {
	uploadRepo repository.UploadRepository
}

func (controller uploadController) UploadItem(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(handlers.NewJError(err))
	}
	for formFieldName, fileHeaders := range form.File {
		for _, fileHeader := range fileHeaders {
			// process uploaded file here

			sess := session.Must(session.NewSession())

			// Create an uploader with the session and default options
			uploader := s3manager.NewUploader(sess)
			print(fileHeader)
			f, err := os.Open(formFieldName)
			if err != nil {
				return fmt.Errorf("failed to open file %q, %v", formFieldName, err)
			}

			// Upload the file to S3.
			result, err := uploader.Upload(&s3manager.UploadInput{
				Bucket: aws.String(os.Getenv("AWS_BUCKET")),
				Key:    aws.String(os.Getenv("AWS_ACCESS_KEY_ID")),
				Body:   f,
			})
			if err != nil {
				return fmt.Errorf("failed to upload file, %v", err)
			}
			fmt.Printf("file uploaded to, %s\n", aws.StringValue(&result.Location))
		}
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"Msg": "OK",
	})
}

func NewUploadController() UploadController {
	return &uploadController{
		uploadRepo: repository.NewUploadRepository(),
	}
}
