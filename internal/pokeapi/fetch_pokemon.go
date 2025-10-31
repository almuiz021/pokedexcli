package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (RespPokemon, error) {

	respPokemon := RespPokemon{}
	url := baseUrl + "pokemon/" + name

	if cacheData, exists := c.myCache.Get(url); exists {
		if err := json.Unmarshal(cacheData, &respPokemon); err != nil {
			return respPokemon, fmt.Errorf("error getting unmarshalled cache data: %s", err)
		}
		return respPokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return respPokemon, fmt.Errorf("error requesting data: %s", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return respPokemon, fmt.Errorf("error in getting response: %s", err)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return respPokemon, fmt.Errorf("error in reading response body: %s", err)
	}
	defer res.Body.Close()

	if err := json.Unmarshal(data, &respPokemon); err != nil {
		return respPokemon, fmt.Errorf("error getting unmarshalled data: %s", err)
	}

	c.myCache.Add(url, data)

	return respPokemon, nil
}
