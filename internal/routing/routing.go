package routing

import (
	"github.com/gyanesh-mishra/slackbot-winston/internal/routing/handlers/health"
	"github.com/gyanesh-mishra/slackbot-winston/internal/routing/handlers/questionanswer"
	"github.com/gyanesh-mishra/slackbot-winston/internal/routing/handlers/slackevents"
	"github.com/gyanesh-mishra/slackbot-winston/internal/routing/handlers/slackinteractions"

	"github.com/julienschmidt/httprouter"
)

// GetRouter returns a router object with mapping of URLs to functions
func GetRouter() *httprouter.Router {
	router := httprouter.New()

	// Handle / path
	router.GET("/", health.HandleGet)

	// Handle /slack-event path
	router.POST("/slack-event", slackevents.HandlePost)

	// Handle /slack-interactions path
	router.POST("/slack-interactions", slackinteractions.HandlePost)

	// Handle /question-answer path
	router.GET("/question-answer", questionanswer.HandleGet)
	router.POST("/question-answer", questionanswer.HandlePost)

	return router
}
