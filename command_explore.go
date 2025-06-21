package main

func commandExplore(areaName string) error {
	explorationResp, err := cfg.pokeapiClient.ExploreArea(areaName)
	return nil
}
