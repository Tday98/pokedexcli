package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(cfg *config) error {
	URL := "https://pokeapi.co/api/v2/location-area"
	if cfg.nextLocationsURL != nil {
		URL = *cfg.nextLocationsURL
	}
	resp, err := http.Get(URL)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var response Response

	if err := json.Unmarshal(dat, &response); err != nil {
		return err
	}

	fmt.Println("next URL:", *response.Next)
	fmt.Println("prev URL:", response.Previous)

	cfg.nextLocationsURL = response.Next
	cfg.prevLocationsURL = response.Previous

	for _, location := range response.Results {
		fmt.Printf("%s\n", location.Name)
	}

	return nil
}
