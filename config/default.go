package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/fcm-service/utils"
	"github.com/kataras/golog"
)

var Props *Config

type Config struct {
	Verbose       uint8
	Env           string
	ServerAddress string
	FirebaseJson  string
}

func Parse(location string) bool {
	var configPath string
	if location == "" {
		configPath = utils.GetEnv("CONFIG", "config/")
		configPath = fmt.Sprintf("%s%s.json", configPath, "config")
	} else {
		configPath = location
	}
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		golog.Error("error in reading config file from ", configPath, " error - ", err)
		return false
	}
	err = json.Unmarshal(data, &Props)
	if err != nil {
		golog.Error("config json unmarshal error - ", configPath, " error - ", err)
		return false
	}
	if Props.Verbose == 1 {
		golog.SetLevel("debug")
	} else {
		golog.SetLevel("warn")
	}
	return true
}
