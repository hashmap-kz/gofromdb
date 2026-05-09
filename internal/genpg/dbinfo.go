package genpg

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
)

type ColumnInfo struct {
	RelPath      string
	AttName      string
	AttType2     string
	ColDesc      string
	AttNotNull   bool
	TabDesc      string
	AttNum       int32
	PrimaryKeys  []string
	GoType       string
	NullifRhs    string
	IsInsertable bool
}

func GetDBInfo(connString string) (map[string][]ColumnInfo, error) {
	slog.Debug("connecting to database")
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("connect: %w", err)
	}
	defer conn.Close(context.Background())

	slog.Debug("querying column info")
	rows, err := conn.Query(context.Background(), GetInfoQuery)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	defer rows.Close()

	colInfo := map[string][]ColumnInfo{}

	for rows.Next() {
		var row ColumnInfo
		if err := rows.Scan(
			&row.RelPath,
			&row.AttName,
			&row.AttType2,
			&row.ColDesc,
			&row.AttNotNull,
			&row.TabDesc,
			&row.AttNum,
			&row.PrimaryKeys,
			&row.GoType,
			&row.NullifRhs,
			&row.IsInsertable,
		); err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}
		colInfo[row.RelPath] = append(colInfo[row.RelPath], row)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("rows: %w", rows.Err())
	}

	slog.Debug("db introspection complete", slog.Int("tables", len(colInfo)))
	return colInfo, nil
}
