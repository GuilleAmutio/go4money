-- name: CreateTransfer :one
INSERT INTO transfers (
  from_account_id, 
  to_account_id, 
  amount
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: ListAllTransfers :many
SELECT * FROM transfers;

-- name: ListTransfers :many
SELECT * FROM transfers
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: ListTransfersFromAccount :many
SELECT * FROM transfers
WHERE from_account_id = $1
ORDER BY id;

-- name: ListTransfersToAccount :many
SELECT * FROM transfers
WHERE to_account_id = $1
ORDER BY id;

-- name: ListTransfersBetweenDates :many
SELECT * FROM transfers
WHERE created_at >= $1 AND created_at <= $2;