package service

import (
	"YandexPra/config"
	"YandexPra/iternal/domain"
	models2 "YandexPra/iternal/models"
	"YandexPra/iternal/repository"
	"YandexPra/iternal/tools"
	"encoding/json"
	"github.com/gofrs/uuid"
	"net/http"
)

type ShortUrl struct {
}

func NewShortUrl() *ShortUrl {
	return &ShortUrl{}
}

var urlRepo = repository.NewUrlRepo()

func (su *ShortUrl) Post(randUrl, url string) (string, error, int) {

	urlEntity := domain.Urls{
		Url:   url,
		Short: randUrl,
	}

	result, err, code := urlRepo.Post(urlEntity)
	if err != nil {
		return result, err, code
	}

	return result, err, code
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

func (su *ShortUrl) PostShorten(url models2.ReqUrl, short string) (models2.ResUrl, error, int) {
	id, _ := uuid.NewV4()

	urlReqEntity := domain.Urls{
		Url:   url.Url,
		Short: short,
	}
	urlReqEntity.ID = id

	result, err, code := urlRepo.PostShorten(urlReqEntity)
	if err != nil {
		return models2.ResUrl{}, nil, code
	}

	urlRes := models2.ResUrl{Url: result}

	return urlRes, nil, code
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

func (su *ShortUrl) PostBatch(urls []models2.ReqUrl, logUUID uuid.UUID) ([]domain.Urls, error, int) {
	var urlsReqEntity []domain.Urls

	for _, url := range urls {
		id, _ := uuid.NewV4()
		if err := tools.ValidUrl(url.Url); err != nil {
			return []domain.Urls{}, err, http.StatusBadRequest
		}
		urlReqEntityPart := domain.Urls{
			Url:   url.Url,
			Short: config.Env.LocalApi + tools.Base62Encode(tools.RundUrl()),
		}
		if logUUID != uuid.Nil {
			urlReqEntityPart.UserID = logUUID
		}
		urlReqEntityPart.ID = id
		urlsReqEntity = append(urlsReqEntity, urlReqEntityPart)
	}

	result, err, code := urlRepo.PostBatch(urlsReqEntity)
	if err != nil {
		return result, err, code
	}

	return result, nil, code
}

func (su *ShortUrl) Register(user models2.RegisterUsers) error {
	id, _ := uuid.NewV4()
	password, _ := tools.HashPassword(user.Password)

	regUsersEntity := domain.RegisterUsers{
		Login:    user.Login,
		Password: password,
	}
	regUsersEntity.ID = id

	err := urlRepo.Register(regUsersEntity)
	if err != nil {
		return err
	}
	return nil
}

func (su *ShortUrl) Login2(user models2.RegisterUsers) (error, domain.RegisterUsers) {

	regUsersEntity := domain.RegisterUsers{
		Login:    user.Login,
		Password: user.Password,
	}

	err, result := urlRepo.Auth(regUsersEntity.Password, regUsersEntity.Login)
	if err != nil {
		return err, domain.RegisterUsers{}
	}
	return nil, result
}

func (su *ShortUrl) GetUserUrls(id string) ([]domain.Urls, error) {

	result, err := urlRepo.GetUserUrls(id)

	if err != nil {
		return []domain.Urls{}, err
	}

	return result, nil

}
