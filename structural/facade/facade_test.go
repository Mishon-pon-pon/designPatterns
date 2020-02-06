package facade

import "testing"

func TestOpenWeatherMap_responseParser(t *testing.T) {
	r := getMockData()
	openWeatherMap := CurrentWeatherData{APIkey: "c5c38d0651c7c91f9b7256d00493cd7"}

	weather, err := openWeatherMap.responseParser(r)
	if err != nil {
		t.Fatal(err)
	}
	if weather.ID != 3117735 {
		t.Errorf("Madrid id is 3117735, not %d\n", weather.ID)
	}
	weather, err = openWeatherMap.GetByCityAndCountryCode("Moscow", "ru")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("The temperature in Moscow: %f", weather.Main.Temp-273.15)
}
