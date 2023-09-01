package util

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

const (
	MockDelayUrl     = "https://run.mocky.io/v3/122c2796-5df4-461c-ab75-87c1192b17f7"
	DefaultMockValue = 15
)

type Response struct {
	Status bool `json:"status"`
	Data   struct {
		Eta int `json:"eta"`
	} `json:"data"`
}

func MockDelay() int {
	res, err := http.Get(MockDelayUrl)
	if err != nil {
		log.Println(err.Error())
		return DefaultMockValue
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err.Error())
		return DefaultMockValue
	}

	response := Response{}
	err = json.Unmarshal(b, &response)
	if err != nil {
		log.Println(err.Error())
		return DefaultMockValue
	}

	return response.Data.Eta
}
