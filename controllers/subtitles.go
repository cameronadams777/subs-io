package controllers

import (
	"app/views/components"

	"io"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

type SubtitlesController struct {}

const MAX_UPLOAD_SIZE = 1024 * 1024 * 1024

func (sc *SubtitlesController) HandleSubtitlesCreate(c echo.Context) error {
	c.Request().Body = http.MaxBytesReader(c.Response(), c.Request().Body, MAX_UPLOAD_SIZE)
	if err := c.Request().ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		return render(c, components.FlashMessage(components.FlashMessageProps{
			Message: err.Error(),
		}))
	}

	file, err := c.FormFile("video")
	if err != nil {
		return render(c, components.FlashMessage(components.FlashMessageProps{
			Message: err.Error(),
		}))
	}

	src, err := file.Open()
	if err != nil {
		return render(c, components.FlashMessage(components.FlashMessageProps{
			Message: err.Error(),
		}))
	}
	defer src.Close()

	file_name := strings.ToLower(strings.ReplaceAll(file.Filename, " ", "_"))

	// Destination
	dst, err := os.Create("uploads/" + file_name)
	if err != nil {
		return render(c, components.FlashMessage(components.FlashMessageProps{
			Message: err.Error(),
		}))
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return render(c, components.FlashMessage(components.FlashMessageProps{
			Message: err.Error(),
		}))
	}

  c.Response().Header().Set("HX-Location", "/")
  return c.String(http.StatusOK, "")
}
