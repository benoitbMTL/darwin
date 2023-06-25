package main

import (
	"darwin/actions"
	"github.com/labstack/echo/v4"
)

func registerActions(e *echo.Echo) {

	// COMMAND INJECTION
	e.POST("/command-injection", actions.handleCommandInjectionAction)

	// SQL INJECTION
	e.POST("/sql-injection", actions.handleSQLInjectionAction)

	// BOT DECEPTION
	e.GET("/view-page-source", actions.handleViewPageSourceAction)
	e.GET("/bot-deception", actions.handleBotDeceptionAction)

	// HEALTH CHECK
	e.GET("/health-check", actions.handleHealthCheckAction)

	// PING
	e.POST("/ping", actions.handlePingAction)
}
