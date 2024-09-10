package services

import (
	"app/models"
	"app/repositories"

	"github.com/google/uuid"
)

func FindUserByID(id string) (*models.User, error) {
  user, err := repositories.FindUser(repositories.FindUserParams{
    ID: id,
  })

  if err != nil {
    return nil, err
  }

  return user, nil
}

type UpdateUserParams struct {
  ID string `form:"id"`
  FirstName string `form:"first_name"`
  LastName string `form:"last_name"`
  Email string `form:"email"`
}

func UpdateUser(updates UpdateUserParams) (*models.User, error) {
  updated_user := models.User {
    FirstName: updates.FirstName,
    LastName: updates.LastName,
    Email: updates.Email,
  }

  user_id, err := uuid.Parse(updates.ID)

  if err != nil {
    return nil, err
  }

  user, err := repositories.UpdateUser(user_id, updated_user)

  if err != nil {
    return nil, err
  }

  return user, nil
}
