package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) FetchLocations(pageUrl *string) (RespLocations, error) {
	respLocations := RespLocations{}

	url := baseUrl + "location-area"
	if pageUrl != nil {

		url = *pageUrl
	}

	var mainData []byte

	cacheData, exists := c.myCache.Get(url)

	if exists {
		mainData = cacheData
	} else {

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespLocations{}, fmt.Errorf("error creating new request: %w", err)
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return RespLocations{}, fmt.Errorf("error getting response: %w", err)
		}

		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return RespLocations{}, fmt.Errorf("error reading data: %w", err)
		}

		mainData = data

		c.myCache.Add(url, data)

	}

	if err := json.Unmarshal(mainData, &respLocations); err != nil {
		return RespLocations{}, fmt.Errorf("error while unmarshalling data: %s into JSON: %w", string(cacheData), err)
	}

	return respLocations, nil
}
