package routes

import (
	"net/http"
)

func StartRoutes(router *http.ServeMux) {
	HandleRouteUser(router)
}