package controllers

import (
	// "fmt"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gofiber/fiber/v2"

	// uuid "github.com/satori/go.uuid"
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
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(handlers.NewJError(err))
	}

	newName := time.Now().Format("20060102150405")
	split := strings.Split(file.Header.Get("Content-Type"), "image/")
	fmt.Println(split)

	fileName := fmt.Sprintf("%s.%s", newName, split[1])

	folder := fmt.Sprintf("tmp/uploads/%s", fileName)
	errSaveTmp := c.SaveFile(file, folder)

	if errSaveTmp != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errSaveTmp})
	}

	f, errOpen := os.Open(folder)

	if errOpen != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	// return c.Status(http.StatusOK).JSON(fiber.Map{"result": f})

	sess := session.Must(session.NewSession())

	uploader := s3manager.NewUploader(sess)

	result, errUploader := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET")),
		Key:    aws.String(os.Getenv("AWS_ACCESS_KEY_ID")),
		Body:   f,
	})

	if errUploader != nil {
		return c.Status(http.StatusBadRequest).JSON(errUploader)
	}

	return c.Status(http.StatusOK).JSON(result)
}

func NewUploadController() UploadController {
	return &uploadController{
		uploadRepo: repository.NewUploadRepository(),
	}
}
