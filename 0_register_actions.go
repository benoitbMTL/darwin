package main

import (
	"net/http"

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

	// PETSTORE API PROTECTION
	e.POST("/petstore-pet-get", handlePetstoreAPIRequestGet)
	e.POST("/petstore-pet-post", handlePetstoreAPIRequestPost)
	e.POST("/petstore-pet-put", handlePetstoreAPIRequestPut)
	e.POST("/petstore-pet-delete", handlePetstoreAPIRequestDelete)

	// REST API
	e.POST("/create-policy", onboardNewApplicationPolicy)
	e.POST("/delete-policy", deleteApplicationPolicy)

	// HEALTH CHECK
	e.GET("/health-check", handleHealthCheckAction)

	// PING
	e.POST("/ping", handlePingAction)

	// CONFIG
	e.GET("/config", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"PETSTORE_URL": PETSTORE_URL,
		})
	})

}
