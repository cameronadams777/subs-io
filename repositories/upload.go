package repositories

import (
	"app/database"
	"app/models"

	"github.com/google/uuid"
)

type FindUploadParams struct {
	ID string
}

func FindUpload(params FindUploadParams) (*models.Upload, error) {
	var post models.Upload

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

type FindUploadsParams struct {
	UserID string
}

func FindUploads(params FindUploadsParams) ([]models.Upload, error) {
	var posts []models.Upload

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

func CreateUpload(post *models.Upload) (*models.Upload, error) {
	err := database.DB.Create(&post).Error

	if err != nil {
		return nil, err
	}

	return post, nil
}

func UpdateUpload(id uuid.UUID, updated_post models.Upload) (*models.Upload, error) {
	var post_to_update models.Upload
	err := database.DB.First(&post_to_update, id).Error

	if err != nil {
		return nil, err
	}

	database.DB.Model(&post_to_update).Updates(updated_post)

	return &post_to_update, nil
}

func DeleteUpload(id uuid.UUID) error {
	var post_to_delete models.Upload
	err := database.DB.First(&post_to_delete, id).Error

	if err != nil {
		return err
	}

	database.DB.Delete(&post_to_delete)

	return nil
}
