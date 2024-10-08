package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/belovetech/pokedexcli.git/internal/pokecache"
)

func (c *Client) ListLocationAreas(pageURL *string, cache *pokecache.Cache) (LocationAreaResp, error) {
	fullUrl := baseURL + "/location-area"

	if pageURL != nil {
		fullUrl = *pageURL
	}

	val, exists := cache.Get(fullUrl)
	if exists {
		var location LocationAreaResp
		err := json.Unmarshal(val, &location)
		if err != nil {
			return LocationAreaResp{}, err
		}
		return location, nil
	}

	if pageURL != nil {
		fullUrl = *pageURL
	}

	req, err := http.NewRequest(http.MethodGet, fullUrl, nil)
	if err != nil {
		return LocationAreaResp{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResp{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationAreaResp{}, fmt.Errorf("error statuscode: %d", res.StatusCode)
	}

	// Read the entire response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}

	// Store the response in cache
	cache.Add(fullUrl, body)

	var location LocationAreaResp
	err = json.Unmarshal(body, &location)
	if err != nil {
		return LocationAreaResp{}, err
	}

	return location, nil
}

func (c *Client) ListPokemonInLocation(locationName string, cache *pokecache.Cache) ([]string, error) {
	fullURL := baseURL + "/location-area/" + locationName

	val, exists := cache.Get(fullURL)
	if exists {
		var pokemonNames []string
		err := json.Unmarshal(val, &pokemonNames)
		if err != nil {
			return []string{}, err
		}
		return pokemonNames, nil
	}

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return []string{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return []string{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 399 {
		return []string{}, fmt.Errorf("error statuscode: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []string{}, err
	}

	var pokemonInLocation PokemonInLocationAreaResp
	err = json.Unmarshal(body, &pokemonInLocation)
	if err != nil {
		return []string{}, err
	}

	pokemanNames, err := getPokemonName(pokemonInLocation)
	if err != nil {
		return []string{}, err
	}

	jsonData, err := json.Marshal(pokemanNames)
	if err != nil {
		return []string{}, err
	}

	cache.Add(fullURL, jsonData)

	return pokemanNames, nil

}

func getPokemonName(pokemonInLocation PokemonInLocationAreaResp) ([]string, error) {
	var pokemonNames []string
	for _, pokeman := range pokemonInLocation.PokemonEncounters {
		pokemanName := pokeman.Pokemon.Name
		pokemonNames = append(pokemonNames, pokemanName)
	}

	return pokemonNames, nil

}
