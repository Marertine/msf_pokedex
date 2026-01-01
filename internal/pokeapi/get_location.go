package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(name string) (Location, error) {
	// similar structure to ListLocations, but:
	// - no pagination
	// - URL includes "/location-area/" + name
	// - unmarshal into Location

	url := baseURL + "/location-area"
	if name != "" {
		url = url + "/" + name
	}

	//if pageURL != nil {
	//	url = *pageURL
	//}

	if data, ok := c.cache.Get(url); ok {
		var locationsResp Location
		if err := json.Unmarshal(data, &locationsResp); err != nil {
			return Location{}, err
		}
		return locationsResp, nil
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

	c.cache.Add(url, dat)

	var locationsResp Location
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return Location{}, err
	}

	return locationsResp, nil
}
