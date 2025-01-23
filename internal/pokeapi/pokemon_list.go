package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemons(locationName string) (RespPokenons, error) {
	url := baseURL + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		pokemonsResp := RespPokenons{}
		err := json.Unmarshal(val, &pokemonsResp)
		if err != nil {
			return RespPokenons{}, err
		}

		return pokemonsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokenons{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokenons{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokenons{}, err
	}

	pokemonsResp := RespPokenons{}
	err = json.Unmarshal(dat, &pokemonsResp)
	if err != nil {
		return RespPokenons{}, err
	}

	c.cache.Add(url, dat)
	return pokemonsResp, nil
}
