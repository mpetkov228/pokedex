package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (PokemonRes, error) {
	url := baseUrl + "/pokemon/" + name

	if data, ok := c.cache.Get(url); ok {
		pokemonRes := PokemonRes{}
		err := json.Unmarshal(data, &pokemonRes)
		if err != nil {
			return PokemonRes{}, err
		}
		return pokemonRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonRes{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonRes{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonRes{}, err
	}

	pokemonRes := PokemonRes{}
	err = json.Unmarshal(data, &pokemonRes)
	if err != nil {
		return PokemonRes{}, err
	}

	c.cache.Add(url, data)
	return pokemonRes, nil
}
