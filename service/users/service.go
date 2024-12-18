package users

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	configs "github.com/MatthewAraujo/anki_ia/config"
	"github.com/MatthewAraujo/anki_ia/repository"
	"github.com/MatthewAraujo/anki_ia/service/auth"
	"github.com/MatthewAraujo/anki_ia/types"
	"github.com/MatthewAraujo/anki_ia/utils"
	"github.com/go-playground/validator/v10"
)

type Service struct {
	db *repository.Queries

	dbTx *sql.DB
}

func NewService(db *repository.Queries, dbTx *sql.DB) *Service {
	return &Service{
		db:   db,
		dbTx: dbTx,
	}
}

func (s *Service) BeginTransaction(ctx context.Context) (*repository.Queries, *sql.Tx, error) {
	tx, err := s.dbTx.BeginTx(ctx, nil)

	if err != nil {
		return nil, nil, err
	}

	defer tx.Rollback()

	return s.db.WithTx(tx), tx, nil
}

func (s *Service) CreateUser(user *types.CreateUserPayload) (int, error) {
	logger.Info("Validating users")
	if err := utils.Validate.Struct(user); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			errorMessages := utils.TranslateValidationErrors(validationErrors)

			response, _ := json.Marshal(errorMessages)
			return http.StatusBadRequest, fmt.Errorf("validation error: %s", response)
		}

		return http.StatusInternalServerError, fmt.Errorf("internal server error: %s", err)
	}

	ctx := context.Background()

	emailAlreadyExists, err := s.db.FindUserByEmail(ctx, user.Email)
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Warn(err.Error())
			return http.StatusInternalServerError, fmt.Errorf("Internal error")
		}
	}

	if emailAlreadyExists.Email != "" {
		return http.StatusConflict, fmt.Errorf("email already has been used")
	}

	logger.Info("inserting users")
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		logger.Warn(err.Error())
		return http.StatusInternalServerError, fmt.Errorf("Internal error")
	}
	_, err = s.db.InsertUsers(ctx,
		repository.InsertUsersParams{
			Name:     user.Name,
			Email:    user.Email,
			Password: hashedPassword,
		})

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (s *Service) Login(user *types.LoginUserPayload) (string, int, error) {
	logger.Info("Service.Login", "Searching user by email")
	u, err := s.db.FindUserByEmail(context.Background(), user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.LogError("Service.Login", fmt.Errorf("user not found: %s", user.Email))
			return "", http.StatusNotFound, fmt.Errorf("user not found")
		}
		logger.LogError("Service.Login", fmt.Errorf("error finding user: %w", err))
		return "", http.StatusInternalServerError, fmt.Errorf("error finding user: %w", err)
	}

	logger.Info("Service.Login", "User found, verifying password")
	if !auth.ComparePasswords(u.Password, []byte(user.Password)) {
		logger.LogError("Service.Login", fmt.Errorf("invalid password for user: %s", user.Email))
		return "", http.StatusUnauthorized, fmt.Errorf("invalid password")
	}

	logger.Info("Service.Login", "Password verified, generating token")
	secret := []byte(configs.Envs.JWT.JWTSecret)
	token, err := auth.CreateJWT(secret, u.ID)
	if err != nil {
		logger.LogError("Service.Login", fmt.Errorf("error creating token: %w", err))
		return "", http.StatusInternalServerError, fmt.Errorf("error creating token: %w", err)
	}

	logger.Info("Service.Login", "Token generated successfully")
	return token, http.StatusAccepted, nil
}
