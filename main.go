package main

import (
	"context"
	"hacktiv8_fp_2/config"
	"hacktiv8_fp_2/controller"
	"hacktiv8_fp_2/repository"
	"hacktiv8_fp_2/routes"
	"hacktiv8_fp_2/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}
	var (
		db                *gorm.DB                     = config.SetupDatabaseConnection()
		userRepository    repository.UserRepository    = repository.NewUserRepository(db)
		photoRepository   repository.PhotoRepository   = repository.NewPhotoRepository(db)
		socmedRepository  repository.SocmedRepository  = repository.NewSocmedRepository(db)
		commentRepository repository.CommentRepository = repository.NewCommentRepository(db)

		jwtService     service.JWTService     = service.NewJWTService()
		userService    service.UserService    = service.NewUserService(userRepository)
		authService    service.AuthService    = service.NewAuthService(userRepository)
		photoService   service.PhotoService   = service.NewPhotoService(photoRepository)
		socmedService  service.SocmedService  = service.NewSocmedService(socmedRepository, userRepository)
		commentService service.CommentService = service.NewCommentService(commentRepository)

		authController    controller.UserController    = controller.NewUserController(userService, authService, jwtService)
		photoController   controller.PhotoController   = controller.NewPhotoController(photoService, jwtService)
		socmedController  controller.SocmedController  = controller.NewSocmedController(userService, socmedService, jwtService)
		commentController controller.CommentController = controller.NewCommentController(commentService, jwtService)
	)

	defer config.CloseDatabaseConnection(db)
	// gin.SetMode(gin.ReleaseMode)
	server := gin.Default()

	routes.UserRoutes(server, authController, jwtService)
	routes.SocMedRoutes(server, socmedController, socmedService, jwtService)
	routes.PhotoRoutes(server, photoController, photoService, jwtService)
	routes.CommentRoutes(server, commentController, commentService, jwtService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: server,
	}

	// Graceful shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("error serving :", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Printf("[%v] - Shutting down server\n", time.Now())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("error shutting down :", err)
	}

	<-ctx.Done()
	log.Println("timeout, exiting")

}
