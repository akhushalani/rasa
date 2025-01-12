package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomMoviePerson(t *testing.T) MoviePeople {
	movie := createRandomMovie(t)
	person := createRandomPerson(t)

	arg := CreateMoviePersonParams{
		MovieID:  movie.MovieID,
		PersonID: person.PersonID,
		Role:     "Actor",
	}

	moviePerson, err := testQueries.CreateMoviePerson(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, moviePerson)

	require.Equal(t, movie.MovieID, moviePerson.MovieID)
	require.Equal(t, person.PersonID, moviePerson.PersonID)
	require.Equal(t, arg.Role, moviePerson.Role)

	return moviePerson
}

func TestCreateMoviePerson(t *testing.T) {
	moviePerson := createRandomMoviePerson(t)
	require.NotEmpty(t, moviePerson)
}

func TestGetMoviesByPersonID(t *testing.T) {
	moviePerson := createRandomMoviePerson(t)

	arg := GetMoviesByPersonIDParams{
		PersonID: moviePerson.PersonID,
		Limit:    10,
		Offset:   0,
	}

	movies, err := testQueries.GetMoviesByPersonID(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, movies)

	require.Equal(t, moviePerson.MovieID, movies[0].MovieID)
	require.Equal(t, moviePerson.PersonID, movies[0].PersonID)
	require.Equal(t, moviePerson.Role, movies[0].Role)
}

func TestGetPeopleByMovieID(t *testing.T) {
	moviePerson := createRandomMoviePerson(t)

	arg := GetPeopleByMovieIDParams{
		MovieID: moviePerson.MovieID,
		Limit:   10,
		Offset:  0,
	}

	people, err := testQueries.GetPeopleByMovieID(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, people)

	require.Equal(t, moviePerson.MovieID, people[0].MovieID)
	require.Equal(t, moviePerson.PersonID, people[0].PersonID)
	require.Equal(t, moviePerson.Role, people[0].Role)
}

func TestDeleteMoviePerson(t *testing.T) {
	moviePerson := createRandomMoviePerson(t)

	arg := DeleteMoviePersonParams{
		MovieID:  moviePerson.MovieID,
		PersonID: moviePerson.PersonID,
	}

	err := testQueries.DeleteMoviePerson(context.Background(), arg)
	require.NoError(t, err)

	// Verify the movie-person association is deleted
	argGet := GetMoviesByPersonIDParams{
		PersonID: moviePerson.PersonID,
		Limit:    10,
		Offset:   0,
	}
	movies, err := testQueries.GetMoviesByPersonID(context.Background(), argGet)
	require.NoError(t, err)
	require.Empty(t, movies)
}

func TestCascadeDeleteMoviePersonOnMovie(t *testing.T) {
	moviePerson := createRandomMoviePerson(t)

	// Delete the movie
	err := testQueries.DeleteMovie(context.Background(), moviePerson.MovieID)
	require.NoError(t, err)

	// Verify the movie-person association is deleted
	arg := GetMoviesByPersonIDParams{
		PersonID: moviePerson.PersonID,
		Limit:    10,
		Offset:   0,
	}
	movies, err := testQueries.GetMoviesByPersonID(context.Background(), arg)
	require.NoError(t, err)
	require.Empty(t, movies)
}

func TestCascadeDeleteMoviePersonOnPerson(t *testing.T) {
	moviePerson := createRandomMoviePerson(t)

	// Delete the person
	err := testQueries.DeletePerson(context.Background(), moviePerson.PersonID)
	require.NoError(t, err)

	// Verify the movie-person association is deleted
	arg := GetPeopleByMovieIDParams{
		MovieID: moviePerson.MovieID,
		Limit:   10,
		Offset:  0,
	}
	people, err := testQueries.GetPeopleByMovieID(context.Background(), arg)
	require.NoError(t, err)
	require.Empty(t, people)
}
