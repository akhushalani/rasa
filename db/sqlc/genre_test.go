package db

import (
	"context"
	"testing"

	"github.com/akhushalani/rasa/util"
	"github.com/stretchr/testify/require"
)

func createRandomGenre(t *testing.T) Genres {
	// Generate a random genre name
	name := "Genre_" + util.RandomString(6)

	// Create the genre
	genre, err := testQueries.CreateGenre(context.Background(), name)
	require.NoError(t, err)
	require.NotEmpty(t, genre)

	// Validate the created genre
	require.Equal(t, name, genre.Name)

	return genre
}

func TestCreateGenre(t *testing.T) {
	genre := createRandomGenre(t)
	require.NotEmpty(t, genre)
}

func TestGetGenre(t *testing.T) {
	// Create a random genre
	genre := createRandomGenre(t)

	// Fetch the genre by ID
	fetchedGenre, err := testQueries.GetGenre(context.Background(), genre.GenreID)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedGenre)

	// Validate the fetched genre
	require.Equal(t, genre.GenreID, fetchedGenre.GenreID)
	require.Equal(t, genre.Name, fetchedGenre.Name)
}

func TestUpdateGenre(t *testing.T) {
	// Create a random genre
	genre := createRandomGenre(t)

	// Update the genre name
	newName := "Updated_" + genre.Name
	arg := UpdateGenreParams{
		GenreID: genre.GenreID,
		Name:    newName,
	}

	// Perform the update
	updatedGenre, err := testQueries.UpdateGenre(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedGenre)

	// Validate the updated genre
	require.Equal(t, genre.GenreID, updatedGenre.GenreID)
	require.Equal(t, newName, updatedGenre.Name)
}

func TestDeleteGenre(t *testing.T) {
	// Create a random genre
	genre := createRandomGenre(t)

	// Delete the genre
	err := testQueries.DeleteGenre(context.Background(), genre.GenreID)
	require.NoError(t, err)

	// Verify the genre is deleted
	deletedGenre, err := testQueries.GetGenre(context.Background(), genre.GenreID)
	require.Error(t, err)
	require.Empty(t, deletedGenre)
}
