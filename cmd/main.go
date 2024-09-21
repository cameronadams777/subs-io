package main

import (
	config "app"
	"app/controllers"
	"app/database"
	middleware_handlers "app/middleware"

	"encoding/json"
	"os"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/providers/tiktok"
)

func main() {
	app := echo.New()

	database.ConnectDB()

	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${method} ${uri} ${status}\n",
	}))

	store := sessions.NewCookieStore([]byte(config.GetConfig("SESSION_SECRET")))

	app.Use(session.Middleware(store))
	app.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:_csrf",
	}))

	app.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "healthy!")
	})

	app.Static("/assets", "assets")

	application_views_controller := controllers.ApplicationViewHandler{}
	app.GET("/error", application_views_controller.HandleErrorIndex)
	app.GET("/not_found", application_views_controller.HandleNotFoundIndex)
	app.GET("/", application_views_controller.HandleHomeIndex)

	authentication_controller := controllers.AuthenticationController{}
	app.GET("/login", authentication_controller.HandleLoginIndex)

	gothic.Store = store

	goth.UseProviders(
		google.New(
      config.GetConfig("GOOGLE_CLIENT_ID"),
      config.GetConfig("GOOGLE_CLIENT_SECRET"),
      "http://localhost:4000/auth/google/callback",
    ),
    tiktok.New(
      config.GetConfig("TIKTOK_CLIENT_ID"),
      config.GetConfig("TIKTOK_CLIENT_SECRET"),
      "http://localhost:4000/auth/tiktok/callback",
    ),
	)

  oauth_controller := controllers.OAuthController{}
	app.GET("/auth/:provider/callback", oauth_controller.HandleOAuthCallback)
	app.GET("/logout/:provider", oauth_controller.HandleOAuthLogout)
  app.GET("/auth/:provider", oauth_controller.HandleOAuthIndex)

	app.Use(middleware_handlers.SetSessionInContext)
	app.Use(middleware_handlers.NoSessionRedirect)

	users := app.Group("/users")
	users_controller := controllers.UsersController{}
	users.GET("/edit", users_controller.HandleUsersEdit)
	users.PATCH("/update", users_controller.HandleUsersUpdate)

	subtitles := app.Group("/subtitles")
	subtitles_controller := controllers.SubtitlesController{}
	subtitles.POST("/create", subtitles_controller.HandleSubtitlesCreate)

	uploads := app.Group("/uploads")
	uploads_controller := controllers.UploadsController{}
	uploads.GET("", uploads_controller.HandleUploadsIndex)
	uploads.GET("/:id", uploads_controller.HandleUploadsShow)

  data, err := json.MarshalIndent(app.Routes(), "", "  ")

  if err != nil {
    panic(err)
  }

  os.WriteFile("routes.json", data, 0644)

	app.Logger.Fatal(app.Start(":4000"))
}
