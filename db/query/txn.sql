-- name: CreateTxn :one
INSERT INTO txn(timestamp, amount, category, description)
    VALUES (?, ?, ?, ?)
RETURNING
    *;

-- name: GetTxn :one
SELECT
    *
FROM
    txn
WHERE
    id = ?;

-- name: ListTxns :many
SELECT
    *
FROM
    txn
WHERE (timestamp > sqlc.narg('start_time')
    OR sqlc.narg('start_time') IS NULL)
AND (timestamp < sqlc.narg('end_time')
    OR sqlc.narg('end_time') IS NULL)
AND (category = sqlc.narg('category')
    OR sqlc.narg('category') IS NULL)
ORDER BY
    timestamp DESC
LIMIT ?1 OFFSET ?2;

-- name: DeleteTxn :one
DELETE FROM txn
WHERE id = ?
RETURNING
    *;

-- name: UpdateTxn :one
UPDATE
    txn
SET
    timestamp = ?,
    amount = ?,
    category = ?,
    description = ?
WHERE
    id = ?
RETURNING
    *;

