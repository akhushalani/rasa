package db

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/akhushalani/rasa/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func createRandomMovie(t *testing.T) Movies {
	uniqueTmdbID := int32(util.RandomInt(0, 1000000))                  // Generate a random unique TmdbID
	uniqueImdbID := fmt.Sprintf("tt%07d", util.RandomInt(0, 10000000)) // Generate a random unique ImdbID

	arg := CreateMovieParams{
		TmdbID:         uniqueTmdbID,
		ImdbID:         pgtype.Text{String: uniqueImdbID, Valid: true},
		Title:          "Test Movie",
		Overview:       pgtype.Text{String: "A test movie overview.", Valid: true},
		ReleaseDate:    pgtype.Date{Time: time.Now(), Valid: true},
		PosterPath:     pgtype.Text{String: "/poster/path.jpg", Valid: true},
		BackdropPath:   pgtype.Text{String: "/backdrop/path.jpg", Valid: true},
		TmdbPopularity: decimal.NewFromInt(90),
	}

	movie, err := testQueries.CreateMovie(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, movie)

	require.Equal(t, arg.TmdbID, movie.TmdbID)
	require.Equal(t, arg.ImdbID.String, movie.ImdbID.String)
	require.Equal(t, arg.Title, movie.Title)

	return movie
}

func TestGetMovie(t *testing.T) {
	movie := createRandomMovie(t)

	fetchedMovie, err := testQueries.GetMovie(context.Background(), movie.MovieID)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedMovie)

	require.Equal(t, movie.MovieID, fetchedMovie.MovieID)
	require.Equal(t, movie.TmdbID, fetchedMovie.TmdbID)
	require.Equal(t, movie.Title, fetchedMovie.Title)
}

func TestGetMovieByTmdbID(t *testing.T) {
	movie := createRandomMovie(t)

	fetchedMovie, err := testQueries.GetMovieByTmdbId(context.Background(), movie.TmdbID)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedMovie)

	require.Equal(t, movie.MovieID, fetchedMovie.MovieID)
	require.Equal(t, movie.TmdbID, fetchedMovie.TmdbID)
	require.Equal(t, movie.Title, fetchedMovie.Title)
}

func TestUpdateMovie(t *testing.T) {
	movie := createRandomMovie(t)

	arg := UpdateMovieParams{
		MovieID:        movie.MovieID,
		TmdbID:         movie.TmdbID,
		ImdbID:         pgtype.Text{String: fmt.Sprintf("tt%07d", util.RandomInt(0, 10000000)), Valid: true},
		Title:          "Updated Test Movie",
		Overview:       pgtype.Text{String: "An updated overview.", Valid: true},
		ReleaseDate:    pgtype.Date{Time: time.Now(), Valid: true},
		PosterPath:     pgtype.Text{String: "/updated/poster.jpg", Valid: true},
		BackdropPath:   pgtype.Text{String: "/updated/backdrop.jpg", Valid: true},
		TmdbPopularity: decimal.NewFromInt(95),
	}

	updatedMovie, err := testQueries.UpdateMovie(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedMovie)

	require.Equal(t, arg.Title, updatedMovie.Title)
	require.Equal(t, arg.ImdbID.String, updatedMovie.ImdbID.String)
}

func TestDeleteMovie(t *testing.T) {
	movie := createRandomMovie(t)

	err := testQueries.DeleteMovie(context.Background(), movie.MovieID)
	require.NoError(t, err)

	// Verify the movie is deleted
	deletedMovie, err := testQueries.GetMovie(context.Background(), movie.MovieID)
	require.Error(t, err)
	require.Empty(t, deletedMovie)
}

func TestListMovies(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomMovie(t)
	}

	arg := ListMoviesParams{
		Limit:  5,
		Offset: 0,
	}

	movies, err := testQueries.ListMovies(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, movies, int(arg.Limit))

	for _, movie := range movies {
		require.NotEmpty(t, movie)
	}
}

func TestCreateMovieUniqueConstraint(t *testing.T) {
	movie := createRandomMovie(t)

	// Attempt to create a duplicate movie with the same tmdb_id
	arg := CreateMovieParams{
		TmdbID:         movie.TmdbID,                                          // Same TmdbID
		ImdbID:         pgtype.Text{String: movie.ImdbID.String, Valid: true}, // Same ImdbID
		Title:          "Duplicate Movie",
		Overview:       pgtype.Text{String: "A duplicate movie overview.", Valid: true},
		ReleaseDate:    pgtype.Date{Time: time.Now(), Valid: true},
		PosterPath:     pgtype.Text{String: "/poster/path.jpg", Valid: true},
		BackdropPath:   pgtype.Text{String: "/backdrop/path.jpg", Valid: true},
		TmdbPopularity: decimal.NewFromInt(100),
	}

	_, err := testQueries.CreateMovie(context.Background(), arg)
	require.Error(t, err) // Expect an error due to UNIQUE constraint violation
}

func TestUpdateMovieUniqueConstraint(t *testing.T) {
	// Create two movies
	movie1 := createRandomMovie(t)
	movie2 := createRandomMovie(t)

	// Attempt to update movie2 to have the same tmdb_id and imdb_id as movie1
	arg := UpdateMovieParams{
		MovieID:        movie2.MovieID,
		TmdbID:         movie1.TmdbID,                                          // Duplicate TmdbID
		ImdbID:         pgtype.Text{String: movie1.ImdbID.String, Valid: true}, // Duplicate ImdbID
		Title:          "Updated Test Movie",
		Overview:       pgtype.Text{String: "An updated overview.", Valid: true},
		ReleaseDate:    pgtype.Date{Time: time.Now(), Valid: true},
		PosterPath:     pgtype.Text{String: "/updated/poster.jpg", Valid: true},
		BackdropPath:   pgtype.Text{String: "/updated/backdrop.jpg", Valid: true},
		TmdbPopularity: decimal.NewFromInt(95),
	}

	_, err := testQueries.UpdateMovie(context.Background(), arg)
	require.Error(t, err) // Expect an error due to UNIQUE constraint violation
}
