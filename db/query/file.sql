-- name: CreateFile :one
INSERT INTO files (
	"name",
	"access",
	"path",
	"ext"
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetFileByName :one
SELECT *
FROM files
WHERE files."name" = $1
LIMIT 1;

-- name: GetFileByNames :many
SELECT *
FROM files
WHERE files."name" = ANY(sqlc.arg(names)::varchar[]);

-- name: CreateFiles :copyfrom
INSERT INTO files (
	"name",
	"access",
	"path",
	"ext"
) VALUES (
    $1, $2, $3, $4
);

-- name: UpdateFile :one
UPDATE files
SET
	"access" = COALESCE(sqlc.narg(access), "access")
WHERE "name" = sqlc.arg(name)
RETURNING *;

-- name: DeleteFilesByNames :many
DELETE
FROM files
WHERE files."name" = ANY(sqlc.arg(names)::varchar[])
RETURNING *;
