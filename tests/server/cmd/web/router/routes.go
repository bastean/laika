package router

import (
	"github.com/bastean/laika/test/server/cmd/web/handler"
)

func InitRoutes() {
	router.NoRoute(handler.NotRoute())

	router.GET("/", handler.IndexPage())
	router.GET("/dashboard", handler.IndexDashboard())
	router.GET("/dashboard/admin", handler.IndexDashboardAdmin())
}
