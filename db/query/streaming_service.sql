-- name: CreateStreamingService :one
INSERT INTO streaming_services (
    name
) VALUES (
    $1
)
RETURNING *;

-- name: GetStreamingService :one
SELECT * FROM streaming_services
WHERE service_id = $1 LIMIT 1;

-- name: ListStreamingServices :many
SELECT * FROM streaming_services
ORDER BY service_id
LIMIT $1
OFFSET $2;

-- name: UpdateStreamingService :one
UPDATE streaming_services
SET 
    name = COALESCE($2, name)
WHERE service_id = $1
RETURNING *;

-- name: DeleteStreamingService :exec
DELETE FROM streaming_services
WHERE service_id = $1;