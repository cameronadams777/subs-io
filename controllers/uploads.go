package controllers

import (
	"app/repositories"
	"app/services"
	"app/structs"
	"app/views/pages/upload_pages"

	"github.com/labstack/echo/v4"
)

type UploadsController struct{}

func (uc *UploadsController) HandleUploadsIndex(c echo.Context) error {
	user_id := c.Get("user_id")

	if user_id == nil {
		user_id = c.Get("uid").(string)
	}

	posts, err := services.FindPosts(repositories.FindPostsParams{
		UserID: user_id.(string),
	})

	if err != nil {
		return c.Redirect(302, "/error")
	}

	app_context := structs.AppContext{
		Key: "session",
		Value: structs.SessionContext{
			UserID: user_id.(string),
		},
	}
	return render_with_context(
		c,
		upload_pages.PostIndex(upload_pages.PostIndexPageProps{
			Posts: posts,
		}),
		app_context,
	)
}

func (uc *UploadsController) HandleUploadsShow(c echo.Context) error {
	post_id := c.Param("id")
	user_id := c.Get("user_id")

	if user_id == nil {
		user_id = c.Get("uid").(string)
	}

	post, err := services.FindPostByID(post_id)

	if err != nil {
		return c.Redirect(302, "/error")
	}

	app_context := structs.AppContext{
		Key: "session",
		Value: structs.SessionContext{
			UserID: user_id.(string),
		},
	}
	return render_with_context(
		c,
		upload_pages.PostShow(upload_pages.PostShowPageProps{
			Post: *post,
		}),
		app_context,
	)
}
