package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
)

// https://go.dev/doc/faq#guarantee_satisfies_interface
var _ Repository[any] = &repository[any]{}

type Repository[Entity any] interface {
	// Inserts a new item in a single table
	Create(ctx context.Context, query string, queryParams map[string]any) (int64, error)
	// Fetches an item from a single table
	Read(ctx context.Context, query, resourceID string) (*Entity, error)
	// Fetches all items from a single table using limit and offset as parameters for results pagination
	List(ctx context.Context, query string, limit, offset int) ([]*Entity, bool, error)
	// Removes an item from table
	Delete(ctx context.Context, query, resourceID string) error
	// Overwrites the information in a single table row
	Update(ctx context.Context, query string, queryParams map[string]any) error
}

type repository[Entity any] struct {
	connection *sqlx.DB
}

// Retrieves a Persistence object customized for the specific data types
func NewRepository[Entity any](conn *sqlx.DB) (*repository[Entity], error) {
	if conn == nil {
		return nil, fmt.Errorf("a sqlx.DB instance is strictly needed to create a persistence")
	}
	return &repository[Entity]{
		connection: conn,
	}, nil
}

func (p *repository[Entity]) Create(ctx context.Context, query string, params map[string]any) (int64, error) {
	lastInsertId := int64(0)
	rows, err := p.connection.NamedQueryContext(ctx, query, params)
	if err != nil {
		return lastInsertId, fmt.Errorf("insert failed, err: %w", err)
	}
	for rows.Next() {
		if err := rows.Scan(&lastInsertId); err != nil {
			return lastInsertId, fmt.Errorf("insert failed, err: %w", err)
		}
	}
	return lastInsertId, nil
}

func (p *repository[Entity]) Read(ctx context.Context, query, ID string) (*Entity, error) {
	dest := new(Entity)
	err := p.connection.QueryRowx(query, ID).StructScan(dest)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("record with ID %s not found in database", ID)
		}
		return nil, fmt.Errorf("[QueryRowxContext]: %w", err)
	}
	return dest, nil
}

func (p *repository[Entity]) List(ctx context.Context, query string, limit, offset int) ([]*Entity, bool, error) {
	var dest []*Entity
	err := p.connection.SelectContext(ctx, &dest, query, strconv.Itoa(limit+1), strconv.Itoa(offset))
	if err != nil {
		return nil, false, fmt.Errorf("[SelectContext]: %w", err)
	}

	hasNextPage := len(dest) > limit
	bound := min(limit, len(dest))
	list := dest[:int(bound)]

	return list, hasNextPage, nil
}

func (p *repository[Entity]) Delete(ctx context.Context, query, ID string) error {
	_, err := p.connection.ExecContext(ctx, query, ID)
	if err != nil {
		return fmt.Errorf("[ExecContext]: %w", err)
	}
	return nil
}

func (p *repository[Entity]) Update(ctx context.Context, query string, params map[string]any) error {
	_, err := p.connection.NamedExecContext(ctx, query, params)
	if err != nil {
		return fmt.Errorf("[NamedExecContext]: %w", err)
	}
	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
