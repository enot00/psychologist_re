package config

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type configuration struct {
	DatabaseName        string
	DatabaseHost        string
	DatabaseUser        string
	DatabasePassword    string
	FileStorageLocation string
	JWTSecretKey        string
	JWTTokenTTL         time.Duration
	Salt                string
	ConfigPath          string
}

func GetConfiguration() *configuration {
	ttl, err := strconv.Atoi(os.Getenv("JWT_TTL"))
	if err != nil {
		log.Fatalf("unsopportable TTL type: %q\n", err)
	}

	return &configuration{
		DatabaseName:        os.Getenv("DB_NAME"),
		DatabaseHost:        os.Getenv("DB_HOST"),
		DatabaseUser:        os.Getenv("DB_USER"),
		DatabasePassword:    os.Getenv("DB_PASSWORD"),
		JWTSecretKey:        os.Getenv("JWT_KEY"),
		JWTTokenTTL:         time.Duration(ttl) * time.Hour,
		Salt:                os.Getenv("HASH_SALT"),
		FileStorageLocation: "file_storage",
		ConfigPath:          filepath.Join("../../", "config"),
	}
}

//func GetConfiguration() *configuration {
//	return &configuration{
//		DatabaseName:        `psychology`,
//		DatabaseHost:        `localhost:5432`,
//		DatabaseUser:        `postgres`,
//		DatabasePassword:    `abcd1234`,
//		FileStorageLocation: "file_storage",
//		JWTsecretKey:        "6ae74c638877b9b6fa5d92dd9f3adb3b1",
//		JWTtokenTTL:         12 * time.Hour,
//	}
//}
