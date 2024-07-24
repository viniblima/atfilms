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
	UploadJobPhoto(c *fiber.Ctx) error
	UploadJobVideo(c *fiber.Ctx) error
	UploadJobComponentFillPhotoHorizontal(c *fiber.Ctx) error
	UploadJobComponentPhotoSlider(c *fiber.Ctx) error
}

type uploadController struct {
	uploadRepo       repository.UploadRepository
	customerRepo     repository.CustomerRepository
	jobRepo          repository.JobRepository
	jobComponentRepo repository.JobComponentRepository
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

func (controller uploadController) UploadJobPhoto(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(handlers.NewJError(err))
	}
	jID := c.Params("id")

	job, errJob := controller.jobRepo.GetJobByID(jID)

	if errJob != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Job not found"})
	}

	newName := time.Now().Format("20060102150405")
	split := strings.Split(file.Header.Get("Content-Type"), "image/")
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
		ReferenceID: job.ID,
		FileName:    fileName,
	}

	photo, errPhoto := controller.uploadRepo.CreatePhoto(&newPhoto)

	if errPhoto != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errPhoto})
	}

	job.Placeholder = photo

	updateJob, errUpdate := controller.jobRepo.UpdateJob(job)

	if errUpdate != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errUpdate})
	}

	return c.Status(http.StatusCreated).JSON(updateJob)
}

func (controller uploadController) UploadJobVideo(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(handlers.NewJError(err))
	}
	jID := c.Params("id")

	job, errJob := controller.jobRepo.GetJobByID(jID)

	if errJob != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Job not found"})
	}

	newName := time.Now().Format("20060102150405")
	split := strings.Split(file.Header.Get("Content-Type"), "video/")
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

	newVideo := models.Video{
		ReferenceID: job.ID,
		FileName:    fileName,
	}

	video, errVideo := controller.uploadRepo.CreateVideo(&newVideo)

	if errVideo != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errVideo})
	}

	job.Video = video

	updateJob, errUpdate := controller.jobRepo.UpdateJob(job)

	if errUpdate != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errUpdate})
	}

	return c.Status(http.StatusCreated).JSON(updateJob)
}

func (controller uploadController) UploadJobComponentFillPhotoHorizontal(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(handlers.NewJError(err))
	}
	cID := c.Params("id")

	jc, errJc := controller.jobComponentRepo.GetJobComponentByID(cID)

	if errJc != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Job Component not found"})
	}

	folder, fileName := handlers.GenerateFileName(file)
	errSaveTmp := c.SaveFile(file, folder)

	if errSaveTmp != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errSaveTmp})
	}

	_, errOpen := os.Open(folder)

	if errOpen != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errOpen})
	}

	newPhoto := models.Photo{
		ReferenceID: jc.ID,
		FileName:    fileName,
	}

	photo, errPhoto := controller.uploadRepo.CreatePhoto(&newPhoto)

	if errPhoto != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errPhoto})
	}

	jc.FillPhotoHorizontal = *photo

	updateJc, errUpdate := controller.jobComponentRepo.UpdateJobComponent(jc)

	if errUpdate != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errUpdate})
	}

	return c.Status(http.StatusCreated).JSON(updateJc)
}

func (controller uploadController) UploadJobComponentPhotoSlider(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(handlers.NewJError(err))
	}
	cID := c.Params("id")

	jc, errJc := controller.jobComponentRepo.GetJobComponentByID(cID)

	if errJc != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Job Component not found"})
	}

	folder, fileName := handlers.GenerateFileName(file)
	errSaveTmp := c.SaveFile(file, folder)

	if errSaveTmp != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errSaveTmp})
	}

	_, errOpen := os.Open(folder)

	if errOpen != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errOpen})
	}

	newPhoto := models.Photo{
		ReferenceID: jc.ID,
		FileName:    fileName,
	}

	photo, errPhoto := controller.uploadRepo.CreatePhoto(&newPhoto)

	if errPhoto != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errPhoto})
	}

	append, errAppend := controller.jobComponentRepo.AppendPhotoToSlider(jc, photo)

	if errAppend != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errAppend})
	}

	return c.Status(http.StatusCreated).JSON(append)
}

func (controller uploadController) UploadJobComponentVideo(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(handlers.NewJError(err))
	}
	cID := c.Params("id")

	jc, errJc := controller.jobComponentRepo.GetJobComponentByID(cID)

	if errJc != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Job Component not found"})
	}

	folder, fileName := handlers.GenerateFileName(file)
	errSaveTmp := c.SaveFile(file, folder)

	if errSaveTmp != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errSaveTmp})
	}

	_, errOpen := os.Open(folder)

	if errOpen != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errOpen})
	}

	newVideo := models.Video{
		ReferenceID: jc.ID,
		FileName:    fileName,
	}

	video, errVideo := controller.uploadRepo.CreateVideo(&newVideo)

	if errVideo != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errVideo})
	}

	append, errAppend := controller.jobComponentRepo.AppendVideo(jc, video)

	if errAppend != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errAppend})
	}

	return c.Status(http.StatusCreated).JSON(append)
}

func NewUploadController() UploadController {
	return &uploadController{
		uploadRepo:       repository.NewUploadRepository(),
		customerRepo:     repository.NewCustomerRepository(),
		jobRepo:          repository.NewJobRepository(),
		jobComponentRepo: repository.NewJobComponentRepository(),
	}
}
