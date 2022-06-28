package main

type apiConfigData struct {
	OpenWeatherMapApiKey string `json: "OpenWeatherMapApiKey"`
}

type weatherData struct {
	Name string `json: "name"`
	Main struct {
		Kelvin float64 `json: "temp"`
	} `json: "main"`
}

func loadApiConfigKey(filename string) {

}