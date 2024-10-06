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

	posts, err := services.FindUploads(repositories.FindUploadsParams{
		UserID: user_id.(string),
	})

	if err != nil {
		return c.Redirect(302, "/error")
	}

	app_context := structs.AppContext{
		Key:   "session",
		Value: structs.SessionContext{},
	}
	return render_with_context(
		c,
		upload_pages.UploadIndex(upload_pages.UploadIndexPageProps{
			Uploads: posts,
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

	post, err := services.FindUploadByID(post_id)

	if err != nil {
		return c.Redirect(302, "/error")
	}

	app_context := structs.AppContext{
		Key:   "session",
		Value: structs.SessionContext{},
	}
	return render_with_context(
		c,
		upload_pages.UploadShow(upload_pages.UploadShowPageProps{
			Upload: *post,
		}),
		app_context,
	)
}
