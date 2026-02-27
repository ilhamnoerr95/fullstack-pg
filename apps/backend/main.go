package main

import (
	"fmt"
	"log"
	"net/http"

	"backend/internal/handler"
	"backend/internal/middleware"
	"backend/internal/repository"
	"backend/internal/service"

	"github.com/gin-gonic/gin"
)


func handlerTest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Go Backend bener boi"))
	fmt.Fprintln(w, "OK") 

}

func main() {
	log.Println("Backend running on :8080")

	r := gin.Default()

	// Serve Swagger UI and OpenAPI spec
	r.StaticFile("/docs", "./docs/swagger.html")
	r.StaticFile("/openapi.yaml", "./openapi.yaml")

	v1 := r.Group("/dashboard/v1")

	// ===== Repository Layer =====
	userRepo := repository.NewUserRepository()
	paymentRepo := repository.NewPaymentRepository()

	// ===== Service Layer =====
	authService := service.NewAuthService(userRepo)
	paymentService := service.NewPaymentService(paymentRepo)

	// ===== Handler Layer =====
	authHandler := handler.NewAuthHandler(authService)
	paymentHandler := handler.NewPaymentHandler(paymentService)

	// ===== Routes =====
	authGroup := v1.Group("/auth")
	authGroup.POST("/login", authHandler.Login)

	api := v1.Group("/payments")
	api.Use(middleware.AuthMiddleware())
	{
		api.GET("", paymentHandler.GetAll)

	}

	r.Run(":8080")
}
