package config

import "os"

type AppConfig struct {
	DBUSERNAME string
	DBPASSWORD string
	DBHOST     string
	DBPORT     string
	DBNAME     string
	DBARGS     string
}

var Secret string

func IniConfig() *AppConfig {
	var cnf = readconfig()
	if cnf == nil {
		return nil
	}

	return cnf
}

func readconfig() *AppConfig {
	var result = new(AppConfig)

	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatalln("cannot load env file for database")
	// }

	result.DBUSERNAME = os.Getenv("DBUSERNAME")
	result.DBPASSWORD = os.Getenv("dbPASSWORD")
	result.DBHOST = os.Getenv("DBHOST")
	result.DBPORT = os.Getenv("DBPORT")
	result.DBNAME = os.Getenv("DBNAME")
	result.DBARGS = os.Getenv("DBARGS")
	Secret = os.Getenv(("SECRET"))

	return result
}
