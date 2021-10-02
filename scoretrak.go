package main

import (
	"context"
	"encoding/json"
	authpb "github.com/ScoreTrak/ScoreTrak/pkg/proto/auth/v1"
	reportpb "github.com/ScoreTrak/ScoreTrak/pkg/proto/report/v1"
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	"google.golang.org/grpc/metadata"
	"log"
)

func getAuth(acli authpb.AuthServiceClient) (string, error) {
	resp, err := acli.Login(context.Background(), &authpb.LoginRequest{Password: conf.ScoreTrakPassword, Username: conf.ScoreTrakUsername})
	if err != nil {
		return "", err
	}
	return resp.AccessToken, nil
}

func getReport(repCli reportpb.ReportServiceClient, token string) (*report.SimpleReport, error) {
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("authorization", token))
	resp, err := repCli.Get(ctx, &reportpb.GetRequest{})
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
