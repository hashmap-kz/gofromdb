package genpg

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
)

type ColumnInfo struct {
	RelPath    string
	AttName    string
	AttType    string
	RefTo      string
	ColDesc    string
	AttNotNull bool
	TabDesc    string
	AttNum     int32
	CharMaxLen *int32
	NPrec      *int32
	NScal      *int32
	Def        *string
}

func GetDBInfo() map[string][]ColumnInfo {
	// Database connection string (update with your credentials)
	connString := "postgres://postgres:postgres@localhost:5432/bookstore"

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
			&row.AttType,
			&row.RefTo,
			&row.ColDesc,
			&row.AttNotNull,
			&row.TabDesc,
			&row.AttNum,
			&row.CharMaxLen,
			&row.NPrec,
			&row.NScal,
			&row.Def,
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
