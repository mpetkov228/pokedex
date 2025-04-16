package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocations(pageUrl *string) (ResLocations, error) {
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	if entry, ok := c.cache.Get(url); ok {
		resLocations := ResLocations{}
		err := json.Unmarshal(entry, &resLocations)
		if err != nil {
			return ResLocations{}, err
		}
		return resLocations, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResLocations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return ResLocations{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return ResLocations{}, err
	}

	resLocations := ResLocations{}
	err = json.Unmarshal(data, &resLocations)
	if err != nil {
		return ResLocations{}, err
	}

	c.cache.Add(url, data)
	return resLocations, nil
}

func (c *Client) ExploreLocation(location string) (ResEncounters, error) {
	url := baseUrl + "/location-area/" + location

	if val, ok := c.cache.Get(url); ok {
		resEncounters := ResEncounters{}
		err := json.Unmarshal(val, &resEncounters)
		if err != nil {
			return ResEncounters{}, err
		}

		return resEncounters, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResEncounters{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return ResEncounters{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return ResEncounters{}, err
	}

	resEncounters := ResEncounters{}
	err = json.Unmarshal(data, &resEncounters)
	if err != nil {
		return ResEncounters{}, err
	}

	c.cache.Add(url, data)
	return resEncounters, nil
}
