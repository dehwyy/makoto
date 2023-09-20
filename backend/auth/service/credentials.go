package service

import (
	"errors"
	"strings"
	"sync"

	"github.com/dehwyy/Makoto/backend/auth/db/models"
	"github.com/dehwyy/Makoto/backend/auth/dto"
	"github.com/dehwyy/Makoto/backend/auth/logger"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

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

func (s *CredentialsService) schema() *gorm.DB {
	return s.db.Model(&models.Credentials{})
}

// Service calls

// ! C - create
func (s *CredentialsService) CreateUser(payload dto.CreateUser) (userId string, err error) {
	// as we hash 2 times, to improve performance, I decided to make it in goroutines
	var hashed_password, hashed_answer string
	var err_password, err_answer error
	var wg sync.WaitGroup

	wg.Add(2)

	// hashing password
	go func() {
		hashed_password, err = s.hasher.Hash(payload.Password)
		if err != nil {
			err_password = errors.New("error hashing password")
		}
		wg.Done()

	}()

	// hashing answer
	go func() {
		hashed_answer, err = s.hasher.Hash(strings.ToLower(payload.Answer))
		if err != nil {
			err_answer = errors.New("error hashing answer")
		}
		wg.Done()
	}()

	wg.Wait()

	if err_password != nil {
		return "", err_password

	} else if err_answer != nil {
		return "", err_answer

	}

	// by default, it should be stringified number which represents len(TotalUser) + 1; f.e. if you have 300 users, it should be 301 to be 100% unique
	unique_user_id := uuid.NewString()

	res := s.db.Create(&models.Credentials{
		UniqueUserId: unique_user_id, // generated (on create, could be changed later)
		Username:     payload.Username,
		Password:     hashed_password, // hashed
		Question:     payload.Question,
		Answer:       hashed_answer, // hashed
	})
	if res.Error != nil {
		return "", res.Error
	}

	s.l.Infof("Created user: \n\t%v", payload, unique_user_id)
	s.l.Infof("HashedPassword is: %s", hashed_password)
	return unique_user_id, nil
}

//! V - validate

func (s *CredentialsService) ValidateUser(username, password string) (userId string, err error) {
	// getting user by provided userId and select
	user := new(struct {
		UniqueUserId string
		Password     string
	})

	result := s.schema().Where("username = ?", username).Select("password", "unique_user_id").Find(user)
	if result.Error != nil {
		return "404", result.Error
	}

	if !s.hasher.Compare(password, user.Password) {
		return "403", errors.New("wrong password")
	}

	s.l.Debugf("Found user: %v", *user)

	return user.UniqueUserId, nil
}

func (s *CredentialsService) ValidateUserPassword(userId, password string) (err error) {
	user := new(struct {
		Password string
	})

	result := s.schema().Where("unique_user_id = ?", userId).Select("password").Find(user)
	if result.Error != nil {
		return result.Error
	}

	if !s.hasher.Compare(password, user.Password) {
		return errors.New("wrong password")
	}

	return nil
}

func (s *CredentialsService) ValidateUserAnswer(userId, answer string) (err error) {
	user := new(struct {
		Answer string
	})

	result := s.schema().Where("unique_user_id = ?", userId).Select("answer").Find(user)
	if result.Error != nil {
		return result.Error
	}

	if !s.hasher.Compare(strings.ToLower(answer), user.Answer) {
		return errors.New("wrong answer")
	}

	return nil
}

//! G - get

func (s *CredentialsService) GetQuestion(userId string) (question string, err error) {
	user := new(struct {
		Question string
	})

	result := s.schema().Where("unique_user_id = ?", userId).Select("question").Find(user)
	if result.Error != nil {
		return "", result.Error
	}

	return user.Question, nil
}

func (s *CredentialsService) GetUserById(userId string) (user *dto.User, err error) {
	user = new(dto.User)

	result := s.schema().Where("unique_user_id = ?", userId).Select("username").Find(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

//! U - update

func (s *CredentialsService) UpdatePassword(userId, new_password string) error {
	hashed_password, err := s.hasher.Hash(new_password)
	if err != nil {
		return err
	}

	res := s.schema().Where("unique_user_id = ?", userId).Update("password", hashed_password)

	return res.Error
}

//! R - remove

func (s *CredentialsService) RemoveToken(userId string) error {
	res := s.schema().Delete(&models.Credentials{}, "unique_user_id = ?", userId)

	return res.Error
}
