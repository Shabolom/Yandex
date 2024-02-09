package api

import (
	"YandexPra/config"
	"YandexPra/iternal/models"
	"YandexPra/iternal/service"
	"YandexPra/iternal/tools"
	"encoding/json"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v4"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
	log "github.com/sirupsen/logrus"
)

type ShortUrl struct{}

func NewShortUrl() *ShortUrl {
	return &ShortUrl{}
}

var urlService = service.NewShortUrl()

// Post заносим в базу данных url если он был не занесен, и возвращаем сокращеный
// url при обоих случаях (если url есть в базе то отдается он)
//
//	@Summary	заносит в базу url в базу если его не было и выдает сокращеный url
//	@Produce	json
//	@Accept		json
//	@Tags		shortUrl
//	@Param		body	body		string	false	"User"
//	@Success	201		{string}	string	"ok"
//	@Failure	400		{object}	models.Error
//	@Failure	404		{object}	models.Error
//	@Router		/ [post]
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
		log.WithField("component", "rest").Warn(err)
		return
	}

	stringContent := string(content)

	if err = tools.ValidUrl(stringContent); err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	stringContent = strings.Replace(stringContent, `"`, "", -1)

	defer c.Request.Body.Close()

	result, err, code := urlService.Post(shortUrl, stringContent)

	if err != nil {
		tools.CreateError(code, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	c.String(code, result)
}

// Get производим переход (Redirect) по url принимая короткий урл как ключ
// в базе данных
//
//	@Summary	переходим по ссылке которая хранится в базе данных используя короткий url как ключ
//	@Produce	json
//	@Accept		json
//	@Tags		shortUrl
//	@Success	307
//	@Failure	400	{object} 	models.Error
//	@Router		/:key [get]
func (a *ShortUrl) Get(c *gin.Context) {
	key := c.Param("key?")

	if key == "" {
		data, err := io.ReadAll(c.Request.Body)

		if err != nil {
			tools.CreateError(http.StatusBadRequest, err, c)
			log.WithField("component", "rest").Warn(err)
			return
		}

		strData := string(data)

		result, err := urlService.Get(strData)

		if err != nil {
			tools.CreateError(http.StatusBadRequest, err, c)
			log.WithField("component", "rest").Warn(err)
			return
		}

		http.Redirect(c.Writer, c.Request, result.Url, http.StatusTemporaryRedirect)
	}

	key = config.Env.LocalApi + key
	result, err := urlService.Get(key)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	http.Redirect(c.Writer, c.Request, result.Url, http.StatusTemporaryRedirect)
}

// GetID производим авторизацию и получение данных student с другого сервера и возвращаем заполненное тело
// json формата
//
//	@Summary	получаем массив объектов и отправляем клиенту
//	@Produce	json
//	@Accept		json
//	@Tags		user
//	@Success	200	{array}		models.SaveStudent
//	@Failure	400	{object}	models.Error
//	@Router		/get/ [get]
func (a *ShortUrl) GetID(c *gin.Context) {

	result, err := urlService.GetID()

	if err != nil {
		fmt.Println(err)
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetUsers производим получение данных user с другого сервера и возвращаем заполненное тело
// json формата
//
//	@Summary	получаем массив объектов и отправляем клиенту
//	@Produce	json
//	@Accept		json
//	@Tags		user
//	@Success	200	{array}		models.SaveUser
//	@Failure	400	{object}	models.Error
//	@Router		/get/user [get]
func (a *ShortUrl) GetUsers(c *gin.Context) {

	result, err := urlService.GetUser()

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetShorten принимаем в теле урл и возвращаем сокращеный
//
//	@Summary	полуяаем сокращеный урл
//	@Produce	json
//	@Accept		json
//	@Tags		url
//	@Param		body	body		models.ReqUrl	false	"User"
//	@Success	200	{array}		models.ResUrl
//	@Failure	400	{object}	models.Error
//	@Router		/api/shorten [post]
func (a *ShortUrl) GetShorten(c *gin.Context) {
	var reqUrl models.ReqUrl

	shortUrl := tools.Base62Encode(tools.RundUrl())

	body, err := io.ReadAll(c.Request.Body)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "ReadAll").Warn(err)
		return
	}

	err = json.Unmarshal(body, &reqUrl)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "Unmarshal").Warn(err)
		return
	}

	result, err, code := urlService.PostShorten(reqUrl, shortUrl)

	if err != nil {
		tools.CreateError(code, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// PostCsv принимаем csv file в форме и заполняем базу его содержимым
//
//	@Summary	принимаем csv file в форме и заполняем базу его содержимым
//	@Produce	json
//	@Accept		mpfd
//	@Tags		url
//	@Param		mpfd	formData	file	false	"Body with CSV file"
//	@Success	200		{string}	string	"успешно заполнили базу"
//	@Failure	400		{object}	models.Error
//	@Router		/api/csv [post]
func (a *ShortUrl) PostCsv(c *gin.Context) {
	var csvs []models.InfVidCsv

	file, _, err := c.Request.FormFile("csv")
	defer file.Close()

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	err = gocsv.UnmarshalMultipartFile(&file, &csvs)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}
	fmt.Println(csvs[0])
	err = urlService.PostCsv(csvs)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	c.String(http.StatusOK, "успешно заполнили базу")
}

// PostBatch принимаем массив джейсонов с контентом в виде ссылок и возвращаем их сокращенную версию
//
//	@Summary	принимаем массив джейсонов с контентом в виде ссылок и возвращаем их сокращенную версию (все делаем
//	с транзакцией те пока не запишется и не пройдет проверку вя полученная информация ничего не заполнится в базе данных)
//	@Produce	json
//	@Accept		json
//	@Tags		url
//	@Param		body	body		[]models.ReqUrl	false	"User"
//	@Success	200		{string}	string	"успешно заполнили базу"
//	@Failure	400		{object}	models.Error
//	@Router		/api/shorten/batch [post]
func (a *ShortUrl) PostBatch(c *gin.Context) {
	var urls []models.ReqUrl
	claim := &tools.Claims{}
	logUUID := uuid.UUID{}

	if c.Request.Header.Get("Authorization") != "" {
		strToken := c.Request.Header.Get("Authorization")
		token, err := jwt.ParseWithClaims(strToken, claim, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(config.Env.SecretKey), nil
		})
		if err != nil {
			tools.CreateError(http.StatusBadRequest, err, c)
			log.WithField("component", "ReadAll").Warn(err)
			return
		}
		if token.Valid {
			logUUID = claim.UserID
		}
	}

	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	err = json.Unmarshal(data, &urls)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	result, err, code := urlService.PostBatch(urls, logUUID)

	if err != nil {
		tools.CreateError(code, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	c.JSON(code, result)
}

func (a *ShortUrl) Register(c *gin.Context) {
	var regUsers models.RegisterUsers

	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	err = json.Unmarshal(data, &regUsers)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	err = urlService.Register(regUsers)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	c.String(http.StatusOK, "успешно зарегестрировались")
}

func (a *ShortUrl) Login2(c *gin.Context) {
	var regUsers models.RegisterUsers

	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	err = json.Unmarshal(data, &regUsers)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	err, result := urlService.Login2(regUsers)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tools.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			// ниже описание части полей и за что они отвечают
			Issuer:    "Timur",                                           // Указывает, кто создал и подписал JWT.
			Subject:   "Authorization",                                   // Определяет субъект (тему), к которой относится JWT.
			IssuedAt:  jwt.NewNumericDate(time.Now()),                    // Указывает время создания JWT.
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)), // Указывает, кто создал и подписал JWT.
		},
		// собственное утверждение
		UserID: result.ID,
		Login:  regUsers.Login,
	})

	tokenString, err := token.SignedString([]byte(config.Env.SecretKey))

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	cookie := http.Cookie{
		Name:  "Token",     //Имя cookie.
		Value: tokenString, //Значение cookie.
		Path:  "/",         //Путь на сервере, для которого cookie действителен. Если установлен "/",
		// то cookie действителен для всего домена.
		Expires: time.Now().Add(time.Hour * 2), // Время истечения срока действия cookie.
		// Если не указано, то cookie действителен только в текущей сессии.

		//MaxAge:     0,// Продолжительность срока действия cookie в секундах.
		//Secure:     false,// Указывает, что cookie должен быть отправлен только по защищенному (HTTPS) соединению.
		//HttpOnly:   false,// Указывает, что cookie должен быть доступен только через HTTP-запросы, а не JavaScript.
		//SameSite:   0,
		//Raw:        "",
		//Unparsed:   nil,
	}
	c.Writer.Header().Set("token", tokenString)
	http.SetCookie(c.Writer, &cookie)

	c.String(http.StatusOK, "успешно авторизировались")
}

func (a *ShortUrl) GetUrls(c *gin.Context) {
	key := c.Param("userid")

	result, err := urlService.GetUserUrls(key)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	c.JSON(http.StatusOK, result)
}
