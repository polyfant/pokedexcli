package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	
)

// GetLocation -
func (c *Client) GetLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName

	if val, exists := c.cache[url]; exists {
		locationResp := Location{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return Location{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	locationResp := Location{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return Location{}, err
	}

	c.cache[url] = dat

	return locationResp, nil
}
