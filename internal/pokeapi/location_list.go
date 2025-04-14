package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/mpetkov228/pokedex/internal/pokecache"
)

func (c *Client) GetLocations(pageUrl *string, cache *pokecache.Cache) (ResLocations, error) {
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	var resLocations ResLocations

	entry, ok := cache.Get(url)
	if ok {
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
	cache.Add(url, data)

	err = json.Unmarshal(data, &resLocations)
	if err != nil {
		return ResLocations{}, err
	}

	return resLocations, nil
}
