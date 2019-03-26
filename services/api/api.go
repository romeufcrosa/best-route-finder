package api

import (
	"net/http"

	v1 "github.com/romeufcrosa/best-route-finder/services/api/v1"

	"github.com/julienschmidt/httprouter"
)

// Router starts a new HTTP Router
func Router() http.Handler {
	// Logging := middlewares.Logging
	params := v1.Params{}

	router := httprouter.New()

	router.POST("/api/v1/nodes", v1.AddNode(params))
	router.POST("/api/v1/edges", v1.AddEdge(params))
	router.GET("/api/v1/routes/from/:from/to/:to", v1.GetRoute(params))
	router.NotFound = http.NotFoundHandler()

	return router
}
