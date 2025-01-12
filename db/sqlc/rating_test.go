package db

import (
	"context"
	"testing"

	decimal "github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func createRandomRating(t *testing.T) Ratings {
	user := createRandomUser(t)
	movie := createRandomMovie(t)

	arg := CreateRatingParams{
		MovieID:     movie.MovieID,
		UserID:      user.UserID,
		RatingScore: decimal.NewFromInt(85), // Rating: 8.5
	}

	rating, err := testQueries.CreateRating(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, rating)

	require.Equal(t, arg.MovieID, rating.MovieID)
	require.Equal(t, arg.UserID, rating.UserID)
	require.Equal(t, arg.RatingScore.Cmp(rating.RatingScore), 0)

	return rating
}

func TestCreateRating(t *testing.T) {
	rating := createRandomRating(t)
	require.NotEmpty(t, rating)
}

func TestGetRating(t *testing.T) {
	rating := createRandomRating(t)

	arg := GetRatingParams{
		MovieID: rating.MovieID,
		UserID:  rating.UserID,
	}

	fetchedRating, err := testQueries.GetRating(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedRating)

	require.Equal(t, rating.MovieID, fetchedRating.MovieID)
	require.Equal(t, rating.UserID, fetchedRating.UserID)
	require.Equal(t, rating.RatingScore.Cmp(fetchedRating.RatingScore), 0)
}

func TestUpdateRating(t *testing.T) {
	rating := createRandomRating(t)

	arg := UpdateRatingParams{
		MovieID:     rating.MovieID,
		UserID:      rating.UserID,
		RatingScore: decimal.NewFromInt(95), // Updated rating: 9.5
	}

	updatedRating, err := testQueries.UpdateRating(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedRating)

	require.Equal(t, arg.MovieID, updatedRating.MovieID)
	require.Equal(t, arg.UserID, updatedRating.UserID)
	require.Equal(t, arg.RatingScore.Cmp(updatedRating.RatingScore), 0)
}

func TestDeleteRating(t *testing.T) {
	rating := createRandomRating(t)

	arg := DeleteRatingParams{
		MovieID: rating.MovieID,
		UserID:  rating.UserID,
	}

	err := testQueries.DeleteRating(context.Background(), arg)
	require.NoError(t, err)

	// Verify the rating is deleted
	argGet := GetRatingParams{
		MovieID: rating.MovieID,
		UserID:  rating.UserID,
	}
	fetchedRating, err := testQueries.GetRating(context.Background(), argGet)
	require.Error(t, err)
	require.Empty(t, fetchedRating)
}

func TestCascadeDeleteRatingOnUser(t *testing.T) {
	rating := createRandomRating(t)

	// Delete the user
	err := testQueries.DeleteUser(context.Background(), rating.UserID)
	require.NoError(t, err)

	// Verify the rating is deleted
	arg := GetRatingParams{
		MovieID: rating.MovieID,
		UserID:  rating.UserID,
	}
	fetchedRating, err := testQueries.GetRating(context.Background(), arg)
	require.Error(t, err)
	require.Empty(t, fetchedRating)
}

func TestCascadeDeleteRatingOnMovie(t *testing.T) {
	rating := createRandomRating(t)

	// Delete the movie
	err := testQueries.DeleteMovie(context.Background(), rating.MovieID)
	require.NoError(t, err)

	// Verify the rating is deleted
	arg := GetRatingParams{
		MovieID: rating.MovieID,
		UserID:  rating.UserID,
	}
	fetchedRating, err := testQueries.GetRating(context.Background(), arg)
	require.Error(t, err)
	require.Empty(t, fetchedRating)
}
