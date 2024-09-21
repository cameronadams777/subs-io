package services

import (
	"app/models"
	"app/repositories"

	"github.com/google/uuid"
)

func FindUploadByID(id string) (*models.Upload, error) {
	post, err := repositories.FindUpload(repositories.FindUploadParams{
		ID: id,
	})

	if err != nil {
		return nil, err
	}

	return post, nil
}

func FindUploads(params repositories.FindUploadsParams) ([]models.Upload, error) {
	posts, err := repositories.FindUploads(params)

	if err != nil {
		return nil, err
	}

	return posts, nil
}

type CreateUploadParams struct {
	Title  string
	UserID string
	URL    string
	Status string
}

func CreateUpload(params CreateUploadParams) (*models.Upload, error) {
	user_id, err := uuid.Parse(params.UserID)

	if err != nil {
		return nil, err
	}

	new_post := models.Upload{
		Title:  params.Title,
		Status: params.Status,
		URL:    params.URL,
		UserID: user_id,
	}

	post, create_err := repositories.CreateUpload(&new_post)

	if create_err != nil {
		return nil, err
	}

	return post, nil
}

type UpdateUploadParams struct {
	ID    uuid.UUID
	Title  string
	Status string
	URL    string
	UserID uuid.UUID
}

func UpdateUpload(updates UpdateUploadParams) (*models.Upload, error) {
	updated_post := models.Upload{
    UUIDBaseModel: models.UUIDBaseModel{
      ID: updates.ID,
    },
		Title:  updates.Title,
		Status: updates.Status,
		URL:    updates.URL,
    UserID: updates.UserID,
	}

	post, err := repositories.UpdateUpload(updates.ID, updated_post)

	if err != nil {
		return nil, err
	}

	return post, nil
}

type DeleteUploadParams struct {
	ID string
}

func DeleteUpload(params DeleteUploadParams) error {
	post_id, err := uuid.Parse(params.ID)

	if err != nil {
		return err
	}

	err = repositories.DeleteUpload(post_id)

	if err != nil {
		return err
	}

	return nil
}
