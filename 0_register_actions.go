package main

import (
	"github.com/labstack/echo/v4"
)

func registerActions(e *echo.Echo) {

	// COMMAND INJECTION
	e.POST("/command-injection", handleCommandInjectionAction)

	// SQL INJECTION
	e.POST("/sql-injection", handleSQLInjectionAction)

	// CROSS SITE SCRIPTING
	e.POST("/cross-site-scripting", handleCrossSiteScriptingAction)

	// COOKIE SECURITY
	e.POST("/cookie-security", handleCookieSecurityAction)

	// CREDENTIAL STUFFING
	e.POST("/credential-stuffing", handleCrendentialStuffingAction)

	// WEB SCANNER
	e.POST("/web-scan", handleNiktoWebScanAction)
	
	// BOT DECEPTION
	e.GET("/view-page-source", handleViewPageSourceAction)
	e.GET("/bot-deception", handleBotDeceptionAction)

	// REST API
	e.POST("/create-policy", onboardNewApplicationPolicy)
	e.POST("/delete-policy", deleteApplicationPolicy)

	// HEALTH CHECK
	e.GET("/health-check", handleHealthCheckAction)

	// PING
	e.POST("/ping", handlePingAction)
}
