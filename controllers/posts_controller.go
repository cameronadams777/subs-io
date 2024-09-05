package controllers

import (
	"app/repositories"
	"app/services"
	"app/structs"
	"app/views/pages/post_pages"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type PostsController struct {
  DB *gorm.DB
}

func (pc *PostsController) HandlePostsIndex(c echo.Context) error {
  user_id := c.Get("user_id")

  if user_id == nil {
    user_id = c.Get("uid").(string)
  }

  posts_service := services.PostService{
    DB: pc.DB,
  }

  posts, err := posts_service.Find(repositories.FindPostsParams{
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
    post_pages.PostIndex(post_pages.PostIndexPageProps{
      Posts: posts,
    }),
    app_context,
  )
}
