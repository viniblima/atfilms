package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/atfilms/controllers"
	"github.com/viniblima/atfilms/middlewares"
)

type JobRouter interface {
	SetupJobRoutes(api fiber.Router)
}

type jobRouter struct {
	jobController          controllers.JobController
	middleware             middlewares.JWTMiddleware
	jobComponentController controllers.JobComponentController
}

func SetupJobRoutes(api fiber.Router) {
	router := &jobRouter{
		jobController:          controllers.NewJobController(),
		middleware:             middlewares.NewJwtMiddleware(),
		jobComponentController: controllers.NewJobComponentController(),
	}

	job_routes := api.Group("/jobs") // Configuracao da rota pai

	job_routes.Post("/", router.middleware.VerifyJWT, router.jobController.CreateJob)      // Criacao de job
	job_routes.Get("/", router.jobController.ListAllJobs)                                  // Listagem de jobs
	job_routes.Get("/:slug", router.jobController.GetJobBySlug)                            // Detalhes de Job via slug
	job_routes.Put("/:id", router.middleware.VerifyJWT, router.jobController.UpdateJob)    // Detalhes de job
	job_routes.Delete("/:id", router.middleware.VerifyJWT, router.jobController.RemoveJob) // Remove job

	job_routes.Post("/job-component/:id", router.middleware.VerifyJWT, router.jobComponentController.Create)   // Cria um novo componente no job
	job_routes.Put("/job-component/:id", router.middleware.VerifyJWT, router.jobComponentController.Update)    // Atualiza um componente
	job_routes.Delete("/job-component/:id", router.middleware.VerifyJWT, router.jobComponentController.Remove) // Remove componente de um job

}
