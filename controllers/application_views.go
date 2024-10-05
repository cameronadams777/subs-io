package controllers

import (
	"app/structs"
	"app/views/pages"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

type ApplicationViewHandler struct{}

func (av *ApplicationViewHandler) HandleHomeIndex(c echo.Context) error {
	session, err := gothic.Store.Get(c.Request(), "subs_io_session")
	if err != nil {
		app_context := structs.AppContext{
			Key: "session",
			Value: structs.SessionContext{
				UserID: "",
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

	u := session.Values["user"]
	if u == nil {
		app_context := structs.AppContext{
			Key: "session",
			Value: structs.SessionContext{
				UserID: "",
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

	app_context := structs.AppContext{
		Key: "session",
		Value: structs.SessionContext{
			UserID: (u.(goth.User)).UserID,
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
