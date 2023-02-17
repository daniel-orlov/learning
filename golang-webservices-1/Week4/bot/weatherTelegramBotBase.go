package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/telegram-bot-api.v4"
	"io/ioutil"
	"net/http"
	"math/rand"
	es "github.com/pkg/errors"
	"math"
)

const (
	BotToken		= ""
	WebhookURL		= ""
	WeatherAPIToken = ""
	DefaultMessage	= "Hi! I am WeartheBot and I cannot spell."
	CoordsAccepted 	= "Gotcha! Please choose the forecast period."
	BackToMainMenu	= "Which place interests you most now?"
	BaseWeatherUrl	= "https://weatherbit-v1-mashape.p.rapidapi.com/"
	Language 		= "lang=en&"
)

var sayings = []string{
	"Let's see weather or not you should leave umbrella at home.",
	"The best and the worst thing about the weather is that it changes.",
	"Conversation about the weather is the last refuge of the unimaginative.",
	"Weather forecast for tonight: dark.",
	"There is no such thing as bad weather, only different kinds of good weather.",
	"I've lived in good climate, and it bores the hell out of me.",
	"Do you know that bad weather always looks worse through a window?",
}

var forecasts = map[string]string{
	"Now": 		"current?",
	"2 days": "forecast/hourly?hours=48&",
	"5 days": 	"forecast/3hourly?",
	"10 days": 	"forecast/daily?",
}

var emojis = map[string]int{
	"Wind": 		127788,
	"Thermometer": 	127777,
	"Comet": 		9732,
	"N": 			11015,
	"E": 			11013,
	"W": 			10145,
	"S": 			11014,
	"SE": 			8598,
	"SW": 			8599,
	"NW": 			8600,
	"NE": 			8601,
	"Balloon": 		127880,
	"Check":		9989,
	"Warning":		9888,
	"Smoking":		128684,
	"Cross":		10060,
	"Question":		10067,
	"ThunderRain":	9928,
	"Thunder":		127785,
	"Umbrella":		9730,
	"UmbrellaRain": 9748,
	"CloudRain": 	127783,
	"CloudSnow": 	127784,
	"Saturn":		129680,
	"Fog":			127787,
	"Sun":			9728,
	"SunSCloud":	127780,
	"SunMCloud":	9925,
	"Cloud":		9729,

}

var loc *tgbotapi.Location

type FullWeatherReport struct {
	Data []struct {
		RelativeHumidity		float64 	`json:"rh"`
		PartOfDay				string 		`json:"pod"`
		PressureMb         		float64 	`json:"pres"` //pressure in mb
		CloudCoverage       	int     	`json:"clouds"`
		CityName     			string  	`json:"city_name"`
		DateTime 				string  	`json:"datetime"`
		WindSpeedMs      		float64 	`json:"wind_spd"`
		WindDirection     		string  	`json:"wind_cdir"`
		SunsetTime       		string  	`json:"sunset"`
		Snow         			int     	`json:"snow"`
		indexUV           		float64 	`json:"uv"`
		Precipitation       	float64    	`json:"precip"`
		SunriseTime      		string  	`json:"sunrise"`
		indexAirQuality			int    		`json:"aqi"`
		Weather      			struct {
			Code        		int    		`json:"code"`
			Description 		string 		`json:"description"`
		} 									`json:"weather"`
		Temperature      		float64 	`json:"temp"`
		FeelsLikeTemp   		float64 	`json:"app_temp"`
	} 										`json:"data"`
	Count 						int 		`json:"count"`
	CityName     				string  	`json:"city_name"`
}

func (wr *FullWeatherReport) formatWindDirection(n int) int {
	direction := wr.Data[n].WindDirection
	code, ok := emojis[direction]
	if ok {
		return code
	} else {
		return emojis["Balloon"]
	}
}

func (wr *FullWeatherReport) formatAqi(n int) (int, string) {
	aqi := wr.Data[n].indexAirQuality
	switch {
	case aqi <= 50:
		return emojis["Check"], "Good"
	case aqi > 50 && aqi <= 100:
		return emojis["Check"], "Moderate"
	case aqi > 101 && aqi <= 150:
		return emojis["Smoking"], "Unhealthy for Sensitive Groups"
	case aqi > 151 && aqi <= 200:
		return emojis["Warning"], "Unhealthy"
	case aqi > 201 && aqi <= 300:
		return emojis["Warning"], "Very Unhealthy"
	case aqi > 301 && aqi <= 500:
		return emojis["Cross"], "Hazardous"
	default:
		return emojis["Question"], "Unknown"
	}
}

func (wr *FullWeatherReport) formatUV(n int) string {
	uv := wr.Data[n].indexUV
	a := map[string]string{
		"no": "👒 No danger to the average person. ",
		"little": "👓 Little risk of harm. ",
		"high": "☝ High risk of harm. ",
		"vhigh": "☝☝ Very high risk of harm. ",
		"extreme": "❌ Extreme risk of harm. Stay indors!",
		"hat": "Wear a hat and/or sunglasses. ",
		"spf15": "Use SPF 15+ Sunscreen. ",
		"spf30": "Use SPF 30+ Sunscreen. ",
		"spf50": "Use SPF 50 Sunscreen. ",
		"cover": "Cover the body with clothing. Avoid the sun if possible. ",
		"sunburn": "Sun burn time is ",
	}
	switch {
	case uv <= 2:
		return a["no"] + a["hat"] + a["sunburn"] + "1h+."
	case uv > 2 && uv <= 5:
		return a["little"] + a["hat"] + a["spf15"] + a["sunburn"] + "40 min."
	case uv > 5 && uv <= 7:
		return a["high"] + a["hat"] + a["spf30"] + a["cover"] + a["sunburn"] + "30 min."
	case uv > 7 && uv <= 10:
		return a["vhigh"] + a["hat"] + a["spf50"] + a["cover"] + a["sunburn"] + "20 min."
	case uv > 10:
		return a["extreme"]
	default:
		return ""
	}
}

func (wr *FullWeatherReport) formatCode(n int) int {
	code := wr.Data[n].Weather.Code
	switch code {
	case 200, 201, 202:
		return emojis["ThunderRain"]
	case 230, 231, 232, 233:
		return emojis["Thunder"]
	case 300, 301, 302:
		return emojis["Umbrella"]
	case 500, 501:
		return emojis["UmbrellaRain"]
	case 502, 511, 520, 521, 522:
		return emojis["CloudRain"]
	case 600, 601, 602, 610, 621, 622, 623:
		return emojis["CloudSnow"]
	case 611, 612:
		return emojis["Saturn"]
	case 700, 711, 721, 731, 741, 751:
		return emojis["Fog"]
	case 800:
		return emojis["Sun"]
	case 801, 802:
		return emojis["SunSCloud"]
	case 803:
		return emojis["SunMCloud"]
	case 804:
		return emojis["Cloud"]
	case 900:
		return emojis["Comet"]
	}
	return emojis["Question"]
}

func (wr *FullWeatherReport) formatTemperature(n int) string {
	temperature := wr.Data[n].Temperature
	if temperature > 0 {
		return "+" + fmt.Sprint(temperature)
	}
	return fmt.Sprint(temperature)
}

func (wr *FullWeatherReport) formatFeelsLikeTemp(n int) string {
	feelsLikeTemp := wr.Data[n].FeelsLikeTemp
	if feelsLikeTemp > 0 {
		return "+" + fmt.Sprint(feelsLikeTemp)
	}
	return fmt.Sprint(feelsLikeTemp)
}

func (wr *FullWeatherReport) formatNow() string {
	n := 0
	l0 := fmt.Sprintf("Weather in %v:\n", wr.Data[n].CityName)
	l1 := fmt.Sprintf(
		"%c %v %c %v\n", emojis["Thermometer"], wr.formatTemperature(n),
		wr.formatCode(n), wr.Data[n].Weather.Description,
	)
	l2 := fmt.Sprintln("Feels like", wr.formatFeelsLikeTemp(n))
	l3 := fmt.Sprintf(
		"%c  Wind: %c %v %v m/sec\n", emojis["Wind"], wr.formatWindDirection(n),
		wr.Data[n].WindDirection, math.Round(wr.Data[n].WindSpeedMs),
	)
	pressure := wr.Data[n].PressureMb
	l4 := fmt.Sprintf(
		"Pressure: %v mb (%v mmHg)\n", math.Round(pressure),
		millibarsToMmHg(math.Round(pressure)),
	)
	l5 := fmt.Sprintf("Humidity: %v%%\n", math.Round(wr.Data[n].RelativeHumidity))
	l6 := fmt.Sprintln("Sunrise:", wr.Data[n].SunriseTime, "Sunset:", wr.Data[n].SunsetTime)
	l7 := ""
	if wr.Data[n].PartOfDay == "d" {
		l7 = fmt.Sprintln("UV Index:", wr.formatUV(n))
	}
	aqiEmoji, aqiDescription := wr.formatAqi(n)
	l8 := fmt.Sprintf("Air Quality Index: %c %v\n", aqiEmoji, aqiDescription)
	return l0 + l1 + l2 + l3 + l4 + l5 + l6 + l7 + l8
}

// func (wr *FullWeatherReport) formatForecast(period, step int) string {
// 	formatted := fmt.Sprintf("%v-h forecast for %v:\n ", period, wr.CityName)
// 	for i := 0; i < period; i += step {
// 		 formatted += fmt.Sprintf(
// 		 	" %c%c%v m/sec\n", emojis["Wind"], wr.formatWindDirection(i),
// 		 math.Round(wr.Data[i].WindSpeedMs),
// 		 )
// 	}
// }

func main() {
	//creating a new BotAPI instance using token
	bot, err := tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		err = es.Wrap(err, "failed to create new BotAPI with given token:")
		fmt.Println(err)
	}

	// bot.Debug = true

	fmt.Printf("Authorized on account %s\n", bot.Self.UserName)
	//setting a webhook
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(WebhookURL))
	if err != nil {
		err = es.Wrap(err, "failed to set WebHook:")
		fmt.Println(err)
	}

	//registers an http handler for a webhook, returns UpdatesChannel (<-chan Update)
	updates := bot.ListenForWebhook("/") //why exactly just root?
	//launching a server on a local host :8080
	go http.ListenAndServe(":8080", nil)
	fmt.Println("start listen :8080")

	mainMenu := tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButtonLocation(`Weather at my location`),
		tgbotapi.NewKeyboardButton(`A different place`),
		tgbotapi.NewKeyboardButton(`Stop`))

	forecastMenu := tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(`Now`),
		tgbotapi.NewKeyboardButton(`2 days`),
		tgbotapi.NewKeyboardButton(`5 days`),
		tgbotapi.NewKeyboardButton(`10 days`),
		tgbotapi.NewKeyboardButton(`< Back`))

	//handling messages from user
	for update := range updates {
		chatID := update.Message.Chat.ID
		userMsg := update.Message
		text := userMsg.Text
		location := userMsg.Location

		switch {
		case location != nil:
			{
				loc = userMsg.Location
				msg := tgbotapi.NewMessage(chatID, CoordsAccepted)
				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(forecastMenu)
				bot.Send(msg)
			}
		}
		switch text {
		case "/start":
			{
    			pick := DefaultMessage + "\n" + pickASaying(sayings)
				msg := tgbotapi.NewMessage(chatID, pick)
				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(mainMenu)
				bot.Send(msg)
			}
		case "< Back":
			{
				msg := tgbotapi.NewMessage(chatID, BackToMainMenu)
				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(mainMenu)
				bot.Send(msg)
			}
		case "Stop":
			{
				msg := tgbotapi.NewMessage(chatID, `Bye!`)
				msg.ReplyMarkup = tgbotapi.ReplyKeyboardHide{HideKeyboard: true}
				bot.Send(msg)
			}
		case "Now":
			{
				forecast, err := getForecast(loc, text)
				if err != nil {
					err = es.Wrap(err, "failed to getForecast:")
					fmt.Println(err)
				}
				wr, err := parseWeather(forecast)
				if err != nil {
					err = es.Wrap(err, "failed to parseWeather:")
					fmt.Println(err)
				}
				repr := wr.formatNow()
				msg := tgbotapi.NewMessage(chatID, repr)
				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(forecastMenu)
				bot.Send(msg)
			}
		case "2 days", "5 days", "10 days":
			{
				text := "Under development"
				msg := tgbotapi.NewMessage(chatID, text)
				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(forecastMenu)
				bot.Send(msg)
			}

		// default:
		// 	{
		// 		msg := tgbotapi.NewMessage(chatID, DefaultMessage)
		// 		bot.Send(msg)
		// 	}
		}
	}
}

func getForecast(loc *tgbotapi.Location, period string) ([]byte, error) {
	lat := fmt.Sprint(loc.Latitude)
	long := fmt.Sprint(loc.Longitude)
	url := BaseWeatherUrl + forecasts[period] + Language + "lat=" + lat + "&lon=" + long

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("x-rapidapi-host", "weatherbit-v1-mashape.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", WeatherAPIToken)

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
	data := FullWeatherReport{}
	err := json.Unmarshal(weather, &data)
	if err != nil {
		err = es.Wrap(err, "failed to unmarshal")
		return data, err
	}
	return data, nil
}

func millibarsToMmHg (millibars float64) float64 {
	/*Converting pressure in millibars into millimiters of mercury
	*/
	return millibars * 0.75
}

func pickASaying(sayings []string) string {
	/*Picking a random saying
	*/
	randomIndex := rand.Intn(len(sayings))
	return sayings[randomIndex]
}
