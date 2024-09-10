package controllers

import (
	"app/services"
	"app/views/components"
	"app/views/pages/user_pages"

	"net/http"

	"github.com/labstack/echo/v4"
)

type UsersController struct {}

func (uc *UsersController) HandleUsersEdit(c echo.Context) error {
	user_id := c.Get("user_id")

	app_context := get_app_context(c)

	user, err := services.FindUserByID(user_id.(string))

	if err != nil {
		return c.Redirect(302, "/error")
	}

	return render_with_context(c, user_pages.UserEdit(user_pages.UserEditPageProps{
		Token: "",
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
}
