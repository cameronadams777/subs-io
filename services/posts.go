package services

import (
	"app/models"
	"app/repositories"

	"github.com/google/uuid"
)

func FindPostByID(id string) (*models.Post, error) {
	post, err := repositories.FindPost(repositories.FindPostParams{
		ID: id,
	})

	if err != nil {
		return nil, err
	}

	return post, nil
}

func FindPosts(params repositories.FindPostsParams) ([]models.Post, error) {
	posts, err := repositories.FindPosts(params)

	if err != nil {
		return nil, err
	}

	return posts, nil
}

type CreatePostParams struct {
	Title  string
	UserID string
	URL    string
	Status string
}

func CreatePost(params CreatePostParams) (*models.Post, error) {
	user_id, err := uuid.Parse(params.UserID)

	if err != nil {
		return nil, err
	}

	new_post := models.Post{
		Title:  params.Title,
		Status: params.Status,
		URL:    params.URL,
		UserID: user_id,
	}

	post, create_err := repositories.CreatePost(&new_post)

	if create_err != nil {
		return nil, err
	}

	return post, nil
}

type UpdatePostParams struct {
	ID    uuid.UUID
	Title  string
	Status string
	URL    string
	UserID uuid.UUID
}

func UpdatePost(updates UpdatePostParams) (*models.Post, error) {
	updated_post := models.Post{
    UUIDBaseModel: models.UUIDBaseModel{
      ID: updates.ID,
    },
		Title:  updates.Title,
		Status: updates.Status,
		URL:    updates.URL,
    UserID: updates.UserID,
	}

	post, err := repositories.UpdatePost(updates.ID, updated_post)

	if err != nil {
		return nil, err
	}

	return post, nil
}

type DeletePostParams struct {
	ID string
}

func DeletePost(params DeletePostParams) error {
	post_id, err := uuid.Parse(params.ID)

	if err != nil {
		return err
	}

	err = repositories.DeletePost(post_id)

	if err != nil {
		return err
	}

	return nil
}
