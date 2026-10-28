package pkg

import (
	"os"

	"github.com/joho/godotenv"
)

func GetMongoDBURI() string {
	_ = godotenv.Load()
	return os.Getenv("MONGODB_URI")
}

func GetMongoDBDatabaseName() string {
	_ = godotenv.Load()
	return os.Getenv("MONGODB_DATABASE_NAME")
}
