package book_translations

import (
	"context"
	"fmt"
	"go-project-template-v7/pkg/pageable"
	"go-project-template-v7/pkg/storage/postgres"

	"github.com/jackc/pgx/v5"
)

type Repository interface {
	Save(ctx context.Context, inputEntity *BookTranslations) (*BookTranslations, error)
	UpdateByID(ctx context.Context, inputEntity *BookTranslations, pkBookID int64, pkLanguageCode string) (*BookTranslations, error)
	DeleteByID(ctx context.Context, pkBookID int64, pkLanguageCode string) error
	FindByID(ctx context.Context, pkBookID int64, pkLanguageCode string) (*BookTranslations, error)
	FindAll(ctx context.Context) ([]BookTranslations, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]BookTranslations, pageable.Page, error)
}

type repo struct {
	db *postgres.Postgres
}

var _ Repository = &repo{}

func NewRepository(_ context.Context, db *postgres.Postgres) Repository {
	return &repo{
		db: db,
	}
}

func (r *repo) Save(ctx context.Context, inputEntity *BookTranslations) (*BookTranslations, error) {
	tag := "repository.Save"

	query := `		
		insert into bookstore_catalog.book_translations (
			book_id,
			language_code,
			translated_title,
			translated_by,
			released_on
		)
		values ($1, $2, $3, $4, $5)
		returning
			book_id,
			language_code,
			translated_title,
			translated_by,
			released_on
		`

	row := r.db.Pool.QueryRow(
		ctx, query,
		inputEntity.BookID,
		inputEntity.LanguageCode,
		inputEntity.TranslatedTitle,
		inputEntity.TranslatedBy,
		inputEntity.ReleasedOn,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *repo) UpdateByID(ctx context.Context, inputEntity *BookTranslations, pkBookID int64, pkLanguageCode string) (*BookTranslations, error) {
	tag := "repository.UpdateByID"

	query := `		
		update bookstore_catalog.book_translations
		set
			translated_title = coalesce(nullif($3, ''), translated_title),
			translated_by    = coalesce(nullif($4, ''), translated_by),
			released_on      = coalesce(nullif($5, '0001-01-01 00:00:00'::date), released_on)
		where book_id = $1 and language_code = $2
		returning
			book_id,
			language_code,
			translated_title,
			translated_by,
			released_on
		`

	row := r.db.Pool.QueryRow(
		ctx, query,
		pkBookID, pkLanguageCode,
		inputEntity.TranslatedTitle,
		inputEntity.TranslatedBy,
		inputEntity.ReleasedOn,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *repo) DeleteByID(ctx context.Context, pkBookID int64, pkLanguageCode string) error {
	tag := "repository.DeleteByID"

	query := `		
		delete from only bookstore_catalog.book_translations
		where book_id = $1 and language_code = $2
		`

	cmdTag, err := r.db.Pool.Exec(ctx, query, pkBookID, pkLanguageCode)
	if err != nil || cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%s. no rows were deleted: %w", tag, err)
	}
	return nil
}

func (r *repo) FindByID(ctx context.Context, pkBookID int64, pkLanguageCode string) (*BookTranslations, error) {
	tag := "repository.FindByID"

	query := `		
		select
			book_id,
			language_code,
			translated_title,
			translated_by,
			released_on
		from bookstore_catalog.book_translations
		where book_id = $1 and language_code = $2
		order by book_id, language_code
		`

	row := r.db.Pool.QueryRow(ctx, query, pkBookID, pkLanguageCode)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *repo) FindAll(ctx context.Context) ([]BookTranslations, error) {
	tag := "repository.FindAll"

	query := `		
		select
			book_id,
			language_code,
			translated_title,
			translated_by,
			released_on
		from bookstore_catalog.book_translations
		order by book_id, language_code
		`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []BookTranslations
	for rows.Next() {
		scannedEntity, err := scanFullRow(rows)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", tag, err)
		}
		scannedEntities = append(scannedEntities, *scannedEntity)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return scannedEntities, nil
}

func (r *repo) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]BookTranslations, pageable.Page, error) {
	tag := "repository.FindAllPageable"

	// retrieve total count
	queryCnt := `select count(*) from bookstore_catalog.book_translations`
	var totalCount int
	if err := r.db.Pool.QueryRow(ctx, queryCnt).Scan(&totalCount); err != nil {
		return nil, pageable.Page{}, err
	}

	// init page
	page := pageable.CreatePage(pq, totalCount)

	// handle empty result
	if totalCount == 0 {
		return nil, page, nil
	}

	// select entities
	query := `		
		select
			book_id,
			language_code,
			translated_title,
			translated_by,
			released_on
		from bookstore_catalog.book_translations
		order by book_id, language_code
		offset $1 limit $2
		`

	rows, err := r.db.Pool.Query(ctx, query, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	var scannedEntities []BookTranslations
	for rows.Next() {
		scannedEntity, err := scanFullRow(rows)
		if err != nil {
			return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
		}
		scannedEntities = append(scannedEntities, *scannedEntity)
	}

	if rows.Err() != nil {
		return nil, page, rows.Err()
	}
	return scannedEntities, page, nil
}

// scan utils

// scanFullRow is expected to scan all columns from a table.
// For simplicity, most methods scan the entire row of the table into the result entity.
// You should adapt methods as needed (e.g., if business logic requires returning only an ID after an UPDATE).
func scanFullRow(row pgx.Row) (*BookTranslations, error) {
	var scannedEntity BookTranslations
	err := row.Scan(
		&scannedEntity.BookID,
		&scannedEntity.LanguageCode,
		&scannedEntity.TranslatedTitle,
		&scannedEntity.TranslatedBy,
		&scannedEntity.ReleasedOn,
	)
	if err != nil {
		return nil, err
	}
	return &scannedEntity, nil
}
