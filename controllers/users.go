package controllers

import (
	"app/services"
)

type UsersController struct {
  auth services.AuthService
}

/*func (uc *UsersController) HandleUsersEdit(c echo.Context) error {
  user, err := uc.auth.GetSessionUser(c.Request())

  if err != nil {
    fmt.Println(err)
    return c.Redirect(http.StatusTemporaryRedirect, "/error")
  }

	app_context := get_app_context(c)

	user, fetch_user_err := services.FindUserByID(&user.UserID)

  if fetch_user_err != nil {
    fmt.Println(err)
    return c.Redirect(http.StatusTemporaryRedirect, "/error")
  }

	if err != nil {
		return c.Redirect(302, "/error")
	}

	return render_with_context(c, user_pages.UserEdit(user_pages.UserEditPageProps{
		Token: c.Get(middleware.DefaultCSRFConfig.ContextKey).(string),
		User:  *user,
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
}*/
