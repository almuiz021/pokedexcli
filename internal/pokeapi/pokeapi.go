package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) ListLocationNext() ([]string, error) {
	baseUrl := "https://pokeapi.co/api/v2/location-area/"

	if c.Next != nil {
		baseUrl = *c.Next
	}

	locs, err := c.GetLocations(baseUrl)
	return locs, err

}

func (c *Client) ListLocationPrevious() ([]string, error) {
	baseUrl := "https://pokeapi.co/api/v2/location-area/"

	if c.Previous != nil {
		baseUrl = *c.Previous
	}

	locs, err := c.GetLocations(baseUrl)
	return locs, err

}

func (c *Client) GetLocations(url string) ([]string, error) {

	var locations RespLocations
	locationNames := make([]string, 0)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []string{}, fmt.Errorf("error: %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return []string{}, fmt.Errorf("error: %w", err)
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locations); err != nil {
		return []string{}, fmt.Errorf("error: %w", err)
	}

	c.Next = locations.Next
	c.Previous = locations.Previous

	for _, loc := range locations.Results {
		locationNames = append(locationNames, loc.Name)
	}
	return locationNames, nil
}
