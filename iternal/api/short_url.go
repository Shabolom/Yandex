package api

import (
	"YandexPra/iternal/service"
	"YandexPra/iternal/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type ShortUrl struct {
}

func NewShortUrl() *ShortUrl {
	return &ShortUrl{}
}

var urlService = service.NewShortUrl()

// Post заносим в базу данных url если он был не занесен, и возвращаем сокращеный
// url при обоих случаях (если url есть в базе то отдается он)
// @Summary  заносит в базу url в базу если его не было и выдает сокращеный url
// @Produce  json
// @Accept   json
// @Tags     shortUrl
// @Param    body  body      string    false  "User"
// @Success  201  {string}  string    "ok"
// @Failure  400  {object}  models.Error
// @Failure  404  {object}  models.Error
// @Router   / [post]
func (a *ShortUrl) Post(c *gin.Context) {

	shortUrl := tools.Base62Encode(tools.RundUrl())

	// c.Request.Body происходит получение тела от клиента поля Body который на даееый момент
	// является reader ом (это оболчка в которую оборачивается байты) и чтобы прочитать
	// контент который хранится в body нужно сначало распаковать reader и для этого мы
	// используем функцию io.ReadAll которая возвращает нам ошибку и байты которые мы потом
	// де сериализируем (преобразуем из байтового формата в объект)
	content, err := io.ReadAll(c.Request.Body)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	stringContent := string(content)
	defer c.Request.Body.Close()

	result, err := urlService.Post(shortUrl, stringContent)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.Writer.Header().Set("content-type", "text/plain")

	c.String(http.StatusCreated, result)
}

// Get производим переход (Redirect) по url принимая короткий урл как ключ
// в базе данных
// @Summary  переходим по ссылке которая хранится в базе данных используя короткий url как ключ
// @Produce  json
// @Accept   json
// @Tags     shortUrl
// @Success  307
// @Failure  400  {object}  models.Error
// @Router   /:key [get]
func (a *ShortUrl) Get(c *gin.Context) {

	id := c.Param("key")

	result, err := urlService.Get(id)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}
	fmt.Println("6")
	http.Redirect(c.Writer, c.Request, result.Url, http.StatusTemporaryRedirect)
}

// GetID производим авторизацию и получение данных student с другого сервера и возвращаем заполненное тело
// json формата
// @Summary  получаем массив объектов и отправляем клиенту
// @Produce  json
// @Accept   json
// @Tags     user
// @Success  200  {array}  models.SaveStudent
// @Failure  400  {object}  models.Error
// @Router   /get/ [get]
func (a *ShortUrl) GetID(c *gin.Context) {

	result, err := urlService.GetID()
	fmt.Println("qwer7")
	if err != nil {
		fmt.Println(err)
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}
	fmt.Println("qwer9")
	c.JSON(http.StatusOK, result)
}

// GetUsers производим получение данных user с другого сервера и возвращаем заполненное тело
// json формата
// @Summary  получаем массив объектов и отправляем клиенту
// @Produce  json
// @Accept   json
// @Tags     user
// @Success  200  {array}  models.SaveUser
// @Failure  400  {object}  models.Error
// @Router   /get/user [get]
func (a *ShortUrl) GetUsers(c *gin.Context) {

	result, err := urlService.GetUser()

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusOK, result)
}