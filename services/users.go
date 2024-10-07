package services

import (
	"app/models"
	"app/repositories"

	"github.com/google/uuid"
	"github.com/markbates/goth"
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

func FindUserByEmail(email string) (*models.User, error) {
  user, err := repositories.FindUser(repositories.FindUserParams{
    Email: email,
  })

  if err != nil {
    return nil, err
  }

  return user, nil
}

type UpdateUserParams struct {
  ID string
  GoogleUserId string
  TiktokUserId string
}

func UpdateUser(updates UpdateUserParams) (*models.User, error) {
  updated_user := models.User {
    GoogleUserId: updates.GoogleUserId,
    TiktokUserId: updates.TiktokUserId,
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

func UpsertUserByEmail(email string, user goth.User) (*models.User, error) {
  existing_user, _ := repositories.FindUser(repositories.FindUserParams{
    Email: email,
  })

  if existing_user != nil {
    if user.Provider == "google" && existing_user.GoogleUserId == "" {
      existing_user.GoogleUserId = user.UserID
      _, err := repositories.UpdateUser(existing_user.ID, *existing_user)

      if err != nil {
        return nil, err
      }
    }

    if user.Provider == "tiktok" && existing_user.TiktokUserId == "" {
      existing_user.TiktokUserId = user.UserID
      _, err := repositories.UpdateUser(existing_user.ID, *existing_user)

      if err != nil {
        return nil, err
      }
    }

    return existing_user, nil
  }

  new_user := models.User{
    Email: user.Email,
  }

  if user.Provider == "google" {
    new_user.GoogleUserId = user.UserID
  }

  if user.Provider == "tiktok" {
    new_user.TiktokUserId = user.UserID
  }

  user_create_err := repositories.CreateUser(new_user)

  if user_create_err != nil {
    return nil, user_create_err
  }

  return &new_user, nil
}
