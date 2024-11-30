package routes

import (
	"net/http"
)

func StartRoutes(router *http.ServeMux) {
	HandleRouteTask(router)
	HandleRouteUser(router)
}