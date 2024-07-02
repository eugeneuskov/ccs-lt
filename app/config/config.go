package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

const (
	ccsUrl         = "CCS_URL"
	ccsClientToken = "CCS_CLIENT_TOKEN"

	rps = "RPS"
)

type Config struct {
	CcsUrl         string
	CcsClientToken string
	Rps            int
}

func (c *Config) Init() (*Config, error) {
	ccsUrlValue, err := envValue(ccsUrl)
	if err != nil {
		return nil, err
	}

	ccsClientTokenValue, err := envValue(ccsClientToken)
	if err != nil {
		return nil, err
	}

	rpsValueString, err := envValue(rps)
	if err != nil {
		return nil, err
	}
	rpsValue, err := strconv.Atoi(rpsValueString)
	if err != nil {
		return nil, err
	}

	return &Config{
		CcsUrl:         ccsUrlValue,
		CcsClientToken: ccsClientTokenValue,
		Rps:            rpsValue,
	}, nil
}

func envValue(envKey string) (string, error) {
	value, found := os.LookupEnv(envKey)
	if !found && value != "" {
		return value, nil
	}

	reader, err := os.Open(".env")
	if err != nil {
		return "", err
	}
	defer func(reader *os.File) {
		_ = reader.Close()
	}(reader)

	env, err := godotenv.Parse(reader)
	value, found = env[envKey]
	if !found {
		return "", fmt.Errorf("%s not found", envKey)
	}

	if value == "" {
		return "", fmt.Errorf("%s is empty", envKey)
	}

	return value, nil
}
