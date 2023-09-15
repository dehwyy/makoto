package service

import (
	"errors"

	"github.com/dehwyy/Makoto/backend/auth/db/models"
	"github.com/dehwyy/Makoto/backend/auth/logger"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type signUpUserPayload struct {
	username string
	password string
	question string
	answer   string
}

type password_hasher interface {
	Hash(string) (string, error)
	Compare(string, string) bool
}

type CredentialsService struct {
	hasher password_hasher
	db     *gorm.DB
	l      logger.AppLogger
}

func NewCredentialsService(hasher password_hasher, db *gorm.DB, l logger.AppLogger) *CredentialsService {
	return &CredentialsService{
		hasher: hasher,
		db:     db,
		l:      l,
	}
}

func (s *CredentialsService) CreateUserPayload(username, password, question, answer string) signUpUserPayload {
	return signUpUserPayload{
		username: username,
		password: password,
		question: question,
		answer:   answer,
	}
}

func (s *CredentialsService) CreateUser(payload signUpUserPayload) (userId string, err error) {
	hashed_password, err := s.hasher.Hash(payload.password)
	if err != nil {
		return "", errors.New("error hashing password")
	}

	// by default, it should be stringified number which represents len(TotalUser) + 1; f.e. if you have 300 users, it should be 301 to be 100% unique
	unique_user_id := uuid.NewString()

	s.db.Create(&models.Credentials{
		UniqueUserId: unique_user_id,
		Username:     payload.username,
		Password:     hashed_password,
		Question:     payload.question,
		Answer:       payload.answer,
	})

	// TODO: save to db
	s.l.Infof("Created user: \n\t%v", payload, unique_user_id)
	s.l.Infof("HashedPassword is: %s", hashed_password)
	return unique_user_id, nil
}
