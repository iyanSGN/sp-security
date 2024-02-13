package environment

import (
	"log"
	"os"
	"smartpatrol/pkg/file"

	"github.com/joho/godotenv"
)

func Init(envFileName string) {
	var err error

	os.Setenv("PROJECT_DIR", "serelo-backend")
	rootPath := file.GetRootDirectory()

	envFilePath := rootPath + envFileName

	err = godotenv.Load(envFilePath)
	if err != nil {
		log.Fatal("Errorrr loading .env file : ", err)
	}
}

func Get(key string) string {
	return os.Getenv(key)
}