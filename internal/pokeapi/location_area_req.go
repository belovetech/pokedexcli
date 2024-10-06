package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreaResp, error) {
	fullUrl := baseURL + "/location-area"

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

	var location LocationAreaResp
	err = json.NewDecoder(res.Body).Decode(&location)
	if err != nil {
		return LocationAreaResp{}, err
	}

	return location, nil
}
