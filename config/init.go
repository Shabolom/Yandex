package config

import (
	"flag"
	"github.com/joho/godotenv"
	"os"
)

// env Структура для хранения переменных среды
type env struct {
	Host       string
	Port       string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
}

// Env глобальная переменная для доступа к переменным среды
var Env env

// CheckFlagEnv Метод проверяющий флаги
func CheckFlagEnv() {
	var host string
	var port string
	var dbHost string
	var dbPort string
	var dbUser string
	var dbPassword string
	var dbName string

	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	var flagHost = flag.String("h", "", "host")
	var flagPort = flag.String("p", "", "port")
	var flagDbHost = flag.String("dh", "", "dbHost")
	var flagDbPort = flag.String("dp", "", "dbPort")
	var flagDbUser = flag.String("du", "", "dbUser")
	var flagDbPassword = flag.String("dpa", "", "dbPassword")
	var flagDbName = flag.String("dn", "", "dbName")

	flag.Parse()

	if os.Getenv("HOST") != "" {
		host = os.Getenv("HOST")
	} else {
		host = "localhost"
	}

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	} else {
		port = "8080"
	}

	if os.Getenv("DB_HOST") != "" {
		dbHost = os.Getenv("DB_HOST")
	} else {
		dbHost = ""
	}

	if os.Getenv("DB_PORT") != "" {
		dbPort = os.Getenv("DB_PORT")
	} else {
		dbPort = ""
	}

	if os.Getenv("DB_USER") != "" {
		dbUser = os.Getenv("DB_USER")
	} else {
		dbUser = ""
	}

	if os.Getenv("DB_PASSWORD") != "" {
		dbPassword = os.Getenv("DB_PASSWORD")
	} else {
		dbPassword = ""
	}

	if os.Getenv("DB_NAME") != "" {
		dbName = os.Getenv("DB_NAME")
	} else {
		dbName = ""
	}

	if *flagHost != "" {
		host = *flagHost
	}

	if *flagPort != "" {
		port = *flagPort
	}

	if *flagDbHost != "" {
		dbHost = *flagDbHost
	}

	if *flagDbPort != "" {
		dbPort = *flagDbPort
	}

	if *flagDbUser != "" {
		dbUser = *flagDbUser
	}

	if *flagDbPassword != "" {
		dbPassword = *flagDbPassword
	}

	if *flagDbName != "" {
		dbName = *flagDbName
	}

	Env = env{
		Host:       host,
		Port:       port,
		DbHost:     dbHost,
		DbPort:     dbPort,
		DbUser:     dbUser,
		DbPassword: dbPassword,
		DbName:     dbName,
	}
}
