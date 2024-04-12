package config

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"

	"github.com/joho/godotenv"
)

func Load() {
	localEnv := os.Getenv("LOCAL")

	if localEnv == "true" {
		envFilePath := fmt.Sprintf("%s/local.env", getConfigDir())
		err := godotenv.Load(envFilePath)
		if err != nil {
			panic(
				fmt.Sprintf("Panic error on load env. Err: %v", err),
			)
		}
	}
}

func GetConfig(key string) string {
	return os.Getenv(key)
}

func GetConfigInt(key string, defaultValue int) int {
	configValue := os.Getenv(key)
	configInt, err := strconv.Atoi(configValue)
	if err != nil {
		return defaultValue
	}
	return configInt
}

func getConfigDir() string {
	_, currentFilePath, _, _ := runtime.Caller(0)
	return path.Join(path.Dir(currentFilePath), "")
}
