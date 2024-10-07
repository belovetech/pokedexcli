package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
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
		log.Println("Cache hit")
		var location LocationAreaResp
		err := json.Unmarshal(val, &location)
		if err != nil {
			return LocationAreaResp{}, err
		}
		return location, nil
	}
	log.Println("Cache miss")

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
