package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) FetchLocations(pageUrl *string) (RespLocations, error) {

	url := baseUrl + "location-area"
	if pageUrl != nil {

		url = *pageUrl
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocations{}, fmt.Errorf("error creating new request: %w", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocations{}, fmt.Errorf("error getting response: %w", err)
	}

	respLocations := RespLocations{}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespLocations{}, fmt.Errorf("error reading data: %w", err)
	}

	if err := json.Unmarshal(data, &respLocations); err != nil {
		return RespLocations{}, fmt.Errorf("error while unmarshalling data: %s into JSON: %w", string(data), err)
	}

	return respLocations, nil
}
