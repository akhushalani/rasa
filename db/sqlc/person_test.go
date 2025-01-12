package db

import (
	"context"
	"testing"
	"time"

	"github.com/akhushalani/rasa/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func createRandomPerson(t *testing.T) People {
	uniqueTmdbID := int32(util.RandomInt(0, 1000000))

	arg := CreatePersonParams{
		TmdbID:             uniqueTmdbID,
		Name:               "Person_" + util.RandomString(6),
		KnownForDepartment: pgtype.Text{String: "Acting", Valid: true},
		Biography:          pgtype.Text{String: "This is a test biography.", Valid: true},
		Birthday:           pgtype.Date{Time: time.Now().AddDate(-30, 0, 0).UTC().Truncate(24 * time.Hour), Valid: true}, // 30 years ago
		Deathday:           pgtype.Date{Valid: false},                                                                    // Not deceased
		Gender:             pgtype.Int2{Int16: 1, Valid: true},                                                           // Gender: Male
		ProfilePath:        pgtype.Text{String: "/path/to/profile.jpg", Valid: true},
		TmdbPopularity:     decimal.NewFromInt(90),
	}

	person, err := testQueries.CreatePerson(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, person)

	require.Equal(t, arg.TmdbID, person.TmdbID)
	require.Equal(t, arg.Name, person.Name)
	require.Equal(t, arg.KnownForDepartment.String, person.KnownForDepartment.String)
	require.Equal(t, arg.Biography.String, person.Biography.String)
	require.WithinDuration(t, arg.Birthday.Time, person.Birthday.Time, time.Second)
	require.Equal(t, arg.Deathday.Valid, person.Deathday.Valid)
	require.Equal(t, arg.Gender.Int16, person.Gender.Int16)
	require.Equal(t, arg.ProfilePath.String, person.ProfilePath.String)

	return person
}

func TestCreatePerson(t *testing.T) {
	person := createRandomPerson(t)
	require.NotEmpty(t, person)
}

func TestGetPerson(t *testing.T) {
	person := createRandomPerson(t)

	fetchedPerson, err := testQueries.GetPerson(context.Background(), person.PersonID)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedPerson)

	require.Equal(t, person.PersonID, fetchedPerson.PersonID)
	require.Equal(t, person.TmdbID, fetchedPerson.TmdbID)
	require.Equal(t, person.Name, fetchedPerson.Name)
	require.Equal(t, person.KnownForDepartment.String, fetchedPerson.KnownForDepartment.String)
	require.Equal(t, person.Biography.String, fetchedPerson.Biography.String)
	require.WithinDuration(t, person.Birthday.Time, fetchedPerson.Birthday.Time, time.Second)
	require.Equal(t, person.Deathday.Valid, fetchedPerson.Deathday.Valid)
	require.Equal(t, person.Gender.Int16, fetchedPerson.Gender.Int16)
	require.Equal(t, person.ProfilePath.String, fetchedPerson.ProfilePath.String)
}

func TestUpdatePerson(t *testing.T) {
	person := createRandomPerson(t)

	arg := UpdatePersonParams{
		PersonID:           person.PersonID,
		TmdbID:             person.TmdbID,
		Name:               "Updated_" + person.Name,
		KnownForDepartment: pgtype.Text{String: "Directing", Valid: true},
		Biography:          pgtype.Text{String: "Updated biography.", Valid: true},
		Birthday:           pgtype.Date{Time: time.Now().AddDate(-40, 0, 0).UTC().Truncate(24 * time.Hour), Valid: true}, // 40 years ago
		Deathday:           pgtype.Date{Time: time.Now().AddDate(-10, 0, 0).UTC().Truncate(24 * time.Hour), Valid: true}, // 10 years ago
		Gender:             pgtype.Int2{Int16: 2, Valid: true},                                                           // Gender: Female
		ProfilePath:        pgtype.Text{String: "/updated/path/to/profile.jpg", Valid: true},
		TmdbPopularity:     decimal.NewFromInt(85), // Popularity: 85
	}

	updatedPerson, err := testQueries.UpdatePerson(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedPerson)

	require.Equal(t, arg.PersonID, updatedPerson.PersonID)
	require.Equal(t, arg.Name, updatedPerson.Name)
	require.Equal(t, arg.KnownForDepartment.String, updatedPerson.KnownForDepartment.String)
	require.Equal(t, arg.Biography.String, updatedPerson.Biography.String)
	require.WithinDuration(t, arg.Birthday.Time, updatedPerson.Birthday.Time, time.Second)
	require.WithinDuration(t, arg.Deathday.Time, updatedPerson.Deathday.Time, time.Second)
	require.Equal(t, arg.Gender.Int16, updatedPerson.Gender.Int16)
	require.Equal(t, arg.ProfilePath.String, updatedPerson.ProfilePath.String)
	require.Equal(t, arg.TmdbPopularity.Cmp(updatedPerson.TmdbPopularity), 0)
}

func TestDeletePerson(t *testing.T) {
	person := createRandomPerson(t)

	err := testQueries.DeletePerson(context.Background(), person.PersonID)
	require.NoError(t, err)

	// Verify the person is deleted
	fetchedPerson, err := testQueries.GetPerson(context.Background(), person.PersonID)
	require.Error(t, err)
	require.Empty(t, fetchedPerson)
}
