// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: file.sql

package sqlc

import (
	"context"
)

const createFile = `-- name: CreateFile :one
INSERT INTO files (
	"name",
	"access",
	"path",
	"ext"
) VALUES (
    $1, $2, $3, $4
) RETURNING id, name, access, path, "createdAt", "updatedAt", ext
`

type CreateFileParams struct {
	Name   string `json:"name"`
	Access Access `json:"access"`
	Path   string `json:"path"`
	Ext    string `json:"ext"`
}

func (q *Queries) CreateFile(ctx context.Context, arg *CreateFileParams) (*File, error) {
	row := q.db.QueryRow(ctx, createFile,
		arg.Name,
		arg.Access,
		arg.Path,
		arg.Ext,
	)
	var i File
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Access,
		&i.Path,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Ext,
	)
	return &i, err
}

type CreateFilesParams struct {
	Name   string `json:"name"`
	Access Access `json:"access"`
	Path   string `json:"path"`
	Ext    string `json:"ext"`
}

const deleteFilesByNames = `-- name: DeleteFilesByNames :many
DELETE
FROM files
WHERE files."name" = ANY($1::varchar[])
RETURNING id, name, access, path, "createdAt", "updatedAt", ext
`

func (q *Queries) DeleteFilesByNames(ctx context.Context, names []string) ([]*File, error) {
	rows, err := q.db.Query(ctx, deleteFilesByNames, names)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*File{}
	for rows.Next() {
		var i File
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Access,
			&i.Path,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Ext,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFileByName = `-- name: GetFileByName :one
SELECT id, name, access, path, "createdAt", "updatedAt", ext
FROM files
WHERE files."name" = $1
LIMIT 1
`

func (q *Queries) GetFileByName(ctx context.Context, name string) (*File, error) {
	row := q.db.QueryRow(ctx, getFileByName, name)
	var i File
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Access,
		&i.Path,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Ext,
	)
	return &i, err
}

const getFileByNames = `-- name: GetFileByNames :many
SELECT id, name, access, path, "createdAt", "updatedAt", ext
FROM files
WHERE files."name" = ANY($1::varchar[])
`

func (q *Queries) GetFileByNames(ctx context.Context, names []string) ([]*File, error) {
	rows, err := q.db.Query(ctx, getFileByNames, names)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*File{}
	for rows.Next() {
		var i File
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Access,
			&i.Path,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Ext,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateFile = `-- name: UpdateFile :one
UPDATE files
SET
	"access" = COALESCE($1, "access")
WHERE "name" = $2
RETURNING id, name, access, path, "createdAt", "updatedAt", ext
`

type UpdateFileParams struct {
	Access NullAccess `json:"access"`
	Name   string     `json:"name"`
}

func (q *Queries) UpdateFile(ctx context.Context, arg *UpdateFileParams) (*File, error) {
	row := q.db.QueryRow(ctx, updateFile, arg.Access, arg.Name)
	var i File
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Access,
		&i.Path,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Ext,
	)
	return &i, err
}