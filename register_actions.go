package main

import (
	"github.com/labstack/echo/v4"
)

func registerActions(e *echo.Echo) {

	// COMMAND INJECTION
	e.POST("/command-injection", handleCommandInjectionAction)

	// SQL INJECTION
	e.POST("/sql-injection", handleSQLInjectionAction)

	// BOT DECEPTION
	e.GET("/view-page-source", handleViewPageSourceAction)
	e.GET("/bot-deception", handleBotDeceptionAction)

	// HEALTH CHECK
	e.GET("/health-check", handleHealthCheckAction)

	// PING
	e.POST("/ping", handlePingAction)
}
