-- name: CreatePerson :one
INSERT INTO people (
    tmdb_id,
    name,
    known_for_department,
    biography,
    birthday,
    deathday,
    gender,
    profile_path,
    tmdb_popularity
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;

-- name: GetPerson :one
SELECT * FROM people
WHERE person_id = $1 LIMIT 1;

-- name: UpdatePerson :one
UPDATE people
SET 
    tmdb_id = COALESCE($2, tmdb_id), 
    name = COALESCE($3, name), 
    known_for_department = COALESCE($4, known_for_department), 
    biography = COALESCE($5, biography), 
    birthday = COALESCE($6, birthday), 
    deathday = COALESCE($7, deathday),
    gender = COALESCE($8, gender),
    profile_path = COALESCE($9, profile_path),
    tmdb_popularity = COALESCE($10, tmdb_popularity),
    last_updated = CURRENT_TIMESTAMP
WHERE person_id = $1
RETURNING *;

-- name: DeletePerson :exec
DELETE FROM people
WHERE person_id = $1;