package protobq

import (
	"context"
	"errors"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/googleapi"
)

func ApplySchema(ctx context.Context, projectID string, datasetID string, views []MaterializedView) error {
	cli, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer func() { _ = cli.Close() }()

	dataset := cli.Dataset(datasetID)

	for _, mv := range views {
		exists, err := existsMaterializedView(ctx, dataset, mv)
		if err != nil {
			return err
		}
		if exists {
			err = dataset.Table(mv.Name()).Delete(ctx)
			if err != nil {
				return err
			}
		}

		md := &bigquery.TableMetadata{
			Name: mv.Name(),
			MaterializedView: &bigquery.MaterializedViewDefinition{
				EnableRefresh:   mv.Options().EnableRefresh,
				RefreshInterval: mv.Options().RefreshInterval,
			},
		}
		table := dataset.Table(mv.Name())
		err = table.Create(ctx, md)
		if err != nil {
			return err
		}
	}
	return nil
}

func existsMaterializedView(ctx context.Context, dataset *bigquery.Dataset, mv MaterializedView) (bool, error) {
	ref := dataset.Table(mv.Name())
	_, err := ref.Metadata(ctx)
	if err != nil {
		var e *googleapi.Error
		if errors.As(err, &e) && e.Code == 404 {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
