package db

import (
	"context"
	"database/sql"
	"sync"

	"github.com/gkits/kurz/internal/types"
	"github.com/jmoiron/sqlx"
)

var (
	db   *sqlx.DB
	once sync.Once
)

func Init(conn *sql.DB, driver string) {
	once.Do(func() {
		db = sqlx.NewDb(conn, driver)
	})
}

func GetLinkByID(ctx context.Context, id int64) (types.Link, error) {
	const query = `SELECT * FROM links WHERE id = $1;`

	var link types.Link
	if err := db.GetContext(ctx, &link, query, id); err != nil {
		return types.Link{}, err
	}
	return link, nil
}

func GetLinkByRef(ctx context.Context, ref string) (types.Link, error) {
	const query = `SELECT * FROM links WHERE ref = $1;`

	var link types.Link
	if err := db.GetContext(ctx, &link, query, ref); err != nil {
		return types.Link{}, err
	}
	return link, nil
}

func GetLinks(ctx context.Context, q types.LinkQuery) ([]types.Link, error) {
	const query = `SELECT * FROM links;`

	rows, err := db.NamedQueryContext(ctx, query, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var links []types.Link
	for rows.Next() {
		var link types.Link
		if err := rows.StructScan(&link); err != nil {
			return nil, err
		}
		links = append(links, link)
	}

	return links, nil
}

func InsertLink(ctx context.Context, link types.Link) (types.Link, error) {
	const query = `
		INSERT INTO links (ref, target, created_by, created_at, expires_at)
		VALUES ($1, $2, $3, $4, $5);`

	res, err := db.ExecContext(ctx, query, link.Ref, link.Target, link.CreatedBy, link.CreatedAt, link.ExpiresAt)
	if err != nil {
		return types.Link{}, err
	}
	link.ID, err = res.LastInsertId()
	if err != nil {
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
