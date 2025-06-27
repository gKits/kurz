package db

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/gkits/kurz/internal/types"
	"github.com/jmoiron/sqlx"
)

var (
	db   *sqlx.DB
	once sync.Once
)

func Init(connStr string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if recoveredErr, ok := r.(error); ok {
				err = fmt.Errorf("db: failed to initialize: %w", recoveredErr)
				return
			}
			err = errors.New("db: failed to initialize")
		}
	}()
	once.Do(func() {
		conn, err := sqlx.Open("sqlite", connStr)
		if err != nil {
			panic(err)
		}
		db = conn
	})
	return nil
}

func GetLinkByID(ctx context.Context, id int64) (types.Link, error) {
	const query = `SELECT * FROM links WHERE id = $1`

	var link types.Link
	if err := db.GetContext(ctx, &link, query, id); err != nil {
		return types.Link{}, err
	}
	return link, nil
}

func GetLinkByRef(ctx context.Context, ref string) (types.Link, error) {
	const query = `SELECT * FROM links WHERE ref = $1`

	var link types.Link
	if err := db.GetContext(ctx, &link, query, ref); err != nil {
		return types.Link{}, err
	}
	return link, nil
}

func GetLinks(ctx context.Context, q types.LinkQuery) ([]types.Link, error) {
	const query = `SELECT * FROM links WHERE id = $1`

	var links []types.Link
	if err := db.SelectContext(ctx, &links, query); err != nil {
		return nil, err
	}
	return links, nil
}

func InsertLink(ctx context.Context, link types.Link) (types.Link, error) {
	const query = `INSERT INTO links VALUES (ref, target, created_by, created_at, expires_at)`

	rows, err := db.NamedQueryContext(ctx, query, link)
	if err != nil {
		return types.Link{}, err
	}
	defer func() { _ = rows.Close() }()

	rows.Next()
	if err := rows.Scan(link.ID); err != nil {
		return types.Link{}, err
	}
	return link, nil
}

func DeleteLink(ctx context.Context, id int64) error {
	const query = `DELETE FROM links WHERE id = $1;`

	if _, err := db.ExecContext(ctx, query, id); err != nil {
		return err
	}
	return nil
}
