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
    fmt.Println("An error occurred fetching the user from session:", err)
    return c.Redirect(http.StatusTemporaryRedirect, "/error")
  }

  app_context := get_app_context(c, uc.AuthService)

  db_user, fetch_user_err := services.FindUserByEmail(user.Email)

  if fetch_user_err != nil {
    fmt.Println("An error occurred fetching the user from the DB:", err)
    return c.Redirect(http.StatusTemporaryRedirect, "/error")
  }

	return render_with_context(c, user_pages.UserEdit(user_pages.UserEditPageProps{
		Token: c.Get(middleware.DefaultCSRFConfig.ContextKey).(string),
    User: *db_user,
	}), app_context)
}

