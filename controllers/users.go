package controllers

import (
	"app/services"
	"app/views/components"
	"app/views/pages/user_pages"
	"fmt"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type UsersController struct {
  AuthService services.AuthService
}

func (uc *UsersController) HandleUsersEdit(c echo.Context) error {
  user, err := uc.AuthService.GetSessionUser(c.Request())

  if err != nil {
    fmt.Println(err)
    return c.Redirect(http.StatusTemporaryRedirect, "/error")
  }

  app_context := get_app_context(c)

  db_user, fetch_user_err := services.FindUserByEmail(user.Email)

  if fetch_user_err != nil {
    fmt.Println(err)
    return c.Redirect(http.StatusTemporaryRedirect, "/error")
  }

	return render_with_context(c, user_pages.UserEdit(user_pages.UserEditPageProps{
		Token: c.Get(middleware.DefaultCSRFConfig.ContextKey).(string),
    User: *db_user,
	}), app_context)
}

func (uc *UsersController) HandleUsersUpdate(c echo.Context) error {
	var form services.UpdateUserParams

	if err := c.Bind(&form); err != nil {
		return render(c, components.FlashMessage(components.FlashMessageProps{
			Message: "Invalid form data",
		}))
	}

	_, err := services.UpdateUser(form)

	if err != nil {
		return render(c, components.FlashMessage(components.FlashMessageProps{
			Message: "Invalid form data",
		}))
	}

	c.Response().Header().Set("HX-Location", "/users/edit")
	return c.String(http.StatusOK, "")
}
