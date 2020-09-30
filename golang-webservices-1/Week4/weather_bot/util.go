package main

import (
	"encoding/json"
	"fmt"
	"github.com/caarlos0/env"
	es "github.com/pkg/errors"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
)

//envDefault:"./.env"

type config struct {
	BotToken   string `env:"BOT_TOKEN"`
	Port       string `env:"PORT"`
	WebhookURL string `env:"WEBHOOK"`
	WeatherAPI string `env:"WEATHER_API"`
	Language   string `env:"LANGUAGE"`
}

func parseConfig() config {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", cfg)
	return cfg
}

type FullWeatherReport struct {
	Data     []Stat `json:"data"`
	Count    int    `json:"count"`
	CityName string `json:"city_name"`
}

type Stat struct {
	RelativeHumidity float64 `json:"rh"`
	PartOfDay        string  `json:"pod"`
	PressureMb       float64 `json:"pres"` //pressure in mbar
	CloudCoverage    int     `json:"clouds"`
	CityName         string  `json:"city_name"`
	DateTime         string  `json:"datetime"`
	WindSpeedMs      float64 `json:"wind_spd"`
	WindDirection    string  `json:"wind_cdir"`
	SunsetTime       string  `json:"sunset"`
	Snow             int     `json:"snow"`
	IndexUV          float64 `json:"uv"`
	Precipitation    float64 `json:"precip"`
	SunriseTime      string  `json:"sunrise"`
	IndexAirQuality  int     `json:"aqi"`
	Weather          `json:"weather"`
	Temperature      float64 `json:"temp"`
	FeelsLikeTemp    float64 `json:"app_temp"`
}

type Weather struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

func formatAqi(data Stat) string {
	/*Formats AQI into a string,
	adding an emoji and a comment
	*/
	fmtAqi := ""
	prefix := "Air Quality Index: "
	var aqi = data.IndexAirQuality
	switch {
	case aqi <= 50:
		fmtAqi = fmt.Sprintf(
			"%v%v %c Good\n", prefix, aqi, emojis["Check"],
		)
	case aqi > 50 && aqi <= 100:
		fmtAqi = fmt.Sprintf(
			"%v%v %c Moderate\n", prefix, aqi, emojis["Check"],
		)
	case aqi > 101 && aqi <= 150:
		fmtAqi = fmt.Sprintf(
			"%v%v %c Unhealthy for Sensitive Groups\n", prefix, aqi, emojis["Smoking"],
		)
	case aqi > 151 && aqi <= 200:
		fmtAqi = fmt.Sprintf(
			"%v%v %c Unhealthy\n", prefix, aqi, emojis["Warning"],
		)
	case aqi > 201 && aqi <= 300:
		fmtAqi = fmt.Sprintf(
			"%v%v %c Very Unhealthy\n", prefix, aqi, emojis["Warning"],
		)
	case aqi > 301 && aqi <= 500:
		fmtAqi = fmt.Sprintf(
			"%v%v %c Hazardous\n", prefix, aqi, emojis["Cross"],
		)
	default:
		fmtAqi = fmt.Sprintf(
			"%v%c Unknown\n", prefix,
			emojis["Question"],
		)
	}
	return fmtAqi
}
func formatCity(data Stat) string {
	/*Formats City into a string
	NB: only to be used inside formatNow
	*/
	city := data.CityName
	fmtCity := fmt.Sprintf(
		"Weather in %v:\n", city,
	)
	return fmtCity
}
func formatFeelsLikeTemp(data Stat) string {
	/*Formats FeelsLike Temperature into a string,
	adding a sign
	*/
	fmtFeelsLikeTemp := ""
	prefix := "Feels like"
	feelsLikeTemp := data.FeelsLikeTemp
	if feelsLikeTemp > 0 {
		fmtFeelsLikeTemp = fmt.Sprintf("%v +%v\n", prefix, feelsLikeTemp)
		return fmtFeelsLikeTemp
	}
	fmtFeelsLikeTemp = fmt.Sprintf("%v %v/n", prefix, feelsLikeTemp)
	return fmtFeelsLikeTemp
}
func formatHumidity(data Stat) string {
	/*Formats humidity into a string,
	adding an emoji and a comment
	*/
	fmtHumidity := ""
	humidity := data.RelativeHumidity
	emoji := emojis["Check"]
	comment := "Comfortable"
	switch {
	case humidity < 30:
		emoji = emojis["Dry"]
		comment = "Dry"
	case humidity > 50:
		emoji = emojis["Wet"]
		comment = "Wet"
	}
	fmtHumidity = fmt.Sprintf(
		"Humidity: %c %v (%v%%)\n", emoji, comment, math.Round(humidity),
	)
	return fmtHumidity
}
func formatPressure(data Stat) string {
	/*Formats pressure into a string,
	adding an emoji and a sign
	*/
	fmtPressure := ""
	pressure := data.PressureMb
	norm := 1013.25
	emoji := emojis["Check"]
	switch {
	case pressure < norm:
		emoji = emojis["Down"]
	case pressure > norm:
		emoji = emojis["Up"]
	}
	fmtPressure = fmt.Sprintf(
		"Pressure: %c %v mb (%v mmHg)\n", emoji, math.Round(pressure), millibarsToMmHg(math.Round(pressure)),
	)
	return fmtPressure
}
func formatSun(data Stat) string {
	/*Formats sun related data into a string
	 */
	fmtSun := fmt.Sprintln(
		"Sunrise:", data.SunriseTime, "Sunset:", data.SunsetTime,
	)
	return fmtSun
}
func formatTemperature(data Stat) string {
	/*Formats Temperature into a string,
	adding an emoji and a sign
	NB: newline character is intentionally omitted
	*/
	fmtTemperature := ""
	temperature := data.Temperature
	tempString := fmt.Sprint(temperature)
	if temperature > 0 {
		tempString = "+" + fmt.Sprint(temperature)
	}
	fmtTemperature = fmt.Sprintf(
		"%c %v ", emojis["Thermometer"], tempString,
	)
	return fmtTemperature
}
func formatUv(data Stat) string {
	/*Formats UV into a string,
	adding an emoji and a piece of advice
	*/
	if data.PartOfDay != "d" {
		return ""
	}
	fmtUv := ""
	prefix := "UV Index:"
	uv := data.IndexUV
	switch {
	case uv <= 2:
		fmtUv = fmt.Sprintf(
			"%v %v %v %v 1h+.\n", prefix, commentsEn["no"], commentsEn["hat"], commentsEn["sunburn"],
		)
	case uv > 2 && uv <= 5:
		fmtUv = fmt.Sprintf(
			"%v %v %v %v %v 40 min.\n", prefix, commentsEn["little"], commentsEn["hat"], commentsEn["spf15"],
			commentsEn["sunburn"],
		)
	case uv > 5 && uv <= 7:
		fmtUv = fmt.Sprintf(
			"%v %v %v %v %v %v 30 min.\n", prefix, commentsEn["high"], commentsEn["hat"], commentsEn["spf30"],
			commentsEn["cover"], commentsEn["sunburn"],
		)
	case uv > 7 && uv <= 10:
		fmtUv = fmt.Sprintf(
			"%v %v %v %v %v %v 20 min.\n", prefix, commentsEn["vhigh"], commentsEn["hat"], commentsEn["spf50"],
			commentsEn["cover"], commentsEn["sunburn"],
		)
	case uv > 10:
		fmtUv = fmt.Sprintf(
			"%v %v 20 min.\n", prefix, commentsEn["extreme"],
		)
	}
	return fmtUv
}
func formatWeatherCode(data Stat) string {
	/*Formats WeatherCode into a string,
	adding an emoji and a description
	*/
	fmtWeatherCode := ""
	emoji := emojis["Question"]
	code := data.Weather.Code
	switch code {
	case 200, 201, 202:
		emoji = emojis["ThunderRain"]
	case 230, 231, 232, 233:
		emoji = emojis["Thunder"]
	case 300, 301, 302:
		emoji = emojis["Umbrella"]
	case 500, 501:
		emoji = emojis["UmbrellaRain"]
	case 502, 511, 520, 521, 522:
		emoji = emojis["CloudRain"]
	case 600, 601, 602, 610, 621, 622, 623:
		emoji = emojis["CloudSnow"]
	case 611, 612:
		emoji = emojis["Saturn"]
	case 700, 711, 721, 731, 741, 751:
		emoji = emojis["Fog"]
	case 800:
		emoji = emojis["Sun"]
	case 801, 802:
		emoji = emojis["SunSCloud"]
	case 803:
		emoji = emojis["SunMCloud"]
	case 804:
		emoji = emojis["Cloud"]
	case 900:
		emoji = emojis["Comet"]
	}
	fmtWeatherCode = fmt.Sprintf(
		"%c %v\n", emoji, data.Weather.Description,
	)
	return fmtWeatherCode
}
func formatWind(data Stat) string {
	/*Formats wind direction data into a string,
	adding emojis and speed information
	*/
	fmtWind := ""
	direction := data.WindDirection
	speed := data.WindSpeedMs
	emoji := emojis["Balloon"]
	code, ok := emojis[direction]
	if ok {
		emoji = code
	}
	fmtWind = fmt.Sprintf(
		"%c  Wind: %c %v %v m/sec\n", emojis["Wind"], emoji, direction, math.Round(speed),
	)
	return fmtWind
}

var lineFormatter = []func(data Stat) string{
	formatCity,
	formatTemperature, formatWeatherCode,
	formatFeelsLikeTemp,
	formatWind,
	formatHumidity,
	formatPressure,
	formatSun,
	formatUv,
	formatAqi,
}

func (wr *FullWeatherReport) formatNow() string {
	n := 0
	var res string
	for _, formatter := range lineFormatter {
		res += formatter(wr.Data[n])
	}
	return res
}

func getForecast(loc *tgbotapi.Location, period string, cfg config) ([]byte, error) {
	/*
	 */
	lat := fmt.Sprint(loc.Latitude)
	long := fmt.Sprint(loc.Longitude)
	baseWeatherURL := "https://weatherbit-v1-mashape.p.rapidapi.com/"
	url := fmt.Sprintf(
		"%v%vlang=%v&lat=%v&lon=%v", baseWeatherURL, forecasts[period], cfg.Language, lat, long,
	)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("x-rapidapi-host", "weatherbit-v1-mashape.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", cfg.WeatherAPI)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		err = es.Wrap(err, "failed to perform request")
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		err = es.Wrap(err, "failed to read from response body")
		return nil, err
	}
	fmt.Printf("BODY: %v\n\n", string(body))
	return body, nil
}

func parseWeather(weather []byte) (FullWeatherReport, error) {
	/*
	 */
	data := FullWeatherReport{}
	err := json.Unmarshal(weather, &data)
	if err != nil {
		err = es.Wrap(err, "failed to unmarshal")
		return data, err
	}
	return data, nil
}

func millibarsToMmHg(millibars float64) float64 {
	/*Converting pressure in millibars into millimeters of mercury
	 */
	return millibars * 0.75
}

func pickASaying(sayings []string) string {
	/*Picking a random saying
	 */
	randomIndex := rand.Intn(len(sayings))
	return sayings[randomIndex]
}
