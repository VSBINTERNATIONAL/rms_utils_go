package pkg

import (
	"encoding/json"
	"log"
	"os"
)

type db struct {
	DBUser     string `json:"dbUser"`
	DBPassword string `json:"dbPassword"`
	DBHost     string `json:"dbHost"`
	DBPort     string `json:"dbPort"`
	DBName     string `json:"dbName"`
}

type AppConfig struct {
	Database *db `json:"Database"`
}

func loadConfig(fileName string) (AppConfig, error) {
	var config AppConfig
	file, err := os.Open(fileName)
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func NewConfig(fileName string) (*AppConfig, error) {
	config := AppConfig{}
	config, err := loadConfig(fileName)
	if err != nil {
		log.Print("Erro ao carregar o arquivo de configuração", err)
		return nil, err
	}
	return &config, nil
}
