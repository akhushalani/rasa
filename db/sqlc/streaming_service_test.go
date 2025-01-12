package db

import (
	"context"
	"testing"

	"github.com/akhushalani/rasa/util"
	"github.com/stretchr/testify/require"
)

func createRandomStreamingService(t *testing.T) StreamingServices {
	name := "Service_" + util.RandomString(6)

	service, err := testQueries.CreateStreamingService(context.Background(), name)
	require.NoError(t, err)
	require.NotEmpty(t, service)

	require.Equal(t, name, service.Name)
	require.Empty(t, service.LogoPath) // Logo path should be empty by default

	return service
}

func TestCreateStreamingService(t *testing.T) {
	service := createRandomStreamingService(t)
	require.NotEmpty(t, service)
}

func TestGetStreamingService(t *testing.T) {
	service := createRandomStreamingService(t)

	fetchedService, err := testQueries.GetStreamingService(context.Background(), service.ServiceID)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedService)

	require.Equal(t, service.ServiceID, fetchedService.ServiceID)
	require.Equal(t, service.Name, fetchedService.Name)
	require.Equal(t, service.LogoPath, fetchedService.LogoPath)
}

func TestUpdateStreamingService(t *testing.T) {
	service := createRandomStreamingService(t)

	arg := UpdateStreamingServiceParams{
		ServiceID: service.ServiceID,
		Name:      "Updated_" + service.Name,
	}

	updatedService, err := testQueries.UpdateStreamingService(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedService)

	require.Equal(t, arg.ServiceID, updatedService.ServiceID)
	require.Equal(t, arg.Name, updatedService.Name)
	require.Equal(t, service.LogoPath, updatedService.LogoPath) // Logo path should remain unchanged
}

func TestDeleteStreamingService(t *testing.T) {
	service := createRandomStreamingService(t)

	err := testQueries.DeleteStreamingService(context.Background(), service.ServiceID)
	require.NoError(t, err)

	// Verify the service is deleted
	fetchedService, err := testQueries.GetStreamingService(context.Background(), service.ServiceID)
	require.Error(t, err)
	require.Empty(t, fetchedService)
}

func TestListStreamingServices(t *testing.T) {
	// Create multiple streaming services
	for i := 0; i < 5; i++ {
		createRandomStreamingService(t)
	}

	arg := ListStreamingServicesParams{
		Limit:  5,
		Offset: 0,
	}

	services, err := testQueries.ListStreamingServices(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, services, int(arg.Limit))

	// Validate each service
	for _, service := range services {
		require.NotEmpty(t, service)
	}
}