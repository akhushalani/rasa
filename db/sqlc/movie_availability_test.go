package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomMovieAvailability(t *testing.T) MovieAvailability {
	movie := createRandomMovie(t)
	service := createRandomStreamingService(t)

	arg := CreateMovieAvailabilityParams{
		MovieID:   movie.MovieID,
		ServiceID: service.ServiceID,
	}

	availability, err := testQueries.CreateMovieAvailability(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, availability)

	require.Equal(t, arg.MovieID, availability.MovieID)
	require.Equal(t, arg.ServiceID, availability.ServiceID)

	return availability
}

func TestCreateMovieAvailability(t *testing.T) {
	availability := createRandomMovieAvailability(t)
	require.NotEmpty(t, availability)
}

func TestGetMovieAvailabilities(t *testing.T) {
	availability := createRandomMovieAvailability(t)

	arg := GetMovieAvailabilitiesParams{
		MovieID: availability.MovieID,
		Limit:   10,
		Offset:  0,
	}

	availabilities, err := testQueries.GetMovieAvailabilities(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, availabilities)

	require.Equal(t, availability.MovieID, availabilities[0].MovieID)
	require.Equal(t, availability.ServiceID, availabilities[0].ServiceID)
}

func TestDeleteMovieAvailability(t *testing.T) {
	availability := createRandomMovieAvailability(t)

	arg := DeleteMovieAvailabilityParams{
		MovieID:   availability.MovieID,
		ServiceID: availability.ServiceID,
	}

	err := testQueries.DeleteMovieAvailability(context.Background(), arg)
	require.NoError(t, err)

	// Verify the association is deleted
	argGet := GetMovieAvailabilitiesParams{
		MovieID: availability.MovieID,
		Limit:   10,
		Offset:  0,
	}
	availabilities, err := testQueries.GetMovieAvailabilities(context.Background(), argGet)
	require.NoError(t, err)
	require.Empty(t, availabilities)
}

func TestCascadeDeleteMovieAvailabilityOnMovie(t *testing.T) {
	availability := createRandomMovieAvailability(t)

	// Delete the movie
	err := testQueries.DeleteMovie(context.Background(), availability.MovieID)
	require.NoError(t, err)

	// Verify the association is deleted
	arg := GetMovieAvailabilitiesParams{
		MovieID: availability.MovieID,
		Limit:   10,
		Offset:  0,
	}
	availabilities, err := testQueries.GetMovieAvailabilities(context.Background(), arg)
	require.NoError(t, err)
	require.Empty(t, availabilities)
}

func TestCascadeDeleteMovieAvailabilityOnStreamingService(t *testing.T) {
	availability := createRandomMovieAvailability(t)

	// Delete the streaming service
	err := testQueries.DeleteStreamingService(context.Background(), availability.ServiceID)
	require.NoError(t, err)

	// Verify the association is deleted
	arg := GetMovieAvailabilitiesParams{
		MovieID: availability.MovieID,
		Limit:   10,
		Offset:  0,
	}
	availabilities, err := testQueries.GetMovieAvailabilities(context.Background(), arg)
	require.NoError(t, err)
	require.Empty(t, availabilities)
}
