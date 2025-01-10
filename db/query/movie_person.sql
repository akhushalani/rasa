-- name: CreateMoviePerson :one
INSERT INTO movie_people (
    movie_id, person_id, role
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: GetPeopleByMovieID :many
SELECT * FROM movie_people
WHERE movie_id = $1
ORDER BY person_id
LIMIT $2
OFFSET $3;

-- name: GetMoviesByPersonID :many
SELECT * FROM movie_people
WHERE person_id = $1
ORDER BY movie_id
LIMIT $2
OFFSET $3;

-- name: DeleteMoviePerson :exec
DELETE FROM movie_people
WHERE movie_id = $1 AND person_id = $2;