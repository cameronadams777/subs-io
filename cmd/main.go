package main

import (
	config "app"
	"app/controllers"
	"app/database"
	middleware_handlers "app/middleware"
	"context"
	"fmt"
	"html/template"

	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "github.com/markbates/goth"
  "github.com/markbates/goth/gothic"
  "github.com/markbates/goth/providers/tiktok"
  "github.com/markbates/goth/providers/google"
)

func main() {
	app := echo.New()

	database.ConnectDB()

	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${method} ${uri} ${status}\n",
	}))
	app.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
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

	store := sessions.NewCookieStore([]byte("secret"))

	gothic.Store = store

	goth.UseProviders(
		google.New(config.GetConfig("GOOGLE_CLIENT_ID"), config.GetConfig("GOOGLE_CLIENT_SECRET"), "http://localhost:4000/auth/google/callback"),
    tiktok.New(config.GetConfig("TIKTOK_CLIENT_ID"), config.GetConfig("TIKTOK_CLIENT_SECRET"), "http://localhost:4000/auth/tiktok/callback"),
	)

	app.GET("/auth/:provider/callback", func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), gothic.ProviderParamKey, c.Param("provider"))

		_, err := gothic.CompleteUserAuth(c.Response(), c.Request().WithContext(ctx))
		if err != nil {
			return err
		}

    return c.Redirect(http.StatusPermanentRedirect, "/")
	})

	app.GET("/logout/:provider", func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), gothic.ProviderParamKey, c.Param("provider"))

		gothic.Logout(c.Response(), c.Request().WithContext(ctx))
		return c.Redirect(http.StatusPermanentRedirect, "/")
	})

	app.GET("/auth/:provider", func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), gothic.ProviderParamKey, c.Param("provider"))

		// try to get the user without re-authenticating
		if _, err := gothic.CompleteUserAuth(c.Response(), c.Request().WithContext(ctx)); err == nil {
      return c.Redirect(http.StatusPermanentRedirect, "/")
		}

		gothic.BeginAuthHandler(c.Response(), c.Request().WithContext(ctx))

		return nil
	})

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

	app.Logger.Fatal(app.Start(":4000"))
}
