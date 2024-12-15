package internal

import (
	"context"

	"cloud.google.com/go/bigquery"
	"github.com/bufbuild/protocompile"
)

func Apply(ctx context.Context, projectID, datasetID, filepath string) error {
	cli, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer func() { _ = cli.Close() }()

	compiler := &protocompile.Compiler{
		Resolver: &protocompile.SourceResolver{
			// ImportPaths は外部から指定すべき？
			ImportPaths: []string{
				"./schema/protobuf",
			},
		},
	}
	_, err = compiler.Compile(ctx, filepath)
	if err != nil {
		return err
	}
	return nil
}
