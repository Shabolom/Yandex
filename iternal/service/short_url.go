package service

import (
	"YandexPra/config"
	"YandexPra/iternal/domain"
	models2 "YandexPra/iternal/models"
	"YandexPra/iternal/repository"
	"YandexPra/iternal/tools"
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
		return result, err
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

func (su *ShortUrl) PostCsv(csv []models2.InfVidCsv) error {
	csvReqEntity := []domain.VideoInfo{}

	for i, _ := range csv {
		csvReqEntityPart := domain.VideoInfo{
			VideoID:             csv[i].VideoID,
			TrendingDate:        csv[i].TrendingDate,
			Title:               csv[i].Title,
			ChannelTitle:        csv[i].ChannelTitle,
			CategoryId:          csv[i].CategoryId,
			PublishTime:         csv[i].PublishTime,
			Tags:                csv[i].Tags,
			Likes:               csv[i].Likes,
			Dislikes:            csv[i].Dislikes,
			CommentCount:        csv[i].CommentCount,
			ThumbnailLink:       csv[i].ThumbnailLink,
			CommentsDisabled:    csv[i].CommentsDisabled,
			RatingsDisabled:     csv[i].RatingsDisabled,
			VideoErrorOrRemoved: csv[i].VideoErrorOrRemoved,
			Description:         csv[i].Description,
		}
		csvReqEntity = append(csvReqEntity, csvReqEntityPart)
	}

	err := urlRepo.PostCsv(csvReqEntity)
	if err != nil {
		return err
	}

	return nil
}

func (su *ShortUrl) PostBatch(urls []models2.ReqUrl) ([]domain.Urls, error) {
	var urlsReqEntity []domain.Urls

	for _, url := range urls {

		if err := tools.ValidUrl(url.Url); err != nil {
			return []domain.Urls{}, err
		}

		urlReqEntityPart := domain.Urls{
			Url:   url.Url,
			Short: config.Env.LocalApi + tools.Base62Encode(tools.RundUrl()),
		}
		urlsReqEntity = append(urlsReqEntity, urlReqEntityPart)
	}

	result, err := urlRepo.PostBatch(urlsReqEntity)
	if err != nil {
		return result, err
	}

	return result, nil
}
