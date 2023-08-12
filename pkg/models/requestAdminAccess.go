package models

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Info struct {
	AdminPassword string `yaml:"ADMIN_SECRET"`
}

var pswd string

func GetAdminPswd() {
	configFile, err := os.Open("config.yaml")
	if err != nil {
		log.Fatalf("failed to open config file: %v", err)
	}
	defer configFile.Close()

	var config Info
	decoder := yaml.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalf("failed to decode config: %v", err)
	}

	pswd = config.AdminPassword
}

func RequestAdminAccess(uname string, enteredPswd string) (bool, error) {
	GetAdminPswd()
	if enteredPswd == pswd {
		db, err := Connection()
		if err != nil {
			return true, err
		}
		query := "UPDATE users SET isAdmin = true WHERE username = (?)"
		_, err = db.Exec(query, uname)
		if err != nil {
			return true, err
		}
		return true, nil
	} else {
		return false, nil
	}
}
