package controllers

import (
	"net/http"

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

	upload, errUpload := handlers.UploadS3(c, file, "image")

	if errUpload != nil {
		return c.Status(http.StatusBadRequest).JSON(errUpload)
	}

	newPhoto := models.Photo{
		LogoID:   &customer.ID,
		FileName: upload.Location,
	}

	photo, errPhoto := controller.uploadRepo.CreatePhoto(&newPhoto)

	if errPhoto != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errPhoto})
	}
	customer.Logo = photo

	updateCustomer, errUpdate := controller.customerRepo.UpdateCustomer(customer)

	if errUpdate != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errUpdate})
	}
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

	upload, errUpload := handlers.UploadS3(c, file, "image")

	if errUpload != nil {
		return c.Status(http.StatusBadRequest).JSON(errUpload)
	}
	newPhoto := models.Photo{
		PlaceholderID: &job.ID,
		FileName:      upload.Location,
	}

	photo, errPhoto := controller.uploadRepo.CreatePhoto(&newPhoto)

	if errPhoto != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errPhoto})
	}

	job.Placeholder = *photo

	updateJob, errUpdate := controller.jobRepo.UpdateJob(job)

	if errUpdate != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errUpdate})
	}

	return c.Status(http.StatusCreated).JSON(updateJob)
}

func (controller uploadController) UploadJobVideo(c *fiber.Ctx) error {
	file, err := c.FormFile("video")
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(handlers.NewJError(err))
	}
	jID := c.Params("id")

	job, errJob := controller.jobRepo.GetJobByID(jID)

	if errJob != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Job not found"})
	}

	upload, errUpload := handlers.UploadS3(c, file, "video")

	if errUpload != nil {
		return c.Status(http.StatusBadRequest).JSON(errUpload)
	}
	newVideo := models.Video{
		MainVideoID: &job.ID,
		FileName:    upload.Location,
	}

	video, errVideo := controller.uploadRepo.CreateVideo(&newVideo)

	if errVideo != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errVideo})
	}

	job.MainVideo = *video

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

	upload, errUpload := handlers.UploadS3(c, file, "image")

	if errUpload != nil {
		return c.Status(http.StatusBadRequest).JSON(errUpload)
	}

	newPhoto := models.Photo{
		FillPhotoHorizontalID: &jc.ID,
		FileName:              upload.Location,
	}

	photo, errPhoto := controller.uploadRepo.CreatePhoto(&newPhoto)

	if errPhoto != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errPhoto})
	}

	jc.FillPhotoHorizontal = photo

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

	upload, errUpload := handlers.UploadS3(c, file, "image")

	if errUpload != nil {
		return c.Status(http.StatusBadRequest).JSON(errUpload)
	}

	newPhoto := models.Photo{
		SliderID: &jc.ID,
		FileName: upload.Location,
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
	file, err := c.FormFile("video")
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(handlers.NewJError(err))
	}
	cID := c.Params("id")

	jc, errJc := controller.jobComponentRepo.GetJobComponentByID(cID)

	if errJc != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Job Component not found"})
	}

	upload, errUpload := handlers.UploadS3(c, file, "image")

	if errUpload != nil {
		return c.Status(http.StatusBadRequest).JSON(errUpload)
	}

	newVideo := models.Video{
		JobComponentVideosID: &jc.ID,
		FileName:             upload.Location,
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
