package main

import (
	"context"
	"encoding/json"
	"github.com/ScoreTrak/ScoreTrak/pkg/auth"
	"github.com/ScoreTrak/ScoreTrak/pkg/report/reportpb"
	"github.com/ubnetdef/payment-processor/report"
	"google.golang.org/grpc/metadata"
	"log"
)

func getAuth(acli auth.AuthServiceClient) (string, error) {
	resp, err := acli.Login(context.Background(), &auth.LoginRequest{Password: conf.ScoreTrakPassword, Username: conf.ScoreTrakUsername})
	if err != nil {
		return "", err
	}
	return resp.AccessToken, nil
}

func getReport(repCli reportpb.ReportServiceClient, token string) (*report.Report, error) {
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("authorization", token))
	resp, err := repCli.Get(ctx, &reportpb.GetRequest{})
	if err != nil {
		return nil, err
	}
	latestReport := new(report.Report)
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
