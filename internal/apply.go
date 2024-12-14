package internal

import (
	"context"

	"cloud.google.com/go/bigquery"
)

func Apply(ctx context.Context, projectID string) error {
	cli, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer func() { _ = cli.Close() }()

	return nil
}
