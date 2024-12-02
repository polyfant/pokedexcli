package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// RespShallowPokemon represents the structure of Pokemon list response from PokeAPI
type RespShallowPokemon struct {
	Count    int `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// ListPokemon retrieves a list of Pokemon from the PokeAPI
func (c *Client) ListPokemon(pageURL *string) (RespShallowPokemon, error) {
	url := baseURL + "/pokemon"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		pokemonResp := RespShallowPokemon{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return RespShallowPokemon{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowPokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowPokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowPokemon{}, err
	}

	pokemonResp := RespShallowPokemon{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return RespShallowPokemon{}, err
	}

	c.cache.Add(url, dat)
	return pokemonResp, nil
}
