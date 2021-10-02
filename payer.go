package main

import (
	"flag"
	"github.com/ScoreTrak/ScoreTrak/pkg/auth"
	authpb "github.com/ScoreTrak/ScoreTrak/pkg/proto/auth/v1"
	reportpb "github.com/ScoreTrak/ScoreTrak/pkg/proto/report/v1"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage"
	"github.com/ScoreTrak/lockdown-payments/report"
	"github.com/golang-jwt/jwt/v4"
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
	db, err := storage.NewDB(conf.DB)
	if err != nil {
		panic(err)
	}
	err = initTables(db)
	if err != nil {
		panic(err)
	}
	cc, err := grpc.Dial(conf.ScoreTrakAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}
	authClient := authpb.NewAuthServiceClient(cc)
	log.Println("Requesting authentication token")
	t, err := getAuth(authClient)
	if err != nil {
		panic(err)
	}
	token, _, err := new(jwt.Parser).ParseUnverified(t, &auth.UserClaims{})
	if err != nil {
		panic(err)
	}
	claims, ok := token.Claims.(*auth.UserClaims)
	if !ok {
		log.Fatalf("invalid token claims")
	}
	log.Println("Authentication token is valid until " + claims.ExpiresAt.Time.String())
	log.Println("Requesting report")
	reportClient := reportpb.NewReportServiceClient(cc)
	latestReport, err := getReport(reportClient, t)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Got report for round %d\n", latestReport.Round)
	current := report.PointsPerTeam(latestReport)
	previous := getMaxPointsPerTeam(db)
	diff := computePointsDiff(current, previous)
	client := http.Client{Timeout: time.Duration(conf.ClientTimeout) * time.Second}
	makeDeposit(client, diff)
	updateMaxPointsPerTeam(db, current)
}
