package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/Turizak/fables-be/database"
	"github.com/Turizak/fables-be/routes"
	sloggin "github.com/samber/slog-gin"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	config := sloggin.Config{
		WithRequestBody:  true,
		WithResponseBody: true,
		Filters: []sloggin.Filter{
			sloggin.IgnoreStatus(401),
			sloggin.IgnoreStatus(404),
			sloggin.IgnorePath("/account/login"),
			sloggin.IgnorePath("/account/token/refresh"),
			sloggin.IgnorePath("/account/create"),
			sloggin.IgnorePath("/account/change-password")},
	}

	router := gin.New()
	router.Use(sloggin.NewWithConfig(logger, config))
	router.Use(gin.Recovery())

	database.ConnectDatabase()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}, // Change this to the origin you want to allow
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
		MaxAge:       12 * time.Hour,
	}))
	routes.Routes(router)
	router.Run(os.Getenv("PORT"))
}
