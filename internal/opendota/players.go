package opendota

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var accountsByPersonalName AccountByPersonalName
var prof Dotka
var winrate Winrate
var heroes HeroesPlayer

func SearchAccountIdByPersonalName(personalName string) {
	response, err := http.Get("https://api.opendota.com/api/search?q=" + personalName)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &accountsByPersonalName); err != nil {
		panic(err)
	}
	fmt.Print(accountsByPersonalName)
}

func SearchByAccountId(account_id string) (profile Dotka) {
	resp, err := http.Get("https://api.opendota.com/api/players/" + account_id)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := json.Unmarshal(body, &prof); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(prof.ComputedMmr)
	return prof
}

func GetWinrateByAccountId(account_id string) (winrate Winrate) {

	resp, err := http.Get("https://api.opendota.com/api/players/" + account_id + "/wl")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := json.Unmarshal(body, &winrate); err != nil {
		fmt.Println(err)
		return
	}

	return winrate

}

func GetPlayerHeroesByAccuontId(account_id string) (HeroesPlayer HeroesPlayer) {
	resp, err := http.Get("https://api.opendota.com/api/players/" + account_id + "/heroes")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	json.Unmarshal(body, &heroes)

	return HeroesPlayer

}
