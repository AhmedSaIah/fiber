package services

import "github.com/AhmedSaIah/fiber/models"

type AuthService interface {
	SignUpUser(*models.SignUpRequest) (*models.DBResponse, error)
	SignInUser(*models.SignInRequest) (*models.DBResponse, error)
}
