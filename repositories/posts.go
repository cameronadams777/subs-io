package repositories

import (
	"app/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostRepository struct {
	DB *gorm.DB
}

type FindPostParams struct {
	ID    string
}

func (ur *PostRepository) FindPost(params FindPostParams) (*models.Post, error) {
	var post models.Post

  tx := ur.DB

  if params.ID != "" {
    tx = tx.Where("id = ?", params.ID)
  }

  err := tx.First(&post).Error

	if err != nil {
		return nil, err
	}

	return &post, nil
}

type FindPostsParams struct {
  UserID string
}

func (ur *PostRepository) FindPosts(params FindPostsParams) ([]models.Post, error) {
  var posts []models.Post

  tx := ur.DB

  if params.UserID != "" {
    tx = tx.Where("user_id = ?", params.UserID)
  }

  err := tx.Find(&posts).Error

  if err != nil {
    return nil, err
  }

  return posts, nil
}

func (ur *PostRepository) CreatePost(post models.Post) error {
  err := ur.DB.Create(&post).Error

  if err != nil {
    return err
  }

  return nil
}

func (ur *PostRepository) Update(id uuid.UUID, updated_post models.Post) (*models.Post, error) {
  var post_to_update models.Post
  err := ur.DB.First(&post_to_update, id).Error

  if err != nil {
    return nil, err
  }

  ur.DB.Model(&post_to_update).Updates(updated_post)

  return &post_to_update, nil
}

func (ur *PostRepository) Delete(id uuid.UUID) error {
  var post_to_delete models.Post
  err := ur.DB.First(&post_to_delete, id).Error

  if err != nil {
    return err
  }

  ur.DB.Delete(&post_to_delete)

  return nil
}
