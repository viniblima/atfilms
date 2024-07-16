package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/atfilms/handlers"
	"github.com/viniblima/atfilms/models"
	"github.com/viniblima/atfilms/repository"
)

type UploadController interface {
	UploadCustomerLogo(c *fiber.Ctx) error
}

type uploadController struct {
	uploadRepo   repository.UploadRepository
	customerRepo repository.CustomerRepository
	jobRepo      repository.JobRepository
}

func (controller uploadController) UploadCustomerLogo(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(handlers.NewJError(err))
	}
	cID := c.Params("id")

	customer, errCustomer := controller.customerRepo.GetCustomerByID(cID)

	if errCustomer != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Customer not found"})
	}

	newName := time.Now().Format("20060102150405")
	split := strings.Split(file.Header.Get("Content-Type"), "image/")
	fmt.Println(split)

	fileName := fmt.Sprintf("%s.%s", newName, split[1])

	folder := fmt.Sprintf("uploads/%s", fileName)
	errSaveTmp := c.SaveFile(file, folder)

	if errSaveTmp != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errSaveTmp})
	}

	_, errOpen := os.Open(folder)

	if errOpen != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errOpen})
	}
	newPhoto := models.Photo{
		ReferenceID: customer.ID,
		FileName:    fileName,
	}

	photo, errPhoto := controller.uploadRepo.CreatePhoto(&newPhoto)

	if errPhoto != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errPhoto})
	}
	customer.Logo = *photo

	updateCustomer, errUpdate := controller.customerRepo.UpdateCustomer(customer)

	if errUpdate != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errUpdate})
	}

	// return c.Status(http.StatusOK).JSON(fiber.Map{"result": f})

	// sess := session.Must(session.NewSession())

	// uploader := s3manager.NewUploader(sess)

	// result, errUploader := uploader.Upload(&s3manager.UploadInput{
	// 	Bucket: aws.String(os.Getenv("AWS_BUCKET")),
	// 	Key:    aws.String(os.Getenv("AWS_ACCESS_KEY_ID")),
	// 	Body:   f,
	// })

	// if errUploader != nil {
	// 	return c.Status(http.StatusBadRequest).JSON(errUploader)
	// }

	return c.Status(http.StatusCreated).JSON(updateCustomer)
}

func NewUploadController() UploadController {
	return &uploadController{
		uploadRepo:   repository.NewUploadRepository(),
		customerRepo: repository.NewCustomerRepository(),
	}
}
