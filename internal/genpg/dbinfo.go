package genpg

import (
	"context"
	"log"

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

func GetDBInfo(connString string) map[string][]ColumnInfo {
	// Connect to the database
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	// Execute the query
	rows, err := conn.Query(context.Background(), GetInfoQuery)
	if err != nil {
		log.Fatalf("Query execution failed: %v\n", err)
	}
	defer rows.Close()

	colInfo := map[string][]ColumnInfo{}

	// Iterate through the rows and print the data
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
			log.Fatalf("Failed to scan row: %v\n", err)
		}
		colInfo[row.RelPath] = append(colInfo[row.RelPath], row)
	}

	// Check for errors encountered during iteration
	if rows.Err() != nil {
		log.Fatalf("Row iteration error: %v\n", rows.Err())
	}

	return colInfo
}
