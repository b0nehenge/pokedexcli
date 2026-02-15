package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (LocationList, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResp := LocationList{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return LocationList{}, err
		}

		return locationsResp, nil
	}

	resp, err := http.Get(url)

	if err != nil {
		return LocationList{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationList{}, err
	}

	locationsResp := LocationList{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return LocationList{}, err
	}

	c.cache.Add(url, dat)
	return locationsResp, nil
}

func (c *Client) ExploreLocation(location string) (Location, error) {
	url := baseURL + "/location-area/" + location

	if val, ok := c.cache.Get(url); ok {
		locationsResp := Location{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return Location{}, err
		}

		return locationsResp, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return Location{}, err
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	locationsResp := Location{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(url, dat)
	return locationsResp, nil
}
