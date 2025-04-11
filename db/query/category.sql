-- name: CreateCategory :one
INSERT INTO category(name)
    VALUES (?)
RETURNING
    *;

-- name: ListCategories :many
SELECT
    *
FROM
    category
WHERE (name = sqlc.narg('name')
    OR name LIKE '%' || sqlc.narg('name') || '%'
    OR sqlc.narg('name') IS NULL)
ORDER BY
    name
