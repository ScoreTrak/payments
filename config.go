package main

import "github.com/jinzhu/configor"

type Config struct {
	BankUsername       string `default:""`
	BankPassword       string `default:""`
	ScoretrakUsername  string `default:""`
	ScoretrakPassword  string `default:""`
	BankBaseUrl        string `default:"http://localhost/"`
	ScoretrakAddress   string `default:"localhost"`
	BankAmountPerPoint uint   `default:"10"`
	DBName             string `default:"bank"`
	ClientTimeout      uint   `default:"5"`
}

func NewConfig(f string) (*Config, error) {
	conf := &Config{}
	err := configor.Load(&conf, f)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
