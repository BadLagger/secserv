package models

import "math/rand"

type StringService struct {
	YandexId          string
	YandexRedirectURI string
}

func NewStringService(id, uri string) *StringService {
	return &StringService{
		YandexId:          id,
		YandexRedirectURI: uri,
	}
}

func (s *StringService) GetOAuthStrs() (string, string) {
	return s.YandexId, s.YandexRedirectURI
}

func (s *StringService) GetWelcomeMsg() string {
	var msgs [4]string = [4]string{
		"Wow, Are You shure?",
		"Hello, You Are in the wrong place!",
		"Do you really want this?",
		"Are you shure that it's a good idea?",
	}

	return msgs[rand.Intn(4)]
}
