package facade

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type CurrentWeatherDataRetriever interface {
	GetByCityAndCountryCode(city, countryCode string) (Weather, error)
	GetByGeoCoordinates(lat, lon float64) (Weather, error)
}

type CurrentWeatherData struct {
	APIkey string
}

func (c *CurrentWeatherData) GetByGeoCoordinates(lat, lon float32) (weather *Weather, err error) {
	return c.doRequest(
		fmt.Sprintf("http://api.openweathermap.org/data/2.5/weatherq=%f,%f&APPID=%s", lat, lon, c.APIkey),
	)
}

func (c *CurrentWeatherData) GetByCityAndCountryCode(city, countryCode string) (weather *Weather, err error) {
	return c.doRequest(
		fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s,%s&APPID=%s", city, countryCode, c.APIkey),
	)
}

func (c *CurrentWeatherData) doRequest(uri string) (weather *Weather, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "aplication/json")
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		byt, errMsg := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if errMsg == nil {
			errMsg = fmt.Errorf("%s", string(byt))
		}
		err = fmt.Errorf("Status code was %d, aborting. Error message was:\n%s\n", resp.StatusCode, errMsg)
		return
	}
	weather, err = c.responseParser(resp.Body)
	resp.Body.Close()
	return
}

func getMockData() io.Reader {
	response := `{
	"coord":{"lon":-3.7,
			"lat":40.42},
	"weather" :[{"id":803,
				"main":"Clouds",
				"description":"brokenclouds",
				"icon":"04n"}],
	"base":"stations",
	"main":{"temp":303.56,
			"pressure":1016.46,
			"humidity":26.8,
			"temp_min":300.95,
			"temp_max":305.93},
	"wind":{"speed":3.17,
			"deg":151.001},
	"rain":{"3h":0.0075},
	"clouds":{"all":68},
	"dt":1471295823,
	"sys":{	"type":3,
			"id":1442829648,
			"message":0.0278,
			"country":"ES",
			"sunrise":1471238808,
			"sunset":1471288232},
	"id":3117735,
	"name":"Madrid",
	"cod":200}`

	r := bytes.NewReader([]byte(response))
	return r
}

func (p *CurrentWeatherData) responseParser(body io.Reader) (*Weather, error) {
	w := new(Weather)
	err := json.NewDecoder(body).Decode(w)
	if err != nil {
		return nil, err
	}
	return w, nil
}
