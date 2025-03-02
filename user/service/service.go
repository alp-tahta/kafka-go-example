package service

import (
	"encoding/json"
	"strconv"
	"user/messaging"
	"user/models"
)

type service struct {
	c messaging.MessagingClient
}

func NewService(c messaging.MessagingClient) *service {
	return &service{c: c}
}

func (s *service) CreateUser(user models.User) error {
	userInBtyes, err := json.Marshal(user)
	if err != nil {
		return err
	}

	ID := strconv.FormatInt(int64(user.ID), 10)

	err = s.c.Publish("user.created", messaging.Message{
		Key:   ID,
		Value: userInBtyes,
	})
	if err != nil {
		return err
	}

	return nil
}
