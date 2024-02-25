package router

import (
	"github.com/bastean/laika/pkg/cmd/server/handler"
)

func InitRoutes() {
	router.NoRoute(handler.NotRoute())

	router.GET("/", handler.IndexPage())
	router.GET("/dashboard", handler.IndexDashboard())
	router.GET("/dashboard/admin", handler.IndexDashboardAdmin())
}
