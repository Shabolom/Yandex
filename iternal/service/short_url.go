package service

import (
	"YandexPra/iternal/domain"
	models2 "YandexPra/iternal/models"
	"YandexPra/iternal/repository"
	"encoding/json"
)

type ShortUrl struct {
}

func NewShortUrl() *ShortUrl {
	return &ShortUrl{}
}

var urlRepo = repository.NewUrlRepo()

func (su *ShortUrl) Post(randUrl, url string) (string, error) {

	urlEntity := domain.Urls{
		Url:   url,
		Short: randUrl,
	}

	result, err := urlRepo.Post(urlEntity)

	if err != nil {
		return "", err
	}

	return result, err
}

func (su *ShortUrl) Get(key string) (domain.Urls, error) {

	result, err := urlRepo.Get(key)

	if err != nil {
		return domain.Urls{}, err
	}

	return result, err
}

func (su *ShortUrl) GetID() ([]models2.SaveStudent, error) {

	var models []models2.SaveStudent

	result, err := urlRepo.GetID()
	if err != nil {
		return []models2.SaveStudent{}, err
	}

	err = json.Unmarshal(result, &models)

	if err != nil {
		return []models2.SaveStudent{}, err
	}

	return models, nil
}

func (su *ShortUrl) GetUser() ([]models2.SaveUser, error) {

	var models []models2.SaveUser

	result, err := urlRepo.GetUser()

	err = json.Unmarshal(result, &models)

	if err != nil {
		return []models2.SaveUser{}, err
	}

	return models, nil

}

func (su *ShortUrl) PostShorten(url models2.ReqUrl, short string) (models2.ResUrl, error) {

	urlReqEntity := domain.Urls{
		Url:   url.Url,
		Short: short,
	}

	result, err := urlRepo.PostShorten(urlReqEntity)
	if err != nil {
		return models2.ResUrl{}, nil
	}

	urlRes := models2.ResUrl{Url: result}

	return urlRes, nil
}
