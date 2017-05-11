package main

import (
	"encoding/json"
	"os"
)

type ConfigLocation struct {
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Distance    float64 `json:"distance"`
	LimitLength float64 `json:"limit_length"`
	Level       int     `json:"level"`
}

func ReadConfigJson(pathFile string) (ConfigLocation, error) {
	var config ConfigLocation
	configFile, err := os.Open(pathFile)
	if err != nil {
		return config, err
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil

}
