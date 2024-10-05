package main

import (
	config "app"
	"app/controllers"
	"app/database"
	"app/services"

	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

  auth_service := services.NewAuthService(store)

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


  oauth_controller := controllers.OAuthController{}
	app.GET("/auth/:provider/callback", oauth_controller.HandleOAuthCallback)
	app.GET("/logout/:provider", oauth_controller.HandleOAuthLogout)
  app.GET("/auth/:provider", oauth_controller.HandleOAuthIndex)

	users := app.Group("/users")
	users_controller := controllers.UsersController{
    auth_service: auth_service,
  }
	users.GET("/edit", services.RequireAuth(users_controller.HandleUsersEdit, auth_service))
	users.PATCH("/update", services.RequireAuth(users_controller.HandleUsersUpdate, auth_service))

	subtitles := app.Group("/subtitles")
	subtitles_controller := controllers.SubtitlesController{}
	subtitles.POST("/create", subtitles_controller.HandleSubtitlesCreate)

	uploads := app.Group("/uploads")
	uploads_controller := controllers.UploadsController{}
	uploads.GET("", uploads_controller.HandleUploadsIndex)
	uploads.GET("/:id", uploads_controller.HandleUploadsShow)

	app.Logger.Fatal(app.Start(":4000"))
}
