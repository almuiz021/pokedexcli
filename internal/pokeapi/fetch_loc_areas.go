package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocation(areaName string) (RespLocationAreas, error) {

	respLocAreas := RespLocationAreas{}
	url := baseUrl + "location-area/" + areaName

	if cacheData, exists := c.myCache.Get(url); exists {
		if err := json.Unmarshal(cacheData, &respLocAreas); err != nil {

			return RespLocationAreas{}, fmt.Errorf("error unmarshalling data: %s", err)
		}
		return respLocAreas, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationAreas{}, fmt.Errorf("error creating new request: %s", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationAreas{}, fmt.Errorf("error getting response: %s", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespLocationAreas{}, fmt.Errorf("error reading response: %s", err)
	}

	if err := json.Unmarshal(data, &respLocAreas); err != nil {

		return RespLocationAreas{}, fmt.Errorf("error unmarshalling data: %s", err)
	}
	c.myCache.Add(url, data)
	return respLocAreas, nil
}
