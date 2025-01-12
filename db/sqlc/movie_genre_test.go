package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomMovieGenre(t *testing.T) MovieGenres {
	// Create random movie and genre
	movie := createRandomMovie(t)
	genre := createRandomGenre(t)

	// Create association
	arg := CreateMovieGenreParams{
		MovieID: movie.MovieID,
		GenreID: genre.GenreID,
	}

	movieGenre, err := testQueries.CreateMovieGenre(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, movieGenre)

	require.Equal(t, movie.MovieID, movieGenre.MovieID)
	require.Equal(t, genre.GenreID, movieGenre.GenreID)

	return movieGenre
}

func TestCreateMovieGenre(t *testing.T) {
	movieGenre := createRandomMovieGenre(t)
	require.NotEmpty(t, movieGenre)
}

func TestGetMovieGenre(t *testing.T) {
	// Create a random movie-genre association
	movieGenre := createRandomMovieGenre(t)

	// Fetch the association by movie_id and genre_id
	arg := GetMovieGenreParams{
		MovieID: movieGenre.MovieID,
		GenreID: movieGenre.GenreID,
	}
	fetchedMovieGenre, err := testQueries.GetMovieGenre(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedMovieGenre)

	// Validate the fetched association
	require.Equal(t, movieGenre.MovieID, fetchedMovieGenre.MovieID)
	require.Equal(t, movieGenre.GenreID, fetchedMovieGenre.GenreID)
}

func TestDeleteMovieGenre(t *testing.T) {
	// Create a random movie-genre association
	movieGenre := createRandomMovieGenre(t)

	// Delete the association
	arg := DeleteMovieGenreParams{
		MovieID: movieGenre.MovieID,
		GenreID: movieGenre.GenreID,
	}
	err := testQueries.DeleteMovieGenre(context.Background(), arg)
	require.NoError(t, err)

	// Verify the association is deleted
	fetchedMovieGenre, err := testQueries.GetMovieGenre(context.Background(), GetMovieGenreParams{
		MovieID: movieGenre.MovieID,
		GenreID: movieGenre.GenreID,
	})
	require.Error(t, err)
	require.Empty(t, fetchedMovieGenre)
}

func TestCascadeDeleteMovieGenreOnMovie(t *testing.T) {
	movieGenre := createRandomMovieGenre(t)

	// Delete the movie
	err := testQueries.DeleteMovie(context.Background(), movieGenre.MovieID)
	require.NoError(t, err)

	// Verify the cascade deletion from movie_genres
	fetchedMovieGenre, err := testQueries.GetMovieGenre(context.Background(), GetMovieGenreParams{
		MovieID: movieGenre.MovieID,
		GenreID: movieGenre.GenreID,
	})
	require.Error(t, err)
	require.Empty(t, fetchedMovieGenre)
}

func TestCascadeDeleteMovieGenreOnGenre(t *testing.T) {
	movieGenre := createRandomMovieGenre(t)

	// Delete the genre
	err := testQueries.DeleteGenre(context.Background(), movieGenre.GenreID)
	require.NoError(t, err)

	// Verify the cascade deletion from movie_genres
	fetchedMovieGenre, err := testQueries.GetMovieGenre(context.Background(), GetMovieGenreParams{
		MovieID: movieGenre.MovieID,
		GenreID: movieGenre.GenreID,
	})
	require.Error(t, err)
	require.Empty(t, fetchedMovieGenre)
}


func TestListMovieGenres(t *testing.T) {
	// Create multiple movie-genre associations
	for i := 0; i < 10; i++ {
		createRandomMovieGenre(t)
	}

	// List associations with limit and offset
	arg := ListMovieGenresParams{
		Limit:  5,
		Offset: 0,
	}
	movieGenres, err := testQueries.ListMovieGenres(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, movieGenres, int(arg.Limit))

	// Validate each association
	for _, movieGenre := range movieGenres {
		require.NotEmpty(t, movieGenre)
	}
}
