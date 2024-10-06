package controllers

import (
	"app/services"
	"app/structs"
	"app/views/pages"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ApplicationViewHandler struct {
	AuthService services.AuthService
}

func (av *ApplicationViewHandler) HandleHomeIndex(c echo.Context) error {
	user, err := av.AuthService.GetSessionUser(c.Request())

	if err != nil {
		app_context := structs.AppContext{
			Key:   "session",
			Value: structs.SessionContext{},
		}
		return render_with_context(
			c,
			pages.HomeIndex(pages.HomePageProps{
				Token: c.Get(middleware.DefaultCSRFConfig.ContextKey).(string),
			}),
			app_context,
		)
	}

	app_context := structs.AppContext{
		Key: "session",
		Value: structs.SessionContext{
      User: &user,
		},
	}
	return render_with_context(
		c,
		pages.HomeIndex(pages.HomePageProps{
			Token: c.Get(middleware.DefaultCSRFConfig.ContextKey).(string),
		}),
		app_context,
	)
}

func (av *ApplicationViewHandler) HandleNotFoundIndex(c echo.Context) error {
	return render(c, pages.NotFoundIndex())
}

func (av *ApplicationViewHandler) HandleErrorIndex(c echo.Context) error {
	return render(c, pages.ErrorIndex())
}
