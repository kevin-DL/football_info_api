package main

import (
	"football_api/general"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestHealth(t *testing.T) {
	s := CreateServer(nil)

	// Create a New Request
	req, _ := http.NewRequest("GET", "/health", nil)

	// Execute Request
	response := general.ExecuteRequest(req, s)

	// Check the response code
	general.CheckResponseCode(t, http.StatusOK, response.Code)

	// We can use testify/require to assert values, as it is more convenient
	require.Equal(t, "\"OK\"", response.Body.String())
}
