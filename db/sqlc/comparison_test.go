package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomComparison(t *testing.T) Comparisons {
	user := createRandomUser(t)
	baseMovie := createRandomMovie(t)
	comparedMovie := createRandomMovie(t)

	arg := CreateComparisonParams{
		UserID:          user.UserID,
		BaseMovieID:     baseMovie.MovieID,
		ComparedMovieID: comparedMovie.MovieID,
		Preference:      1, // Prefers base movie
	}

	comparison, err := testQueries.CreateComparison(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, comparison)

	require.Equal(t, arg.UserID, comparison.UserID)
	require.Equal(t, arg.BaseMovieID, comparison.BaseMovieID)
	require.Equal(t, arg.ComparedMovieID, comparison.ComparedMovieID)
	require.Equal(t, arg.Preference, comparison.Preference)

	return comparison
}

func TestCreateComparison(t *testing.T) {
	comparison := createRandomComparison(t)
	require.NotEmpty(t, comparison)
}

func TestGetComparison(t *testing.T) {
	comparison := createRandomComparison(t)

	fetchedComparison, err := testQueries.GetComparison(context.Background(), comparison.ComparisonID)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedComparison)

	require.Equal(t, comparison.ComparisonID, fetchedComparison.ComparisonID)
	require.Equal(t, comparison.UserID, fetchedComparison.UserID)
	require.Equal(t, comparison.BaseMovieID, fetchedComparison.BaseMovieID)
	require.Equal(t, comparison.ComparedMovieID, fetchedComparison.ComparedMovieID)
	require.Equal(t, comparison.Preference, fetchedComparison.Preference)
}

func TestUpdateComparison(t *testing.T) {
	comparison := createRandomComparison(t)

	arg := UpdateComparisonParams{
		ComparisonID: comparison.ComparisonID,
		Preference:   0, // Prefers compared movie
	}

	updatedComparison, err := testQueries.UpdateComparison(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedComparison)

	require.Equal(t, arg.ComparisonID, updatedComparison.ComparisonID)
	require.Equal(t, arg.Preference, updatedComparison.Preference)
}

func TestDeleteComparison(t *testing.T) {
	comparison := createRandomComparison(t)

	err := testQueries.DeleteComparison(context.Background(), comparison.ComparisonID)
	require.NoError(t, err)

	// Verify the comparison is deleted
	fetchedComparison, err := testQueries.GetComparison(context.Background(), comparison.ComparisonID)
	require.Error(t, err)
	require.Empty(t, fetchedComparison)
}

func TestCascadeDeleteComparisonOnUser(t *testing.T) {
	comparison := createRandomComparison(t)

	// Delete the user
	err := testQueries.DeleteUser(context.Background(), comparison.UserID)
	require.NoError(t, err)

	// Verify the comparison is deleted
	fetchedComparison, err := testQueries.GetComparison(context.Background(), comparison.ComparisonID)
	require.Error(t, err)
	require.Empty(t, fetchedComparison)
}

func TestCascadeDeleteComparisonOnMovie(t *testing.T) {
	comparison := createRandomComparison(t)

	// Delete the base movie
	err := testQueries.DeleteMovie(context.Background(), comparison.BaseMovieID)
	require.NoError(t, err)

	// Verify the comparison is deleted
	fetchedComparison, err := testQueries.GetComparison(context.Background(), comparison.ComparisonID)
	require.Error(t, err)
	require.Empty(t, fetchedComparison)
}