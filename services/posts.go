package services

import (
	"app/models"
	"app/repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostService struct {
	DB *gorm.DB
}

func (ps *PostService) FindByID(id string) (*models.Post, error) {
	post_repo := repositories.PostRepository{
		DB: ps.DB,
	}

	post, err := post_repo.FindPost(repositories.FindPostParams{
		ID: id,
	})

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (ps *PostService) Find(params repositories.FindPostsParams) ([]models.Post, error) {
	post_repo := repositories.PostRepository{
		DB: ps.DB,
	}

	posts, err := post_repo.FindPosts(params)

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

func (ps *PostService) Create(params CreatePostParams) (*models.Post, error) {
	post_repo := repositories.PostRepository{
		DB: ps.DB,
	}

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

	create_err := post_repo.CreatePost(new_post)

	if create_err != nil {
		return nil, err
	}

	return &new_post, nil
}

type UpdatePostParams struct {
	ID    uuid.UUID
	Title  string
	Status string
	URL    string
	UserID string
}

func (ps *PostService) Update(updates UpdatePostParams) (*models.Post, error) {
	post_repo := repositories.PostRepository{
		DB: ps.DB,
	}

	updated_post := models.Post{
		Title:  updates.Title,
		Status: updates.Status,
		URL:    updates.URL,
	}

	post, err := post_repo.Update(updates.ID, updated_post)

	if err != nil {
		return nil, err
	}

	return post, nil
}

type DeletePostParams struct {
	ID string
}

func (ps *PostService) Delete(params DeletePostParams) error {
	post_repo := repositories.PostRepository{
		DB: ps.DB,
	}

	post_id, err := uuid.Parse(params.ID)

	if err != nil {
		return err
	}

	err = post_repo.Delete(post_id)

	if err != nil {
		return err
	}

	return nil
}
