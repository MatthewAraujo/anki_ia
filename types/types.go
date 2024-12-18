package types

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type ValidationErrorResponse struct {
	Field      string `json:"field"`
	Validation string `json:"validation"`
	Value      string `json:"value,omitempty"`
	Message    string `json:"message"`
}

type CostumersService interface {
	CreateUser(c *CreateUserPayload) (int, error)
	Login(c *LoginUserPayload) (string, int, error)
}

type CreateUserPayload struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=100"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=100"`
}

type AnkiService interface {
	CreateAnki(a *CreateAnkiPayload) (CreateAnkiResponse, int, error)
}

type CreateAnkiPayload struct {
	File multipart.File
	Name string
}

type CreateAnkiResponse struct {
	Question []Question `json:"questions"`
}

type Question struct {
	ID           uuid.UUID         `json:"id"`
	Pergunta     string            `json:"pergunta"`
	Alternativas map[string]string `json:"alternativas"`
	Gabarito     string            `json:"gabarito"`
}
