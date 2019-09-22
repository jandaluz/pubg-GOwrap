package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	data "pubGO/pkg/data"
	"reflect"
	"strings"
)

type PubgApi struct {
	Platform string
	key      string
	url      string
}

func (api *PubgApi) ApiKey(key string) {
	api.key = key
}

func NewApi(platform, key string) *PubgApi {
	url := fmt.Sprintf("https://api.pubg.com/shards/%s", platform)
	return &PubgApi{
		Platform: platform,
		key:      key,
		url:      url,
	}
}

func (api *PubgApi) GenerateApiUrl(url string) {
	api.url = url
}

func (api PubgApi) requestWithToken(url, method string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	h := fmt.Sprintf("Bearer %s", api.key)
	fmt.Println(h)
	req.Header.Add("Authorization", h)
	req.Header.Add("accept", "application/vnd.api+json")
	c := &http.Client{}

	return c.Do(req)
}

func (api PubgApi) requestWithoutToken(url, method string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("accept", "application/vnd.api+json")
	c := &http.Client{}

	return c.Do(req)
}

func (api PubgApi) GetPlayerInfo(names ...string) (string, []data.PlayerData) {
	fmt.Printf("Hey, I want %s's playerInfo\n", names)
	url := fmt.Sprintf("%s/players?filter[playerNames]=%s", api.url, strings.Join(names, ","))
	fmt.Printf("url: %s", url)

	resp, err := api.requestWithToken(url, "GET")
	if err != nil {
		panic(err)
	}

	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("")
	var responseData data.PlayerResponse
	var responseData2 data.ResponseData
	err = json.Unmarshal(r, &responseData)
	err = json.Unmarshal(r, &responseData2)
	responseString := string(r)

	for _, dataObject := range responseData2.Data {
		fmt.Printf("%T\n", dataObject)
		x := reflect.ValueOf(dataObject)

		dataType := x.MapIndex(reflect.ValueOf("type")).Elem().String()
		fmt.Println("val of type", dataType)
		if dataType == "player" {
			playerData := x.MapIndex(reflect.ValueOf("attributes"))
			fmt.Printf("%v+\n", playerData.String())
			var dataObject data.PlayerData
			fmt.Printf("playerdata: %s\n", playerData.Elem())
			fmt.Println("name:", dataObject.Attributes.Name)
			fmt.Printf("%v+\n", dataObject)
		}

	}
	//fmt.Println(y)
	return responseString, responseData.Data
}

func (api PubgApi) GetPlayerInfoReflection(names ...string) (string, []data.PlayerData) {
	fmt.Printf("Hey, I want %s's playerInfo\n", names)
	url := fmt.Sprintf("https://api.pubg.com/shards/%s/players?filter[playerNames]=%s", api.Platform, strings.Join(names, ","))
	fmt.Printf("url: %s", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	h := fmt.Sprintf("Bearer %s", api.key)
	fmt.Println(h)
	req.Header.Add("Authorization", h)
	req.Header.Add("accept", "application/vnd.api+json")
	c := &http.Client{}

	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}

	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("")
	var responseData data.PlayerResponse
	var responseData2 data.ResponseData
	err = json.Unmarshal(r, &responseData)
	err = json.Unmarshal(r, &responseData2)
	responseString := string(r)

	for _, dataObject := range responseData2.Data {
		fmt.Printf("%T\n", dataObject)
		x := reflect.ValueOf(dataObject)

		dataType := x.MapIndex(reflect.ValueOf("type")).Elem().String()
		fmt.Println("val of type", dataType)
		if dataType == "player" {
			playerData := x.MapIndex(reflect.ValueOf("attributes"))
			fmt.Printf("%v+\n", playerData.String())
			var dataObject data.PlayerData
			fmt.Printf("playerdata: %s\n", playerData.Elem())
			fmt.Println("name:", dataObject.Attributes.Name)
			fmt.Printf("%v+\n", dataObject)
		}

	}
	//fmt.Println(y)
	return string(responseString), responseData.Data
}

func (api PubgApi) GetPlayerInfoByID(iD string) string {
	fmt.Printf("Hey, I want %s's playerInfo\n", iD)
	url := fmt.Sprintf("https://api.pubg.com/shards/%s/players/%s", api.Platform, iD)
	fmt.Printf("url: %s", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	h := fmt.Sprintf("Bearer %s", api.key)
	fmt.Println(h)
	req.Header.Add("Authorization", h)
	req.Header.Add("accept", "application/vnd.api+json")
	c := &http.Client{}

	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}

	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("")
	var responseData data.PlayerResponse
	err = json.Unmarshal(r, &responseData)
	responseString := string(r)
	fmt.Println(responseString)

	//x := responseData.Data
	fmt.Println(responseData.Data)
	for _, player := range responseData.Data {
		fmt.Printf("%v+\n", player.Attributes)
	}
	//fmt.Println(y)
	return string(responseString)
}
