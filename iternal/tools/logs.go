package tools

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	"os"

	"YandexPra/config"

	log "github.com/sirupsen/logrus"
)

func InitLogger() error {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}

	defer file.Close()

	log.SetReportCaller(true)

	if config.Env.Production {
		log.SetLevel(log.WarnLevel)
		log.SetOutput(file)
		log.SetFormatter(&log.JSONFormatter{})

	} else {
		log.SetLevel(log.DebugLevel)
		log.SetOutput(os.Stdout)
		log.SetFormatter(&nested.Formatter{
			ShowFullLevel: true,
			HideKeys:      true,
			FieldsOrder:   []string{"component", "category"},
		})
	}
	return nil
}
