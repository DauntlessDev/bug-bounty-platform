-- name: GetBounties :many
SELECT id, title, description, points FROM bounties;

-- name: GetBountyByID :one
SELECT id, title, description, points FROM bounties WHERE id = $1;

-- name: CreateBounty :exec
INSERT INTO bounties (id, title, description, points) VALUES ($1, $2, $3, $4);

-- name: UpdateBounty :exec
UPDATE bounties SET title = $1, description = $2, points = $3 WHERE id = $4;