package controllers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/atfilms/handlers"
	"github.com/viniblima/atfilms/models"
	"github.com/viniblima/atfilms/repository"
)

type JobController interface {
	CreateJob(c *fiber.Ctx) error
	UpdateJob(c *fiber.Ctx) error
	ListAllJobs(c *fiber.Ctx) error
	GetJobBySlug(c *fiber.Ctx) error
	RemoveJob(c *fiber.Ctx) error
}

type jobController struct {
	jobRepo          repository.JobRepository
	customerRepo     repository.CustomerRepository
	jobComponentRepo repository.JobComponentRepository
	tagRepo          repository.TagRepository
	awardRepo        repository.AwardRepository
}

type CreateJobStruct struct {
	Name       string                `json:"Name" validate:"required,min=3,max=32"`
	Slug       string                `json:"Slug" validate:"required,min=3,max=32"`
	CustomerID string                `json:"CustomerID" validate:"required"`
	ShowInHome bool                  `json:"ShowInHome" `
	Components []models.JobComponent `json:"Components"`
	Tags       []string
	Awards     []string
}

func (controller jobController) UpdateJob(c *fiber.Ctx) error {
	body := new(CreateJobStruct)

	c.BodyParser(&body)

	err := validator.New().Struct(body)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(handlers.NewJError(err))
	}
	id := c.Params("id")

	jobFound, errJobFound := controller.jobRepo.GetJobByID(id)

	if errJobFound != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Job not found",
		})
	}

	Customer, errCustomer := controller.customerRepo.GetCustomerByID(body.CustomerID)

	if errCustomer != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Customer not found",
		})
	}

	var tags []*models.Tag
	for i := 0; i < len(body.Tags); i++ {
		tag, errTag := controller.tagRepo.GetByName(body.Tags[i])

		if errTag != nil {
			objTag := models.Tag{
				Name: body.Tags[i],
			}

			newTag, errNewTag := controller.tagRepo.Create(&objTag)

			if errNewTag != nil {
				tags = append(tags, newTag)
			}
		} else {
			tags = append(tags, tag)
		}
	}

	var awards []*models.Award
	for i := 0; i < len(body.Awards); i++ {
		a, errA := controller.awardRepo.GetByID(body.Awards[i])

		if errA == nil {
			awards = append(awards, a)
		}
	}

	jobFound.Customer = Customer
	jobFound.Name = body.Name
	jobFound.Slug = body.Slug
	jobFound.ShowInHome = body.ShowInHome

	update, errUpdate := controller.jobRepo.UpdateJob(jobFound)

	_, errClear, errAppend := controller.jobRepo.AppendTag(update, tags)

	if errClear != nil {
		return c.Status(http.StatusBadRequest).JSON(errClear)
	}

	if errAppend != nil {
		return c.Status(http.StatusBadRequest).JSON(errAppend)
	}

	appendAward, errClearAward, errAppendAward := controller.jobRepo.AppendAward(update, awards)

	if errClearAward != nil {
		return c.Status(http.StatusBadRequest).JSON(errClearAward)
	}

	if errAppendAward != nil {
		return c.Status(http.StatusBadRequest).JSON(errAppendAward)
	}

	for i := 0; i < len(body.Components); i++ {
		found, errFound := controller.jobComponentRepo.GetJobComponentByID(body.Components[i].ID)
		if errFound != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Job Component not found"})
		}
		found.Text = body.Components[i].Text
		found.Title = body.Components[i].Title

		_, errCpnt := controller.jobComponentRepo.UpdateJobComponent(found)

		if errCpnt != nil {
			return c.Status(http.StatusBadRequest).JSON(errCpnt)
		}
	}

	if errUpdate != nil {
		return c.Status(http.StatusBadRequest).JSON(errUpdate)
	}

	return c.Status(http.StatusOK).JSON(appendAward)
}

func (controller jobController) CreateJob(c *fiber.Ctx) error {
	body := new(CreateJobStruct)

	c.BodyParser(&body)

	err := validator.New().Struct(body)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(handlers.NewJError(err))
	}

	customer, errCustomer := controller.customerRepo.GetCustomerByID(body.CustomerID)

	if errCustomer != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Customer not found",
		})
	}

	var tags []*models.Tag
	for i := 0; i < len(body.Tags); i++ {
		tag, errTag := controller.tagRepo.GetByName(body.Tags[i])

		if errTag != nil {
			objTag := models.Tag{
				Name: body.Tags[i],
			}

			newTag, errNewTag := controller.tagRepo.Create(&objTag)

			if errNewTag != nil {
				tags = append(tags, newTag)
			}
		} else {
			tags = append(tags, tag)
		}
	}

	var awards []*models.Award

	for i := 0; i < len(body.Awards); i++ {
		a, errA := controller.awardRepo.GetByID(body.Awards[i])

		if errA == nil {
			awards = append(awards, a)
		}
	}

	newJob := models.Job{
		Name:       body.Name,
		Slug:       body.Slug,
		ShowInHome: body.ShowInHome,
	}
	newJob.Tags = tags
	newJob.Customer = customer
	newJob.CustomerID = customer.ID

	job, errJob := controller.jobRepo.CreateJob(&newJob)

	_, errClear, errAppend := controller.jobRepo.AppendTag(job, tags)

	if errClear != nil {
		return c.Status(http.StatusBadRequest).JSON(errClear)
	}

	if errAppend != nil {
		return c.Status(http.StatusBadRequest).JSON(errAppend)
	}

	_, errClearAward, errAppendAward := controller.jobRepo.AppendAward(job, awards)

	if errClearAward != nil {
		return c.Status(http.StatusBadRequest).JSON(errClearAward)
	}

	if errAppendAward != nil {
		return c.Status(http.StatusBadRequest).JSON(errAppendAward)
	}

	if errJob != nil {
		return c.Status(http.StatusBadRequest).JSON(errJob)
	}

	var components []models.JobComponent

	for i := 0; i < len(body.Components); i++ {
		l := &body.Components[i]

		newC := models.JobComponent{
			JobID:    job.ID,
			Type:     l.Type,
			Title:    l.Title,
			Text:     l.Text,
			Position: l.Position,
		}
		newComponent, errNewComponent := controller.jobComponentRepo.CreateJobComponent(&newC)

		if errNewComponent != nil {
			return c.Status(http.StatusBadRequest).JSON(errNewComponent)
		}

		components = append(components, *newComponent)
	}

	job.Components = &components

	update, errUpdate := controller.jobRepo.UpdateJob(job)

	if errUpdate != nil {
		return c.Status(http.StatusBadRequest).JSON(errUpdate)
	}

	return c.Status(http.StatusCreated).JSON(update)
}

func (controller jobController) ListAllJobs(c *fiber.Ctx) error {
	jobs, err := controller.jobRepo.ListAllJobs()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}
	return c.Status(http.StatusOK).JSON(&jobs)
}

func (controller jobController) GetJobBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")

	job, err := controller.jobRepo.GetJobByID(slug)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Job not found",
		})
	}
	return c.Status(http.StatusOK).JSON(job)
}

func (controller jobController) RemoveJob(c *fiber.Ctx) error {
	id := c.Params("id")

	job, errJob := controller.jobRepo.GetJobByID(id)

	if errJob != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Job not found",
		})
	}

	resultRemove, errRemove := handlers.RemoveS3(job.Placeholder.FileName)

	if errRemove != nil {
		return c.Status(http.StatusBadRequest).JSON(errRemove)
	}

	errR := controller.jobRepo.RemoveJobByID(job)

	if errR != nil {
		return c.Status(http.StatusBadRequest).JSON(errR)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"Msg":    "Job removed",
		"Result": resultRemove,
	})
}

func NewJobController() JobController {
	return &jobController{
		jobRepo:          repository.NewJobRepository(),
		customerRepo:     repository.NewCustomerRepository(),
		jobComponentRepo: repository.NewJobComponentRepository(),
		tagRepo:          repository.NewTagRepository(),
		awardRepo:        repository.NewAwardRepository(),
	}
}
