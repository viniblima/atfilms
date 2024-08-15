package controllers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
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
	UploadJobComponentVideo(c *fiber.Ctx) error

	UpdatePhotoPositionInSlider(c *fiber.Ctx) error
	RemovePhotoFromSlider(c *fiber.Ctx) error

	UpdateVideoPositionInSlider(c *fiber.Ctx) error
	RemoveVideoFromList(c *fiber.Ctx) error

	UploadAwardImage(c *fiber.Ctx) error
}

type uploadController struct {
	uploadRepo       repository.UploadRepository
	customerRepo     repository.CustomerRepository
	jobRepo          repository.JobRepository
	jobComponentRepo repository.JobComponentRepository
	awardRepo        repository.AwardRepository
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

func (controller uploadController) UploadAwardImage(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(handlers.NewJError(err))
	}
	cID := c.Params("id")

	award, errAward := controller.awardRepo.GetByID(cID)

	if errAward != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Customer not found"})
	}
	if award.AwardImage != nil {
		_, errRemove := handlers.RemoveS3(award.AwardImage.FileName)

		if errRemove != nil {
			return c.Status(http.StatusBadRequest).JSON(errRemove)
		}
	}

	upload, errUpload := handlers.UploadS3(c, file, "image")

	if errUpload != nil {
		return c.Status(http.StatusBadRequest).JSON(errUpload)
	}

	newPhoto := models.Photo{
		AwardImageID: &award.ID,
		FileName:     upload.Location,
	}

	photo, errPhoto := controller.uploadRepo.CreatePhoto(&newPhoto)

	if errPhoto != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errPhoto})
	}

	award.AwardImage = photo

	updateAward, errUpdate := controller.awardRepo.Update(award)

	if errUpdate != nil {
		return c.Status(http.StatusBadRequest).JSON(errUpdate)
	}

	return c.Status(http.StatusCreated).JSON(updateAward)
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

	if jc.FillPhotoHorizontal != nil {
		_, errUpload := handlers.RemoveS3(jc.FillPhotoHorizontal.FileName)

		if errUpload != nil {
			return c.Status(http.StatusBadRequest).JSON(errUpload)
		}

		errRemove := controller.uploadRepo.RemovePhotoByID(jc.FillPhotoHorizontal.ID)

		if errRemove != nil {
			return c.Status(http.StatusBadRequest).JSON(errRemove)
		}
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

	newPhoto := &models.Photo{
		SliderID: &jc.ID,
		FileName: upload.Location,
	}

	_, errPhoto := controller.uploadRepo.CreatePhoto(newPhoto)

	if errPhoto != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errPhoto})
	}

	list, errList := controller.uploadRepo.GetPhotosBySliderID(jc.ID)

	if errList != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errList})
	}

	jc.Slider = list

	append, errAppend := controller.jobComponentRepo.UpdateJobComponent(jc)

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

	upload, errUpload := handlers.UploadS3(c, file, "video")

	if errUpload != nil {
		return c.Status(http.StatusBadRequest).JSON(errUpload)
	}

	newVideo := &models.Video{
		JobComponentVideosID: &jc.ID,
		FileName:             upload.Location,
	}

	_, errVideo := controller.uploadRepo.CreateVideo(newVideo)

	if errVideo != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errVideo})
	}

	// append, errAppend := controller.jobComponentRepo.AppendVideo(jc, video)

	list, errList := controller.uploadRepo.GetVideosByComponentID(jc.ID)

	if errList != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errList})
	}

	jc.Videos = list

	append, errAppend := controller.jobComponentRepo.UpdateJobComponent(jc)

	if errAppend != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errAppend})
	}

	return c.Status(http.StatusCreated).JSON(append)
}

func (controller uploadController) UpdatePhotoPositionInSlider(c *fiber.Ctx) error {
	body := new(models.UpdateFilePositionStruct)

	c.BodyParser(&body)

	err := validator.New().Struct(body)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(handlers.NewJError(err))
	}

	cID := c.Params("id")

	photo, errPhoto := controller.uploadRepo.GetPhotoByID(cID)

	if errPhoto != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Photo not found"})
	}

	photo.Position = &body.Position

	update, errUpdate := controller.uploadRepo.UpdatePhoto(photo)

	if errUpdate != nil {
		return c.Status(http.StatusBadRequest).JSON(errUpdate)
	}

	return c.Status(http.StatusOK).JSON(update)
}

func (controller uploadController) RemovePhotoFromSlider(c *fiber.Ctx) error {
	body := new(models.UpdateFilePositionStruct)

	c.BodyParser(&body)

	err := validator.New().Struct(body)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(handlers.NewJError(err))
	}

	cID := c.Params("id")

	photo, errPhoto := controller.uploadRepo.GetPhotoByID(cID)

	if errPhoto != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Photo not found"})
	}

	errRemove := controller.uploadRepo.RemovePhotoByID(photo.ID)

	if errRemove != nil {
		return c.Status(http.StatusBadRequest).JSON(errRemove)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"Msg": "Photo removed"})
}

func (controller uploadController) UpdateVideoPositionInSlider(c *fiber.Ctx) error {
	body := new(models.UpdateFilePositionStruct)

	c.BodyParser(&body)

	err := validator.New().Struct(body)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(handlers.NewJError(err))
	}

	cID := c.Params("id")

	video, errVideo := controller.uploadRepo.GetVideoByID(cID)

	if errVideo != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Video not found"})
	}

	video.Position = &body.Position

	update, errUpdate := controller.uploadRepo.UpdateVideo(video)

	if errUpdate != nil {
		return c.Status(http.StatusBadRequest).JSON(errUpdate)
	}

	return c.Status(http.StatusOK).JSON(update)
}

func (controller uploadController) RemoveVideoFromList(c *fiber.Ctx) error {
	body := new(models.UpdateFilePositionStruct)

	c.BodyParser(&body)

	err := validator.New().Struct(body)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(handlers.NewJError(err))
	}

	cID := c.Params("id")

	video, errVideo := controller.uploadRepo.GetVideoByID(cID)

	if errVideo != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Video not found"})
	}

	errRemove := controller.uploadRepo.RemoveVideoByID(video.ID)

	if errRemove != nil {
		return c.Status(http.StatusBadRequest).JSON(errRemove)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"Msg": "Video removed"})
}

func NewUploadController() UploadController {
	return &uploadController{
		uploadRepo:       repository.NewUploadRepository(),
		customerRepo:     repository.NewCustomerRepository(),
		jobRepo:          repository.NewJobRepository(),
		jobComponentRepo: repository.NewJobComponentRepository(),
		awardRepo:        repository.NewAwardRepository(),
	}
}
