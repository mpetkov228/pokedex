package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PokeapiRes struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string
	Url  string
}

func GetLocationData(url string) (PokeapiRes, error) {
	res, err := http.Get(url)
	if err != nil {
		return PokeapiRes{}, fmt.Errorf("error making request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokeapiRes{}, fmt.Errorf("error reading request data: %w", err)
	}

	var apiResult PokeapiRes

	err = json.Unmarshal(data, &apiResult)
	if err != nil {
		return PokeapiRes{}, fmt.Errorf("error unmarshaling json: %w", err)
	}

	return apiResult, nil
}

func commandMap(config *Config) error {
	url := config.Next

	data, err := GetLocationData(url)
	if err != nil {
		return err
	}

	config.Next = data.Next
	config.Previous = data.Previous

	for _, area := range data.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandMapb(config *Config) error {
	url := config.Previous
	if url == "" {
		fmt.Println("You are on the first page.")
		return nil
	}

	data, err := GetLocationData(url)
	if err != nil {
		return err
	}

	config.Next = data.Next
	config.Previous = data.Previous

	for _, area := range data.Results {
		fmt.Println(area.Name)
	}

	return nil
}
