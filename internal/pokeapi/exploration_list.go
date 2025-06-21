package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ExploreArea(areaName string) (RespExploreLocations, error) {
	url := ""
	if areaName != "" {
		url = baseURL + "/location-area/" + areaName
	} else {
		return RespExploreLocations{}, nil
	}

	if cached, ok := c.cache.Get(url); ok {
		explorationResp := RespExploreLocations{}
		err := json.Unmarshal(cached, &explorationResp)
		if err == nil {
			return explorationResp, nil
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespExploreLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespExploreLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespExploreLocations{}, err
	}

	c.cache.Add(url, dat)

	explorationResp := RespExploreLocations{}
	err = json.Unmarshal(dat, &explorationResp)
	if err != nil {
		return RespExploreLocations{}, err
	}

	return explorationResp, nil
}
