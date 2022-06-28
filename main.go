package main

import (
	"encoding/json"
	"io/ioutil"
)

type apiConfigData struct {
	OpenWeatherMapApiKey string `json: "OpenWeatherMapApiKey"`
}

type weatherData struct {
	Name string `json: "name"`
	Main struct {
		Kelvin float64 `json: "temp"`
	} `json: "main"`
}

func loadApiConfigKey(filename string) (apiConfigData, error) {

	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return apiConfigData{}, err
	}

	var apiData apiConfigData

	err = json.Unmarshal(bs, &apiData)
	if err != nil {
		return apiConfigData{}, err
	}

	return apiData, nil
}
