package fs

import (
	"embed"
	"io/fs"
)

var (
	//go:embed public/*
	public embed.FS

	//go:embed migrations/*.sql
	migrations embed.FS
)

func Public() fs.FS {
	return public
}

func Migrations() fs.FS {
	return migrations
}
