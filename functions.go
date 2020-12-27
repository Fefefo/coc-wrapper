package cocwrapper

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

var token string

// SetAPIToken is the function to set your clash of clans token
func SetAPIToken(api string) {
	token = "Bearer " + api
}

// GetClans is the function to get all the clans that meet all the requirements
func GetClans(stats ReqClans) CocClans {
	var res CocClans
	client := &http.Client{}

	query := "https://api.clashofclans.com/v1/clans"

	req, err := http.NewRequest("GET", query, nil)
	if err != nil {
		log.Panic(err)
		return CocClans{}
	}
	req.Header.Set("User-Agent", "Golang_Wrapper")
	req.Header.Set("Authorization", token)

	q := req.URL.Query()

	if stats.Name != "" {
		q.Add("name", stats.Name)
	}
	if stats.MinLevel != 0 {
		q.Add("minClanLevel", strconv.Itoa(stats.MinLevel))
	}
	if stats.MaxLevel != 0 {
		q.Add("maxClanLevel", strconv.Itoa(stats.MaxLevel))
	}
	if stats.Limit != 0 {
		q.Add("limit", strconv.Itoa(stats.Limit))
	}
	if stats.MinPoints != 0 {
		q.Add("minClanPoints", strconv.Itoa(stats.MinPoints))
	}
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
		return CocClans{}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
		return CocClans{}
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Panic(err)
		return CocClans{}
	}
	return res
}

// GetPlayer is the function to get the player with the given playerTag
func GetPlayer(playerTag string) CocPlayer {
	var res CocPlayer
	client := &http.Client{}

	query := "https://api.clashofclans.com/v1/players/" + url.QueryEscape(playerTag)

	req, err := http.NewRequest("GET", query, nil)
	if err != nil {
		log.Panic(err)
		return CocPlayer{}
	}
	req.Header.Set("User-Agent", "Golang_Wrapper")
	req.Header.Set("Authorization", token)

	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
		return CocPlayer{}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
		return CocPlayer{}
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Panic(err)
		return CocPlayer{}
	}
	return res
}
