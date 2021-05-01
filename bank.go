package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func computePointsDiff(currentPoints, previousPoints map[int]uint64) map[int]uint64 {
	diff := make(map[int]uint64)
	for team, current := range currentPoints {
		previous := previousPoints[team]
		diff[team] = current - previous
		log.Printf("Team %d: %d - %d = %d\n", team, current, previous, diff[team])
	}
	return diff
}

type Deposit struct {
	Team   int    `json:"team"`
	Amount string `json:"amount"`
}

func makeDeposit(client http.Client, teamPoints map[int]uint64) {
	depositUrl := conf.BankBaseUrl + "/api/bank/accounts/deposits/"

	var deposits []Deposit
	for team, points := range teamPoints {
		deposit := Deposit{
			Team:   team,
			Amount: fmt.Sprintf("%d.00", points*conf.BankAmountPerPoint),
		}
		deposits = append(deposits, deposit)
	}

	data := map[string]interface{}{
		"deposits":    deposits,
		"description": "Payment for service uptime",
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", depositUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(conf.BankUsername, conf.BankPassword)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		var errBody string
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			errBody = string(body)
		}
		panic(fmt.Sprintf("Failed to deposit, Code: %s. Error: %s", resp.Status, errBody))
	}
}
