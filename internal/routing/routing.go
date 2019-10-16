package routing

import (
	"github.com/gyanesh-mishra/slackbot-winston/internal/routing/handlers/health"
	"github.com/gyanesh-mishra/slackbot-winston/internal/routing/handlers/root"

	"github.com/julienschmidt/httprouter"
)

func GetRouter() *httprouter.Router {
	router := httprouter.New()

	// Handle / path
	router.GET("/", root.HandleGet)
	router.POST("/", root.HandlePost)

	// Handle /health path
	router.GET("/health", health.HandleGet)

	return router
}
