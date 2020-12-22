package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func computePointsDiff(currentPoints, previousPoints map[int]uint) map[int]uint {
	diff := make(map[int]uint)
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

func makeDeposit(client http.Client, teamPoints map[int]uint) {
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
		panic("Failed to deposit")
	}
}
