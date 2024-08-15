package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/atfilms/controllers"
	"github.com/viniblima/atfilms/middlewares"
)

type UploadRouter interface {
	SetupUploadRoutes(api fiber.Router)
}

type uploadRouter struct {
	uploadController controllers.UploadController
	middleware       middlewares.JWTMiddleware
}

func SetupUploadRoutes(api fiber.Router) {
	router := &uploadRouter{
		uploadController: controllers.NewUploadController(),
		middleware:       middlewares.NewJwtMiddleware(),
	}

	upload_routes := api.Group("/upload")                                                                        // Configuracao da rota pai
	upload_routes.Post("/customer/:id", router.middleware.VerifyJWT, router.uploadController.UploadCustomerLogo) // Upload de foto

	upload_routes.Post("/award/:id", router.middleware.VerifyJWT, router.uploadController.UploadAwardImage)

	upload_routes.Post("/job/placeholder/:id", router.middleware.VerifyJWT, router.uploadController.UploadJobPhoto) // Upload de placeholder
	upload_routes.Post("/job/video/:id", router.middleware.VerifyJWT, router.uploadController.UploadJobVideo)       // Upload de video

	upload_routes.Post("job-component/fill-photo-horizontal/:id", router.middleware.VerifyJWT, router.uploadController.UploadJobComponentFillPhotoHorizontal) // Upload de foto que preenche o componente

	upload_routes.Post("job-component/slider/:id", router.middleware.VerifyJWT, router.uploadController.UploadJobComponentPhotoSlider) // Upload de foto à ser adicionado ao slider de fotos
	upload_routes.Put("job-component/slider/:id", router.middleware.VerifyJWT, router.uploadController.UpdatePhotoPositionInSlider)    // Altera posição de foto no slider de fotos
	upload_routes.Delete("job-component/slider/:id", router.middleware.VerifyJWT, router.uploadController.RemovePhotoFromSlider)       // Exclui foto do slider

	upload_routes.Post("job-component/videos/:id", router.middleware.VerifyJWT, router.uploadController.UploadJobComponentVideo)    // Upload de vídeo à ser adicionado ao slider de vídeos
	upload_routes.Put("job-component/videos/:id", router.middleware.VerifyJWT, router.uploadController.UpdateVideoPositionInSlider) // Upload de video à ser adicionado à lista de videos
	upload_routes.Delete("job-component/videos/:id", router.middleware.VerifyJWT, router.uploadController.RemoveVideoFromList)      // Exclui vídeo da lista do componente

}
