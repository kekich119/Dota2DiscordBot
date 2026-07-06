package opendota

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var accountsByPersonalName AccountByPersonalName
var prof Dotka

func searchAccountIdByPersonalName(personalName string) {
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

func searchByAccountId(account_id string) {
	resp, err := http.Get("https://api.opendota.com/api/players/" + account_id)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &prof); err != nil {
		panic(err)
	}

	fmt.Println(prof.ComputedMmr)
}
