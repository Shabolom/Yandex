package repository

import (
	"YandexPra/config"
	"YandexPra/iternal/domain"
	"YandexPra/iternal/tools"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"strings"
)

type UrlRepo struct {
}

func NewUrlRepo() *UrlRepo {
	return &UrlRepo{}
}

func (ur *UrlRepo) Post(url domain.Urls) (string, error, int) {

	if result, err := ur.Get(url.Url); err == nil {
		return result.Short, nil, http.StatusConflict
	}

	err := ur.saveFile(url)
	if err != nil {
		return "", err, http.StatusBadRequest
	}

	// метод библиотеки для сохранения сущности в базе данных
	err = config.DB.Create(&url).Error
	if err != nil {
		return "", err, http.StatusBadRequest
	}

	return config.Env.LocalApi + url.Short, nil, http.StatusCreated
}

func (ur *UrlRepo) Get(key string) (domain.Urls, error) {
	var url domain.Urls

	err := config.DB.
		Where("short = ?", key).
		Or("url = ?", key).
		First(&url).
		Error

	if err != nil {
		return domain.Urls{}, err
	}

	return url, nil
}

func (ur *UrlRepo) GetID() ([]byte, error) {
	c := config.Env

	//data := []byte(`{"login": "admin", "password": "admin"}`)
	data := `{"login": "admin", "password": "admin"}`
	response, err := http.Post(c.ConnectionApi, "application/json", strings.NewReader(data))
	//response, err := http.Post("http://localhost:8000/api/user/login", "application/json", bytes.NewBuffer(data))
	// если что смотреть тут https://practicum.yandex.ru/learn/go-advanced/courses/fe725b51-a888-4c0a-809a-611a1ef8e2ba/sprints/161327/topics/6582b0e4-a277-4b6e-b492-a9b98681d530/lessons/a535ccde-9231-48a6-9ac2-ca7b8149c8bf/

	if err != nil {
		return nil, err
	}

	if response.Header.Get("Token") != "" {
		req, err := http.NewRequest(http.MethodGet, c.ConnectionGet, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Authorization", "Bearer "+response.Header.Get("Token"))

		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode != http.StatusOK {
			return nil, errors.New("сервер не отвечает")
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()

		return body, nil
	}
	return nil, errors.New("ошибка в response.Header.Get(\"Token\") != \"\"")
}

func (ur *UrlRepo) GetUser() ([]byte, error) {
	c := config.Env

	req, err := http.NewRequest(http.MethodGet, c.JsonApi, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return body, nil
}

func (ur *UrlRepo) PostShorten(model domain.Urls) (string, error, int) {

	result, err, code := ur.Post(model)
	if err != nil {
		return "", err, code
	}

	return result, nil, code
}

func (ur *UrlRepo) PostCsv(masDomainCsv []domain.VideoInfo) error {

	go func([]domain.VideoInfo) {
		tx := config.DB.Begin()
		for _, csv := range masDomainCsv {
			err := tx.
				Create(&csv).
				Error
			if err != nil {
				tx.Rollback()
				panic(err)
			}
		}
		tx.Commit()
	}(masDomainCsv)

	return nil
}

func (ur *UrlRepo) PostBatch(urls []domain.Urls) ([]domain.Urls, error, int) {
	var res []domain.Urls
	tx := config.DB.Begin()
	var routErr bool

	for _, url := range urls {
		result, err := ur.Get(url.Url)
		if err == nil {
			res = append(res, result)
			routErr = true
		}
	}

	if routErr == true {
		return res, nil, http.StatusConflict
	}

	for _, url := range urls {
		err := tx.
			Create(&url).
			Error
		if err != nil {
			tx.Rollback()
		}
	}
	tx.Commit()

	return urls, nil, http.StatusCreated
}

func (ur *UrlRepo) Register(user domain.RegisterUsers) error {
	var users domain.RegisterUsers

	if err := config.DB.
		Find(&users, "login = ?", user.Login).
		Error; err == nil {
		return errors.New("пользователь уже есть в базе данных")
	}

	err := config.DB.
		Create(&user).
		Error
	if err != nil {
		return err
	}

	return nil
}

func (ur *UrlRepo) Auth(password, login string) (error, domain.RegisterUsers) {
	var user domain.RegisterUsers

	err := config.DB.
		Where("login= ?", login).
		First(&user).
		Error

	if err != nil {
		return errors.New("не верный логин или пароль"), domain.RegisterUsers{}
	}

	if !tools.CheckPasswordHash(password, user.Password) {
		return errors.New("не верный логин или пароль"), domain.RegisterUsers{}
	}

	return nil, user
}

func (ur *UrlRepo) GetUserUrls(id string) ([]domain.Urls, error) {
	var urls []domain.Urls
	fmt.Println(id, "qweqweqwe")
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	err = config.DB.Where("user_id=?", parsedUUID).
		Find(&urls).
		Error
	if err != nil {
		return []domain.Urls{}, err
	}
	return urls, nil
}

func (ur *UrlRepo) saveFile(url domain.Urls) error {

	file, err := os.OpenFile("url.save", os.O_CREATE|os.O_APPEND|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}

	data, err := json.Marshal(url)
	if err != nil {
		return err
	}

	_, err = file.WriteString("\n")
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	defer file.Close()
	return nil
}
