package repository

import (
	"YandexPra/config"
	"YandexPra/iternal/domain"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	parseUrl "net/url"
	"os"
	"strings"
)

type UrlRepo struct {
}

func NewUrlRepo() *UrlRepo {
	return &UrlRepo{}
}

func (ur *UrlRepo) Post(url domain.Urls) (string, error) {

	_, err := parseUrl.ParseRequestURI(url.Url)
	if err != nil {
		return "", errors.New("не валидный урл " + url.Url)
	}

	if result, err := ur.Get(url.Url); err == nil {
		fmt.Println("такой url уже есть в базе")
		return config.Env.LocalApi + result.Short, nil
	}

	err = ur.saveFile(url)
	if err != nil {
		return "", err
	}

	// метод библиотеки для сохранения сущности в базе данных
	err = config.DB.Create(&url).Error
	if err != nil {
		return "", err
	}

	return config.Env.LocalApi + url.Short, nil
}

func (ur *UrlRepo) Get(key string) (domain.Urls, error) {
	var url domain.Urls

	// метод библиотеки для сохранения сущности в базе данных
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

	req, err := http.NewRequest(http.MethodGet, config.Env.JsonApi, nil)
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

func (ur *UrlRepo) PostShorten(model domain.Urls) (string, error) {

	result, err := ur.Post(model)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (ur *UrlRepo) saveFile(url domain.Urls) error {

	file, err := os.OpenFile("url.save", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
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
