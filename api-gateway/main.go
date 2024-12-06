package main

import (
	"api-gateway/clients"
	"api-gateway/handlers"
	"api-gateway/utils"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	godotenv.Load()

	handler := handlers.Handlers{
		UserClient:    clients.NewUserClient(),
		ProfileClient: clients.NewProfileClient(),
		PaymentClient: clients.NewPaymentClient(),
		DateClient:    clients.NewDateClient(),
	}

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	//set error handler
	e.HTTPErrorHandler = utils.ErrorHandler

	//user-subs
	userSubs := e.Group("/subscribe")
	userSubs.POST("", handler.HandleCreateUserSubscription)
	//callback for xendit
	e.POST("/payment-callback", handler.HandlePaymentCallback)

	//user
	users := e.Group("/users")
	users.POST("/register", handler.HandleRegister)
	users.POST("/login", handler.HandleLogin)

	//profile
	profiles := e.Group("/profiles")
	profiles.POST("", handler.HandleCreateProfile)
	profiles.GET("/:id", handler.HandleGetProfile)
	profiles.PUT("/:id", handler.HandleUpdateProfile)
	profiles.DELETE("/:id", handler.HandleDeleteProfile)

	//date
	swipes := e.Group("/swipes")
	swipes.POST("", handler.HandleRecordSwipe)
	swipes.GET("", handler.HandleGetSuggestions)
	swipes.GET("/history", handler.HandleSwipeHistory)

	//start server
	log.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
