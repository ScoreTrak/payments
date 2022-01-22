package main

import (
	"context"
	"encoding/json"
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	authv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/auth/v1"
	reportv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/report/v1"
	"google.golang.org/grpc/metadata"
	"log"
)

func getAuth(acli authv1.AuthServiceClient) (string, error) {
	resp, err := acli.Login(context.Background(), &authv1.LoginRequest{Password: conf.ScoreTrakPassword, Username: conf.ScoreTrakUsername})
	if err != nil {
		return "", err
	}
	return resp.AccessToken, nil
}

func getReport(repCli reportv1.ReportServiceClient, token string) (*report.SimpleReport, error) {
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("authorization", token))
	resp, err := repCli.Get(ctx, &reportv1.GetRequest{})
	if err != nil {
		return nil, err
	}
	latestReport := new(report.SimpleReport)
	res, err := resp.Recv()
	if err != nil {
		log.Fatal("stream was closed: ", err)
	}
	defer resp.Context().Done()
	err = json.Unmarshal([]byte(res.Report.Cache), latestReport)
	if err != nil {
		return nil, err
	}
	return latestReport, nil
}
