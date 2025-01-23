package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) ListPokemons(locationName string) (PokemonEncounters, error) {
	url := baseURL + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		pokemonsResp := PokemonEncounters{}
		err := json.Unmarshal(val, &pokemonsResp)
		if err != nil {
			return PokemonEncounters{}, err
		}

		return pokemonsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonEncounters{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil || resp.StatusCode > 299 {
		return PokemonEncounters{}, errors.New("location not found")
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonEncounters{}, err
	}

	pokemonsResp := PokemonEncounters{}
	err = json.Unmarshal(dat, &pokemonsResp)
	if err != nil {
		return PokemonEncounters{}, err
	}

	c.cache.Add(url, dat)
	return pokemonsResp, nil
}
