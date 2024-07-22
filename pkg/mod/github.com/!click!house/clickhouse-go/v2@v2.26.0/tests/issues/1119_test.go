package issues

import (
	"context"
	"testing"

	"github.com/ClickHouse/clickhouse-go/v2"
	clickhouse_tests "github.com/ClickHouse/clickhouse-go/v2/tests"
	"github.com/stretchr/testify/require"
)

func Test1119(t *testing.T) {
	clickhouse_tests.SkipOnCloud(t, "The JSON data type is an obsolete feature on Cloud.")
	var (
		conn, err = clickhouse_tests.GetConnection("issues", clickhouse.Settings{
			"max_execution_time":             60,
			"allow_experimental_object_type": true,
		}, nil, &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		})
	)
	ctx := context.Background()
	require.NoError(t, err)
	const ddl = "CREATE TABLE test_1119 (col JSON) Engine MergeTree() ORDER BY tuple()"
	require.NoError(t, conn.Exec(ctx, ddl))
	defer func() {
		conn.Exec(ctx, "DROP TABLE IF EXISTS test_1119")
	}()

	batch, err := conn.PrepareBatch(context.Background(), "INSERT INTO test_1119")
	require.NoError(t, err)

	v := map[string]any{
		"str": "foo",
		"int": 12,
	}

	require.NoError(t, batch.Append(v))
	require.NoError(t, batch.Send())
}
