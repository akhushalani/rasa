package db

import (
	"context"
	"testing"
	"time"

	"github.com/akhushalani/rasa/util"
	"github.com/stretchr/testify/require"
)

func logRandomMovieCache(t *testing.T) MovieCacheLog {
	// Generate a random tmdbID
	tmdbID := int32(util.RandomInt(0, 1000000))

	// Insert into movie_cache_log
	movieCache, err := testQueries.LogMovieCache(context.Background(), tmdbID)
	require.NoError(t, err)
	require.NotEmpty(t, movieCache)

	// Validate the fields
	require.Equal(t, tmdbID, movieCache.TmdbID)
	require.WithinDuration(t, time.Now(), movieCache.LastFetched.Time, time.Second)

	return movieCache
}

func TestLogMovieCache(t *testing.T) {
	movieCache := logRandomMovieCache(t)

	require.Equal(t, movieCache.TmdbID, movieCache.TmdbID) // Double-check the ID matches
	require.WithinDuration(t, time.Now(), movieCache.LastFetched.Time, time.Second)
}

func TestGetMovieCache(t *testing.T) {
	movieCache := logRandomMovieCache(t)

	fetchedCache, err := testQueries.GetMovieCache(context.Background(), movieCache.TmdbID)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedCache)

	require.Equal(t, movieCache.TmdbID, fetchedCache.TmdbID)
	require.WithinDuration(t, movieCache.LastFetched.Time, fetchedCache.LastFetched.Time, time.Second)
}

func TestUpdateMovieCache(t *testing.T) {
	movieCache := logRandomMovieCache(t)

	updatedCache, err := testQueries.UpdateMovieCache(context.Background(), movieCache.TmdbID)
	require.NoError(t, err)
	require.NotEmpty(t, updatedCache)

	require.Equal(t, movieCache.TmdbID, updatedCache.TmdbID)
	require.WithinDuration(t, time.Now(), updatedCache.LastFetched.Time, time.Second)
}

func TestDeleteMovieCache(t *testing.T) {
	movieCache := logRandomMovieCache(t)

	err := testQueries.DeleteMovieCache(context.Background(), movieCache.TmdbID)
	require.NoError(t, err)

	// Verify the movie cache log is deleted
	fetchedCache, err := testQueries.GetMovieCache(context.Background(), movieCache.TmdbID)
	require.Error(t, err)
	require.Empty(t, fetchedCache)
}
