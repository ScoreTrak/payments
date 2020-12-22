package main

import (
	"database/sql"
	"flag"
	"github.com/ScoreTrak/ScoreTrak/pkg/auth"
	"github.com/ScoreTrak/ScoreTrak/pkg/report/reportpb"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"time"
)

var conf *Config

func main() {
	var err error
	confPath := flag.String("config", "configs/config.yml", "Please enter a path to config file")
	flag.Parse()
	conf, err = NewConfig(*confPath)
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("sqlite3", conf.DBName)
	if err != nil {
		panic(err) // TODO: Change this?
	}
	defer db.Close()
	initTables(db)
	cc1, err := grpc.Dial(conf.ScoretrakAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}
	acli := auth.NewAuthServiceClient(cc1)
	log.Println("Requesting authentication token")
	t, err := getAuth(acli)
	token, _, err := new(jwt.Parser).ParseUnverified(t, &auth.UserClaims{})
	if err != nil {
		panic(err)
	}
	claims, ok := token.Claims.(*auth.UserClaims)
	if !ok {
		log.Fatalf("invalid token claims")
	}
	log.Println("Authentication token is valid until " + time.Unix(claims.ExpiresAt, 0).String())
	log.Println("Requesting report")

	respcli := reportpb.NewReportServiceClient(cc1)

	r, err := getReport(respcli, t)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Got report for round %d\n", r.Round)

	current := r.PointsPerTeam()
	previous := getMaxPointsPerTeam(db)
	diff := computePointsDiff(current, previous)
	client := http.Client{Timeout: time.Duration(conf.ClientTimeout) * time.Second}
	makeDeposit(client, diff)
	updateMaxPointsPerTeam(db, current)
}
