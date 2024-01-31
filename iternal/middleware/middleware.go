package middleware

import (
	"compress/gzip"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

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
