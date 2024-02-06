package main

import (
	"YandexPra/config"
	_ "YandexPra/docs"
	migrate "YandexPra/init"
	"YandexPra/iternal/routes"
	"YandexPra/iternal/tools"

	log "github.com/sirupsen/logrus"
)

func main() {
	//	@title		cmd Service
	//	@version	1.0.0
	//	@host		localhost:8000

	config.CheckFlagEnv()
	tools.InitLogger()

	// config.InitPgSQL инициализируем подключение к базе данных
	err := config.InitPgSQL()
	if err != nil {
		log.WithField("component", "initialization").Panic(err)
	}

	// вызываем миграцию структуры в базу данных
	migrate.Migrate()

	//test.ClientGet()
	//test.Redirect()

	// конфигурация (инициализация) end point или ручка (можно назвать имя запроса)
	// (как api student) URLов пример (localhost, 8080) конфигурация всех URLов которые будет
	// обрабатывать сервер

	r := routes.SetupRouter()

	// запуск сервера
	if err = r.Run(config.Env.Host + ":" + config.Env.Port); err != nil {
		log.WithField("component", "run").Panic(err)
	}

}
