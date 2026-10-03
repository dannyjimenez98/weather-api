package env

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Load() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("error loading .env: %v\n", err)
	}
}

func APIKey() string {
	return os.Getenv("API_KEY")
}

func Address() string {
	return os.Getenv("ADDRESS")
}

func Password() string {
	return os.Getenv("PASSWORD")
}

func Database() int {
	db, err := strconv.Atoi(os.Getenv("DATABASE"))
	if err != nil {
		log.Fatalln("cannot read database from env")
	}

	return db
}
