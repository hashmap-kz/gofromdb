package books

import (
	"context"
	"errors"
	"fmt"
	"go-project-template-v7/pkg/apperrors"
	"go-project-template-v7/pkg/pageable"
	"go-project-template-v7/pkg/storage/postgres"

	"github.com/jackc/pgx/v5"
)

type Repository interface {
	Save(ctx context.Context, inputEntity *Books) (*Books, error)
	UpdateByID(ctx context.Context, inputEntity *Books, pkBookID int64) (*Books, error)
	DeleteByID(ctx context.Context, pkBookID int64) error
	FindByID(ctx context.Context, pkBookID int64) (*Books, error)
	FindAll(ctx context.Context) ([]Books, error)
	FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]Books, pageable.Page, error)
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

func (r *repo) Save(ctx context.Context, inputEntity *Books) (*Books, error) {
	tag := "repository.Save"

	query := `		
		insert into bookstore_catalog.books (
			publisher_code,
			isbn13,
			title,
			subtitle,
			description,
			price,
			weight_grams,
			rating,
			published_on,
			tags,
			attrs,
			cover_image,
			archived_at
		)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		returning
			book_id,
			publisher_code,
			isbn13,
			title,
			subtitle,
			description,
			price,
			weight_grams,
			rating,
			published_on,
			tags,
			attrs,
			cover_image,
			archived_at,
			title_search
		`

	row := r.db.Pool.QueryRow(
		ctx, query,
		inputEntity.PublisherCode,
		inputEntity.Isbn13,
		inputEntity.Title,
		inputEntity.Subtitle,
		inputEntity.Description,
		inputEntity.Price,
		inputEntity.WeightGrams,
		inputEntity.Rating,
		inputEntity.PublishedOn,
		inputEntity.Tags,
		inputEntity.Attrs,
		inputEntity.CoverImage,
		inputEntity.ArchivedAt,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *repo) UpdateByID(ctx context.Context, inputEntity *Books, pkBookID int64) (*Books, error) {
	tag := "repository.UpdateByID"

	query := `		
		update bookstore_catalog.books
		set
			publisher_code = coalesce(nullif($2, ''), publisher_code),
			isbn13         = coalesce(nullif($3, ''), isbn13),
			title          = coalesce(nullif($4, ''), title),
			subtitle       = coalesce(nullif($5, ''), subtitle),
			description    = coalesce(nullif($6, ''), description),
			price          = coalesce(nullif($7, 0::numeric), price),
			weight_grams   = coalesce(nullif($8, 0::int4), weight_grams),
			rating         = coalesce(nullif($9, 0::numeric), rating),
			published_on   = coalesce(nullif($10, '0001-01-01 00:00:00'::date), published_on),
			tags           = coalesce(nullif($11, '{}'::text[]), tags),
			attrs          = coalesce(nullif($12, '{}'::jsonb), attrs),
			cover_image    = coalesce(nullif($13, ''::bytea), cover_image),
			archived_at    = coalesce(nullif($14, '0001-01-01 00:00:00'::timestamptz), archived_at)
		where book_id = $1
		returning
			book_id,
			publisher_code,
			isbn13,
			title,
			subtitle,
			description,
			price,
			weight_grams,
			rating,
			published_on,
			tags,
			attrs,
			cover_image,
			archived_at,
			title_search
		`

	row := r.db.Pool.QueryRow(
		ctx, query,
		pkBookID,
		inputEntity.PublisherCode,
		inputEntity.Isbn13,
		inputEntity.Title,
		inputEntity.Subtitle,
		inputEntity.Description,
		inputEntity.Price,
		inputEntity.WeightGrams,
		inputEntity.Rating,
		inputEntity.PublishedOn,
		inputEntity.Tags,
		inputEntity.Attrs,
		inputEntity.CoverImage,
		inputEntity.ArchivedAt,
	)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *repo) DeleteByID(ctx context.Context, pkBookID int64) error {
	tag := "repository.DeleteByID"

	query := `		
		delete from only bookstore_catalog.books
		where book_id = $1
		`

	cmdTag, err := r.db.Pool.Exec(ctx, query, pkBookID)
	if err != nil {
		return fmt.Errorf("%s: %w", tag, err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("%s: %w", tag, apperrors.ErrNotFound)
	}
	return nil
}

func (r *repo) FindByID(ctx context.Context, pkBookID int64) (*Books, error) {
	tag := "repository.FindByID"

	query := `		
		select
			book_id,
			publisher_code,
			isbn13,
			title,
			subtitle,
			description,
			price,
			weight_grams,
			rating,
			published_on,
			tags,
			attrs,
			cover_image,
			archived_at,
			title_search
		from bookstore_catalog.books
		where book_id = $1
		order by book_id
		`

	row := r.db.Pool.QueryRow(ctx, query, pkBookID)

	scannedEntity, err := scanFullRow(row)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	return scannedEntity, nil
}

func (r *repo) FindAll(ctx context.Context) ([]Books, error) {
	tag := "repository.FindAll"

	query := `		
		select
			book_id,
			publisher_code,
			isbn13,
			title,
			subtitle,
			description,
			price,
			weight_grams,
			rating,
			published_on,
			tags,
			attrs,
			cover_image,
			archived_at,
			title_search
		from bookstore_catalog.books
		order by book_id
		`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	scannedEntities := make([]Books, 0)
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

func (r *repo) FindAllPageable(ctx context.Context, pq *pageable.PaginationQuery) ([]Books, pageable.Page, error) {
	tag := "repository.FindAllPageable"

	// retrieve total count
	queryCnt := `select count(*) from bookstore_catalog.books`
	var totalCount int
	if err := r.db.Pool.QueryRow(ctx, queryCnt).Scan(&totalCount); err != nil {
		return nil, pageable.Page{}, err
	}

	// init page
	page := pageable.CreatePage(pq, totalCount)

	// handle empty result
	if totalCount == 0 {
		return make([]Books, 0), page, nil
	}

	// select entities
	query := `		
		select
			book_id,
			publisher_code,
			isbn13,
			title,
			subtitle,
			description,
			price,
			weight_grams,
			rating,
			published_on,
			tags,
			attrs,
			cover_image,
			archived_at,
			title_search
		from bookstore_catalog.books
		order by book_id
		offset $1 limit $2
		`

	rows, err := r.db.Pool.Query(ctx, query, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, pageable.Page{}, fmt.Errorf("%s: %w", tag, err)
	}
	defer rows.Close()

	scannedEntities := make([]Books, 0)
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
func scanFullRow(row pgx.Row) (*Books, error) {
	var scannedEntity Books
	err := row.Scan(
		&scannedEntity.BookID,
		&scannedEntity.PublisherCode,
		&scannedEntity.Isbn13,
		&scannedEntity.Title,
		&scannedEntity.Subtitle,
		&scannedEntity.Description,
		&scannedEntity.Price,
		&scannedEntity.WeightGrams,
		&scannedEntity.Rating,
		&scannedEntity.PublishedOn,
		&scannedEntity.Tags,
		&scannedEntity.Attrs,
		&scannedEntity.CoverImage,
		&scannedEntity.ArchivedAt,
		&scannedEntity.TitleSearch,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.ErrNotFound
		}
		return nil, err
	}
	return &scannedEntity, nil
}
