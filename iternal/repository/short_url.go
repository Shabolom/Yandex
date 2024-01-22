package repository

import (
	"YandexPra/config"
	"YandexPra/iternal/domain"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type UrlRepo struct {
}

func NewUrlRepo() *UrlRepo {
	return &UrlRepo{}
}

func (ur *UrlRepo) Post(url domain.Urls) (string, error) {

	if result, err := ur.Get(url.Url); err == nil {
		fmt.Println("такой url уже есть в базе")
		return "http://localhost:8080/" + result.Short, nil
	}

	// метод библиотеки для сохранения сущности в базе данных
	err := config.DB.Create(&url).Error

	if err != nil {
		return "", err
	}
	return "http://localhost:8080/" + url.Short, nil
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

	//data := []byte(`{"login": "admin", "password": "admin"}`)
	data := `{"login": "admin", "password": "admin"}`
	response, err := http.Post("http://localhost:8080/api/user/login", "application/json", strings.NewReader(data))
	//response, err := http.Post("http://localhost:8000/api/user/login", "application/json", bytes.NewBuffer(data))
	// если что смотреть тут https://practicum.yandex.ru/learn/go-advanced/courses/fe725b51-a888-4c0a-809a-611a1ef8e2ba/sprints/161327/topics/6582b0e4-a277-4b6e-b492-a9b98681d530/lessons/a535ccde-9231-48a6-9ac2-ca7b8149c8bf/

	if err != nil {
		return nil, err
	}

	if response.Header.Get("Token") != "" {

		req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/api/student", nil)
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

	req, err := http.NewRequest(http.MethodGet, "https://jsonplaceholder.typicode.com/users", nil)
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
