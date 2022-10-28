package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	var cityName string
	fmt.Println("Введите город, погоду которого, хотите узнать: ")
	_, err := fmt.Scan(&cityName)
	if err != nil {
		fmt.Println("произошла ошибка")
		os.Exit(1) // в операционных системах, если программа возвращает не нулевой код возврата это считается что программа завершилась с ошибкой
	}
	APIKey := os.Getenv("W_API_KEY")
	geo, err := GetGeo(cityName, APIKey)
	if err != nil {
		panic(err)
	}

	weather, err := GetWeather(geo, APIKey)
	if err != nil {
		panic(err)
	}

	result := fmt.Sprintf("Погода в городе %s\nТемпература: %2.1fС°, ощущается как %2.1fC°\n",
		weather.Name, weather.Main.Temp, weather.Main.FeelsLike)

	fmt.Println(result)

}

func GetGeo(cityName, APIKey string) (Geo, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, _ := http.NewRequest("GET", fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s&appid=%s", cityName, APIKey), nil)
	res, err := client.Do(req)
	if err != nil {
		return Geo{}, fmt.Errorf("cannot make request: %w", err)
	}
	defer res.Body.Close()

	var gr GeoResponse
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Geo{}, fmt.Errorf("cannot read from response: %w", err)
	}

	err = json.Unmarshal(body, &gr)
	if err != nil {
		return Geo{}, fmt.Errorf("cannot unmarshal to struct: %w", err)
	}

	var geo Geo
	for _, v := range gr {
		geo.Country = v.Country
		geo.Lat = v.Lat
		geo.Name = v.Name
		geo.Lon = v.Lon
		break
	}

	return geo, nil
}

func GetWeather(geo Geo, APIKey string) (WeatherResponse, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s&lang=ru&units=metric", geo.Lat, geo.Lon, APIKey), nil)
	res, err := client.Do(req)
	if err != nil {
		return WeatherResponse{}, fmt.Errorf("cannot make request: %w", err)
	}
	defer res.Body.Close()

	var wr WeatherResponse
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return WeatherResponse{}, fmt.Errorf("cannot read from response: %w", err)
	}

	err = json.Unmarshal(body, &wr)
	if err != nil {
		return WeatherResponse{}, fmt.Errorf("cannot unmarshal to struct: %w", err)
	}
	return wr, nil
}
