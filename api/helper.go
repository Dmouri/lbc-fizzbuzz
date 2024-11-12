package api

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// GET registers a route that works with or without a trailing slash.
func GET(g *gin.RouterGroup, route string, handlers ...gin.HandlerFunc) {
	var otherRoute string
	if route[len(route)-1] == '/' {
		otherRoute = strings.TrimRight(route, "/")
	} else {
		otherRoute = route + "/"
	}
	g.GET(route, handlers...)
	g.GET(otherRoute, handlers...)
}
