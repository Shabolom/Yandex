package main

import (
	"YandexPra/config"
	_ "YandexPra/docs"
	migrate "YandexPra/init"
	"YandexPra/iternal/routes"
)

func main() {
	// @title    cmd Service
	// @version  1.0.0
	// @host     localhost:8080

	// сканируем env файл
	//_ = godotenv.Load(".env")
	config.CheckFlagEnv()

	// config.InitPgSQL инициализируем подключение к базе данных
	err := config.InitPgSQL()
	if err != nil {
		panic(err)
	}

	// вызываем миграцию труктуры в базу данных
	migrate.Migrate()

	//test.ClientGet()
	//test.Redirect()

	// конфигурация (инициализация) end point или ручка (можно назвать имя запроса)
	// (как api student) URLов пример (localhost, 8080) конфигурация всех URLов которые будет
	// обрабатывать сервер

	r := routes.SetupRouter()

	// запуск сервера
	if err := r.Run(config.Env.DbHost + ":" + config.Env.Port); err != nil {
		panic(err)
	}

}
