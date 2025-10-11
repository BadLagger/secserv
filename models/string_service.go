package models

import "math/rand"

type StringService struct {
	/*YandexId           string
	YandexRedirectURI  string
	YandexClientSecret string*/
}

func NewStringService(id, uri, secret string) *StringService {
	return &StringService{
		/*YandexId:           id,
		YandexRedirectURI:  uri,
		YandexClientSecret: secret,*/
	}
}

/*func (s *StringService) GetOAuthStrs() (string, string) {
	return s.YandexId, s.YandexRedirectURI
}*/

func (s *StringService) GetWelcomeMsg() string {
	var msgs [4]string = [4]string{
		"Wow, Are You shure?",
		"Hello, You Are in the wrong place!",
		"Do you really want this?",
		"Are you shure that it's a good idea?",
	}

	return msgs[rand.Intn(4)]
}
