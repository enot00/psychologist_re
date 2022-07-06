package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/test_server/internal/infra/middlewares"
	"net/http"

	"github.com/go-chi/cors"
	"github.com/test_server/internal/infra/http/controllers"

	"github.com/go-chi/chi/v5/middleware"
)

func Router(
	authController *controllers.ApiAuthenticationController,
	meetingController *controllers.MeetingController,
	clientController *controllers.ClientController,
	psychologistController *controllers.PsychologistsController,
	userController *controllers.UserController,
	workingHoursController *controllers.WorkingHoursController,
	chatController *controllers.ChatController,
	messageController *controllers.MessageController,
	md *middlewares.Middlewares) http.Handler {

	router := chi.NewRouter()

	// Health
	router.Group(func(healthRouter chi.Router) {
		healthRouter.Use(middleware.RedirectSlashes)

		healthRouter.Route("/ping", func(healthRouter chi.Router) {
			healthRouter.Get("/", PingHandler())

			healthRouter.Handle("/*", NotFoundJSON())
		})
	})

	router.Group(func(apiRouter chi.Router) {
		apiRouter.Use(middleware.RedirectSlashes, cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300,
		}))

		apiRouter.Route("/v1", func(apiRouter chi.Router) {
			AuthRoutes(&apiRouter, authController, md)

			apiRouter.Group(func(apiRouter chi.Router) {
				apiRouter.Use(md.Authorization)

				ChatRoutes(&apiRouter, chatController)
				MessageRoutes(&apiRouter, messageController)
				MeetingRouter(&apiRouter, meetingController)
				ClientRouter(&apiRouter, clientController)
				PsychologistsRouter(&apiRouter, psychologistController)
				UserRouter(&apiRouter, userController)
				WorkingHoursRouter(&apiRouter, workingHoursController)

				apiRouter.Handle("/*", NotFoundJSON())
			})
			apiRouter.Handle("/*", NotFoundJSON())
		})
	})

	return router
}

func AuthRoutes(router *chi.Router, controller *controllers.ApiAuthenticationController, md *middlewares.Middlewares) {
	(*router).Route("/auth", func(apiRouter chi.Router) {
		apiRouter.Post("/signup", controller.SignUp())
		apiRouter.Post("/login", controller.Login())

		apiRouter.Group(func(r chi.Router) {
			r.Use(md.Authorization)
			r.Get("/refresh", controller.Logout())
			r.Get("/logout", controller.Refresh())
		})
	})
}

func ChatRoutes(router *chi.Router, controller *controllers.ChatController) {
	(*router).Route("/chats", func(apiRouter chi.Router) {
		apiRouter.Post("/", controller.Save())
		apiRouter.Get("/{chatId}", controller.FindOne())
		apiRouter.Delete("/{chatId}", controller.Delete())
		apiRouter.Get("/{userId}/user", controller.FindUserChats())
	})
}

func MessageRoutes(router *chi.Router, messageController *controllers.MessageController) {
	(*router).Route("/messages", func(apiRouter chi.Router) {
		apiRouter.Post("/file", messageController.SaveFile())
		apiRouter.Post(
			"/",
			messageController.AddMessage(),
		)
		apiRouter.Put(
			"/",
			messageController.UpdateMessage(),
		)
		apiRouter.Get(
			"/{page}/{chatID}",
			messageController.PaginateAll(),
		)
		apiRouter.Delete(
			"/{id}",
			messageController.DeleteMessage(),
		)
	})
}

func PsychologistsRouter(router *chi.Router, psychologistsController *controllers.PsychologistsController) {
	(*router).Route("/psychologists", func(apiRouter chi.Router) {
		apiRouter.Get(
			"/",
			psychologistsController.PaginateAll(),
		)
		apiRouter.Get(
			"/{psyId}",
			psychologistsController.GetOne(),
		)
		apiRouter.Put(
			"/{psyId}",
			psychologistsController.Update(),
		)
	})
}
func UserRouter(router *chi.Router, userController *controllers.UserController) {
	(*router).Route("/users", func(apiRouter chi.Router) {
		apiRouter.Get(
			"/",
			userController.PaginateAll(),
		)
		apiRouter.Get(
			"/{userID}",
			userController.GetOneByID(),
		)
		apiRouter.Put(
			"/{userID}",
			userController.Update(),
		)
		apiRouter.Put(
			"/reset-password/{id}",
			userController.ResetPassword(),
		)
		apiRouter.Delete(
			"/{userID}",
			userController.Delete(),
		)
	})
}
func WorkingHoursRouter(router *chi.Router, workingHoursController *controllers.WorkingHoursController) {
	(*router).Route("/working_hours", func(apiRouter chi.Router) {
		apiRouter.Get(
			"/by_id_psy/{psyID}",
			workingHoursController.GetOneByID(),
		)
		apiRouter.Post(
			"/",
			workingHoursController.Save(),
		)
		apiRouter.Put(
			"/update_hours/{psyID}",
			workingHoursController.Update(),
		)
		apiRouter.Delete(
			"/delete_hours/{psyID}",
			workingHoursController.Delete(),
		)
	})
}

func MeetingRouter(router *chi.Router, meetingController *controllers.MeetingController) {
	(*router).Route("/meeting", func(apiRouter chi.Router) {
		apiRouter.Post(
			"/create_by_psychologist",
			meetingController.CreateByPsychologist(),
		)
		apiRouter.Post(
			"/create_by_client",
			meetingController.CreateByClient(),
		)
		apiRouter.Get(
			"/page/{page}",
			meetingController.PaginateAll(),
		)
		apiRouter.Get(
			"/psychologist/{psychologist_id}/page/{page}",
			meetingController.PaginateAllByPsychologist(),
		)
		apiRouter.Get(
			"/client/{client_id}/page/{page}",
			meetingController.PaginateAllByClient(),
		)
		apiRouter.Get(
			"/{id}",
			meetingController.FindOne(),
		)
		apiRouter.Put(
			"/{id}/update_by_psychologist",
			meetingController.UpdateByPsychologist(),
		)
		apiRouter.Put(
			"/{id}/update_by_client",
			meetingController.UpdateByClient(),
		)
		apiRouter.Delete(
			"/{id}/delete_by_psychologist",
			meetingController.DeleteByPsychologist(),
		)
		apiRouter.Delete(
			"/{id}/delete_by_client",
			meetingController.DeleteByClient(),
		)
	})
}

func ClientRouter(router *chi.Router, clientController *controllers.ClientController) {
	(*router).Route("/client", func(apiRouter chi.Router) {
		apiRouter.Post(
			"/",
			clientController.Save(),
		)
		apiRouter.Get(
			"/page/{page}",
			clientController.PaginateAll(),
		)
		apiRouter.Get(
			"/{id}",
			clientController.FindOne(),
		)
		apiRouter.Put(
			"/{id}",
			clientController.Update(),
		)
		apiRouter.Delete(
			"/{id}",
			clientController.Delete(),
		)
	})
}
