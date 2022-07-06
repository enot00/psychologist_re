package main

import (
	"context"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/test_server/internal/infra/middlewares"
	"path/filepath"

	"github.com/test_server/config"
	"github.com/test_server/internal/app"
	"github.com/test_server/internal/infra/database"
	"github.com/test_server/internal/infra/http/controllers"
	"github.com/upper/db/v4/adapter/postgresql"

	"log"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	"github.com/test_server/internal/infra/http"
)

// @title                       Psychology API
// @version                     0.0.1
// @description                 REST Service for Psychology application

func main() {
	exitCode := 0
	ctx, cancel := context.WithCancel(context.Background())

	// Recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("The system panicked!: %v\n", r)
			fmt.Printf("Stack trace form panic: %s\n", string(debug.Stack()))
			exitCode = 1
		}
		os.Exit(exitCode)
	}()

	// Signals
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-c
		fmt.Printf("Received signal '%s', stopping... \n", sig.String())
		cancel()
		fmt.Printf("Sent cancel to all threads...")
	}()

	var conf = config.GetConfiguration()

	sess, err := postgresql.Open(
		postgresql.ConnectionURL{
			User:     conf.DatabaseUser,
			Host:     conf.DatabaseHost,
			Password: conf.DatabasePassword,
			Database: conf.DatabaseName,
		})
	if err != nil {
		log.Fatalf("Unable to create new DB session: %q\n", err)
	}

	_, err = os.Stat(conf.FileStorageLocation)
	if err != nil {
		err = os.MkdirAll(conf.FileStorageLocation+"/tmp", os.ModePerm)
	}

	if err != nil {
		log.Fatalf("Storage folder is not available %s", err)
	}

	// ===== Our code here =====
	e, err := casbin.NewEnforcer(
		filepath.Join(conf.ConfigPath, "auth/model.conf"),
		filepath.Join(conf.ConfigPath, "auth/policy.csv"),
	)
	if err != nil {
		log.Fatalf("Unable to create new Casbin enforcer: %q\n", err)
	}

	hashService := app.NewHashService(conf.Salt)
	storageService := app.NewStorageService(conf.FileStorageLocation)
	userRepository := database.NewUserRepository(&sess)
	userService := app.NewUserService(&userRepository, &hashService)
	userController := controllers.NewUserController(&userService)

	tokenRepository := database.NewTokenRepository(&sess)
	tokenService := app.NewTokenService(&tokenRepository, conf.JWTSecretKey, conf.JWTTokenTTL)

	authService := app.NewAuthenticationService(&userService, &tokenService, &hashService)

	md := middlewares.NewMiddlewares(e, &tokenService, &userService)
	authController := controllers.NewApiAuthenticationController(&authService)

	messageRepository := database.NewMessageRepository(&sess)
	messageService := app.NewMessageService(&messageRepository)
	messageController := controllers.NewMessageController(&messageService, storageService)
	// =====
	chatRepository := database.NewChatRepository(&sess)
	chatService := app.NewChatService(&chatRepository)
	chatController := controllers.NewChatController(chatService)
	// =====               =====

	psychologistsRepository := database.NewClientPsychologistRepository(&sess)
	psychologistsService := app.NewPsychologistService(&psychologistsRepository)
	psychologistsController := controllers.NewPsychologistsController(&psychologistsService)

	workingHoursRepository := database.NewWorkingHoursRepository(&sess)
	workingHoursService := app.NewWorkingHoursService(&workingHoursRepository)
	workingHoursController := controllers.NewWorkingHoursController(&workingHoursService)

	meetingRepository := database.NewMeetingRepository(&sess)
	meetingService := app.NewMeetingService(&meetingRepository)
	meetingController := controllers.NewMeetingController(&meetingService)

	clientRepository := database.NewClientRepository(&sess)
	clientService := app.NewClientService(&clientRepository)
	clientController := controllers.NewClientController(&clientService)

	// HTTP Server
	err = http.Server(
		ctx,
		http.Router(
			authController,
			meetingController,
			clientController,
			psychologistsController,
			userController,
			workingHoursController,
			chatController,
			messageController,
			md,
		),
	)

	if err != nil {
		fmt.Printf("http server error: %s", err)
		exitCode = 2
		return
	}
}
