package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	data "pubGO/pkg/data"
	"strings"
)

func (api *PubgApi) getSeasons() (string, []data.SeasonData) {
	fmt.Printf("Hey, I want the seasons\n")
	url := fmt.Sprintf("%s/seasons", api.url)
	resp, err := api.requestWithToken(url, "GET")

	if err != nil {
		panic(err)
	}
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("")
	var responseData data.SeasonResponse
	err = json.Unmarshal(r, &responseData)
	fmt.Printf("%v+\n", responseData)
	fmt.Println(responseData)

	return string(r), responseData.Data
}

func (api *PubgApi) GetCurrentSeasonID() string {
	_, seasons := api.getSeasons()
	for _, season := range seasons {
		if season.Attributes.IsCurrentSeason {
			return season.ID
		}
	}
	return ""
}

func (api *PubgApi) GetSeasonStats(playerId, seasonId string) (string, data.PlayerSeasonData) {
	fmt.Printf("Hey, I want season %s stats for player %s!\n", seasonId, playerId)
	url := fmt.Sprintf("%s/players/%s/seasons/%s", api.url, playerId, seasonId)
	resp, err := api.requestWithToken(url, "GET")
	r, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(r))
	if err != nil {
		panic(err)
	}
	var seasonData data.PlayerSeasonResponseData
	json.Unmarshal([]byte(r), &seasonData)
	fmt.Printf("%v+\n", seasonData)
	x := seasonData.Data.Attributes.GameModeStats
	fmt.Println(x)
	fmt.Println("relationships", seasonData.Data.Relationships)
	fmt.Println(seasonData.Data.Relationships["player"])
	return string(r), seasonData.Data
}

func (api *PubgApi) GetSeasonStatsForGameMode(seasonId, gameMode string, playerIDs ...string) string {
	fmt.Printf("Hey, I want %s's playerInfo\n", playerIDs)
	url := fmt.Sprintf("%s/seasons/%s/gameMode/%s/players?filter[playerIds]=%s", api.url, seasonId, gameMode, strings.Join(playerIDs, ","))
	resp, err := api.requestWithToken(url, "GET")
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(r))
	return string(r)
}
