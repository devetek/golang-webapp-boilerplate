package member

import (
	"context"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UseCase struct {
	DB         *gorm.DB
	Log        *logrus.Logger
	Validate   *validator.Validate
	Repository *Repository
}

func NewUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, repository *Repository) *UseCase {
	return &UseCase{
		DB:         db,
		Log:        logger,
		Validate:   validate,
		Repository: repository,
	}
}

func (c *UseCase) Create(ctx context.Context, request *RegisterRequest) (*Response, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, err
	}

	total, err := c.Repository.CountByUsername(tx, request.Username)
	if err != nil {
		c.Log.Warnf("Failed count user from database : %+v", err)
		return nil, err
	}

	if total > 0 {
		c.Log.Warnf("User already exists : %+v", err)
		return nil, errors.New("user already exists")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Log.Warnf("Failed to generate bcrype hash : %+v", err)
		return nil, err
	}

	// new user
	user := &Entity{
		Password: string(password),
		Username: request.Username,
	}

	if err := c.Repository.Create(tx, user); err != nil {
		c.Log.Warnf("Failed create user to database : %+v", err)
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, err
	}

	return UserToResponse(user), nil
}
