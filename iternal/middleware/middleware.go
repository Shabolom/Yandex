package middleware

import (
	"YandexPra/config"
	"YandexPra/iternal/tools"
	"compress/gzip"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	log "github.com/sirupsen/logrus"
)

type JWTclaims struct {
	jwt.RegisteredClaims
	Login    string
	Password string
	ID       uuid.UUID
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// продолжаем работать с хэндлером который идет после мидлвейра  (который был вызван изначально))
		c.Next()

		latency := time.Since(t)
		log.WithField("component", "latency").Info(latency)
	}
}

func Gzip() gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.Request.Header.Get("Content-Encoding") == `gzip` {
			reader, err := gzip.NewReader(c.Request.Body)

			if err != nil {
				log.WithField("component", "Gzip").Warn(err)
				http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
				return
			}
			c.Request.Body = reader
			defer reader.Close()
		}
	}
}

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.Request.Header.Get("Authorization") == `` {
			//c.Writer.WriteHeader(http.StatusUnauthorized)
			//_, err := c.Writer.Write([]byte("You're Unauthorized!"))
			//if err != nil {
			//	fmt.Println(err)
			//	return
			//}
			tools.CreateError(http.StatusUnauthorized, errors.New("You're Unauthorized"), c)
			return
		}

		strToken := c.Request.Header.Get("Authorization")
		fmt.Println(strToken, "qwe")
		token, err := jwt.Parse(strToken,
			func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					tools.CreateError(http.StatusUnauthorized, errors.New("You're Unauthorized"), c)
					//c.Writer.WriteHeader(http.StatusUnauthorized)
					//_, err := c.Writer.Write([]byte("You're Unauthorized!"))
					//return nil, err
				}
				return []byte(config.Env.SecretKey), nil
			})

		if err != nil {
			tools.CreateError(http.StatusUnauthorized, errors.New("You're Unauthorized"), c)
			return
			//c.Writer.WriteHeader(http.StatusUnauthorized)
			//_, err = c.Writer.Write([]byte("You're Unauthorized due to error parsing the JWT!"))
			//if err != nil {
			//	fmt.Println("ошибка при записи в тело ответа")
			//	return
			//}
		}

		if !token.Valid {
			tools.CreateError(http.StatusUnauthorized, errors.New("You're Unauthorized"), c)
			return
			//c.Writer.WriteHeader(http.StatusUnauthorized)
			//_, err = c.Writer.Write([]byte("You're Unauthorized due to No token in the header"))
			//if err != nil {
			//	fmt.Println("ошибка при записи в тело ответа")
			//	return
			//}
		} else {
			c.Next()
		}
	}
}
