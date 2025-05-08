//go:build applicationtest

package applicationtest

import (
	"fmt"
	"github.com/go-tstr/golden"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestTasks(t *testing.T) {

	tests := []struct {
		name               string
		method             string
		path               string
		body               io.Reader
		expectedStatusCode int
	}{
		{
			name:               "Create Task",
			method:             http.MethodPost,
			path:               "/api/v1/tasks",
			body:               strings.NewReader(`{"name":"new test"}`),
			expectedStatusCode: 200,
		},
		{
			name:               "Create second Task",
			method:             http.MethodPost,
			path:               "/api/v1/tasks",
			body:               strings.NewReader(`{"name":"new test 2"}`),
			expectedStatusCode: 200,
		},
		{
			name:               "Get All Task",
			method:             http.MethodGet,
			path:               "/api/v1/tasks",
			body:               nil,
			expectedStatusCode: 200,
		},
		{
			name:               "Delete Task 1",
			method:             http.MethodDelete,
			path:               "/api/v1/tasks/1",
			body:               nil,
			expectedStatusCode: 200,
		},
		{
			name:               "Get All Task",
			method:             http.MethodGet,
			path:               "/api/v1/tasks",
			body:               nil,
			expectedStatusCode: 200,
		},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("%d %s", i, tc.name), func(t *testing.T) {
			req, err := http.NewRequest(tc.method, appURL+tc.path, tc.body)
			require.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")
			golden.Request(t, http.DefaultClient, req, tc.expectedStatusCode)
		})

	}

	//t.Run("tasks", func(t *testing.T) {
	//	req, err := http.NewRequest(http.MethodPost, appURL+"/api/v1/tasks", strings.NewReader(`{"name":"new test"}`))
	//	require.NoError(t, err)
	//	req.Header.Set("Content-Type", "application/json")
	//	golden.Request(t, http.DefaultClient, req, 200)
	//})
	//
	//t.Run("getAll", func(t *testing.T) {
	//	req, err := http.NewRequest(http.MethodGet, appURL+"/api/v1/tasks", nil)
	//	require.NoError(t, err)
	//	req.Header.Set("Content-Type", "application/json")
	//	golden.Request(t, http.DefaultClient, req, 200)
	//
	//})
}
