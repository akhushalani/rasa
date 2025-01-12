package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	decimal "github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func createRandomUserMovie(t *testing.T) UserMovies {
	user := createRandomUser(t)
	movie := createRandomMovie(t)

	arg := CreateUserMovieParams{
		UserID:    user.UserID,
		MovieID:   movie.MovieID,
		Rating:    decimal.NewFromFloat(4.5), // Rating: 4.5
		Review:    pgtype.Text{String: "Great movie!", Valid: true},
		Watchlist: pgtype.Bool{Bool: true, Valid: true},
		Watched:   pgtype.Bool{Bool: true, Valid: true},
		Favorited: pgtype.Bool{Bool: false, Valid: true},
	}

	userMovie, err := testQueries.CreateUserMovie(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, userMovie)

	require.Equal(t, arg.UserID, userMovie.UserID)
	require.Equal(t, arg.MovieID, userMovie.MovieID)
	require.Equal(t, arg.Rating.String(), userMovie.Rating.String())
	require.Equal(t, arg.Review.String, userMovie.Review.String)

	return userMovie
}

func TestCreateUserMovie(t *testing.T) {
	userMovie := createRandomUserMovie(t)
	require.NotEmpty(t, userMovie)
}

func TestGetUserMovie(t *testing.T) {
	userMovie := createRandomUserMovie(t)

	arg := GetUserMovieParams{
		UserID:  userMovie.UserID,
		MovieID: userMovie.MovieID,
	}

	fetchedUserMovie, err := testQueries.GetUserMovie(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedUserMovie)

	require.Equal(t, userMovie.UserID, fetchedUserMovie.UserID)
	require.Equal(t, userMovie.MovieID, fetchedUserMovie.MovieID)
	require.Equal(t, userMovie.Rating.String(), fetchedUserMovie.Rating.String())
	require.Equal(t, userMovie.Review.String, fetchedUserMovie.Review.String)
}

func TestUpdateUserMovie(t *testing.T) {
	userMovie := createRandomUserMovie(t)

	arg := UpdateUserMovieParams{
		UserID:    userMovie.UserID,
		MovieID:   userMovie.MovieID,
		Rating:    decimal.NewFromFloat(3.8), // Updated Rating: 3.8
		Review:    pgtype.Text{String: "Good movie.", Valid: true},
		Watchlist: pgtype.Bool{Bool: false, Valid: true},
		Watched:   pgtype.Bool{Bool: true, Valid: true},
		Favorited: pgtype.Bool{Bool: true, Valid: true},
	}

	updatedUserMovie, err := testQueries.UpdateUserMovie(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedUserMovie)

	require.Equal(t, arg.UserID, updatedUserMovie.UserID)
	require.Equal(t, arg.MovieID, updatedUserMovie.MovieID)
	require.Equal(t, arg.Rating.String(), updatedUserMovie.Rating.String())
	require.Equal(t, arg.Review.String, updatedUserMovie.Review.String)
	require.Equal(t, arg.Favorited.Bool, updatedUserMovie.Favorited.Bool)
}

func TestListUserMovies(t *testing.T) {
	userMovie := createRandomUserMovie(t)

	arg := ListUserMoviesParams{
		UserID: userMovie.UserID,
		Limit:  10,
		Offset: 0,
	}

	userMovies, err := testQueries.ListUserMovies(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, userMovies)

	require.Equal(t, userMovie.UserID, userMovies[0].UserID)
	require.Equal(t, userMovie.MovieID, userMovies[0].MovieID)
}
