package models

import (
	"context"
	"os"

	"cloud.google.com/go/datastore"
)

type ModelClient struct {
	dsClient *datastore.Client
}

func NewClient() (*ModelClient, error) {
	// 参考URL(https://cloud.google.com/datastore/docs/datastore-api-tutorial?hl=ja)
	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, os.Getenv("DATASTORE_PROJECT_ID"))
	if err != nil {
		return nil, err
	}

	client := ModelClient{dsClient: dsClient}
	return &client, nil
}
