package repositories

import (
	"app/database"
	"app/models"

	"github.com/google/uuid"
)

type FindPostParams struct {
	ID string
}

func FindPost(params FindPostParams) (*models.Post, error) {
	var post models.Post

	tx := database.DB

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

func FindPosts(params FindPostsParams) ([]models.Post, error) {
	var posts []models.Post

	tx := database.DB

	if params.UserID != "" {
		tx = tx.Where("user_id = ?", params.UserID)
	}

	err := tx.Find(&posts).Error

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func CreatePost(post *models.Post) (*models.Post, error) {
	err := database.DB.Create(&post).Error

	if err != nil {
		return nil, err
	}

	return post, nil
}

func UpdatePost(id uuid.UUID, updated_post models.Post) (*models.Post, error) {
	var post_to_update models.Post
	err := database.DB.First(&post_to_update, id).Error

	if err != nil {
		return nil, err
	}

	database.DB.Model(&post_to_update).Updates(updated_post)

	return &post_to_update, nil
}

func DeletePost(id uuid.UUID) error {
	var post_to_delete models.Post
	err := database.DB.First(&post_to_delete, id).Error

	if err != nil {
		return err
	}

	database.DB.Delete(&post_to_delete)

	return nil
}
