package httpapi

import (
	"github.com/gofiber/fiber/v2"
	"webapi/internal/app/queue"
	"webapi/internal/app/user"
	"webapi/internal/repository"

	httpAuth "webapi/internal/http/controllers/auth"
	httpHealthz "webapi/internal/http/controllers/healthz"
	httpMiscellaneous "webapi/internal/http/controllers/miscellaneous"
	httpQueue "webapi/internal/http/controllers/queue"
	httpUser "webapi/internal/http/controllers/user"
)

// ====================================================
// =================== DEFINE ROUTE ===================
// ====================================================
var repo *repository.Repository

func RegisterRoute(r *fiber.App) {
	repo = repository.NewRepository()
	api := r.Group("/api")
	v1 := api.Group("/v1")

	// Healthz API
	healthAPI := api.Group("/healthz")
	healthHandler := httpHealthz.NewHealthzHTTPHandler()
	healthAPI.Get("/", healthHandler.Healthz)

	// User API
	userAPI := v1.Group("/users")
	userApp := user.NewUserApp(repo)
	userHandler := httpUser.NewUserHTTPHandler(userApp)
	userAPI.Get("/", userHandler.GetUsers)
	userAPI.Get("/:id", userHandler.GetUserByID)
	userAPI.Post("/", userHandler.CreateUser)
	userAPI.Put("/:id", userHandler.UpdateUser)
	userAPI.Delete("/:id", userHandler.DeleteUser)

	// auth
	authApi := v1.Group("/auth")
	registerHandler := httpAuth.NewRegisterHTTPHandler(userApp)
	authApi.Post("/register", registerHandler.Register)

	// Queue API
	queueAPI := v1.Group("/queues")
	queueApp := queue.NewQueueApp(repo)
	queueHandler := httpQueue.NewQueueHTTPHandler(queueApp)
	queueAPI.Get("/", queueHandler.GetQueues)
	// queueAPI.Get("/:key", queueHandler.GetQueueByKey)

	// Error Case Handler
	miscellaneousHandler := httpMiscellaneous.NewMiscellaneousHTTPHandler()
	r.All("*", miscellaneousHandler.NotFound)
}
