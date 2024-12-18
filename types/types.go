package types

import (
	"mime/multipart"
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
	File   multipart.File `json:"file"`
	Name   string
	UserID int32
}

type CreateAnkiResponse struct {
	Question []Question `json:"questions"`
}

type Question struct {
	Question     string            `json:"question"`
	Alternatives map[string]string `json:"alternatives"`
	Right_answer string            `json:"right_answer"`
}
