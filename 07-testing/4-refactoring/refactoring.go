package refactoring

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Rates - курсы валют.
type Rates struct {
	Valute struct {
		USD struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"USD"`
	}
}

func rublesToUSD(rubles float64) (float64, error) {
	var data Rates
	const url = "https://www.cbr-xml-daily.ru/daily_json.js"
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return 0, err
	}

	if data.Valute.USD.Value == 0 {
		return 0, errors.New("деление на 0")
	}
	return rubles / float64(data.Valute.USD.Value), nil
}
