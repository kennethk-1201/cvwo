package helper

import (
	"log"
	"database/sql"
	"os"
	"github.com/spf13/viper"
)

func CheckErr(err error) {
	if err != nil {
        log.Fatalf("Error occured: %q", err)
	}
} 

// use viper package to read .env file
// return the value of the key
func GetEnvVariable(key string) string {

	// SetConfigFile explicitly defines the path, name and extension of the config file.
	// Viper will use this and not check any of the config paths.
	// .env - It will search for the .env file in the current directory
	viper.SetConfigFile(".env")

	// Find and read the config file
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	// viper.Get() returns an empty interface{}
	// to get the underlying type of the key,
	// we have to do the type assertion, we know the underlying value is string
	// if we type assert to other type it will throw an error
	value, ok := viper.Get(key).(string)

	// If the type is a string then ok will be true
	// ok will make sure the program not break
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}

func SetupDB() *sql.DB {
	os.Setenv("DATABASE_URL", GetEnvVariable("DATABASE_URL")) // remove in production
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	CheckErr(err)

    return db
}