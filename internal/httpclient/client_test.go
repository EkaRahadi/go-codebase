package httpclient_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/EkaRahadi/go-codebase/internal/httpclient"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method) // check if request method is equal
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"message":"success"}`))
	}))
	defer mockServer.Close()

	client := httpclient.NewClient()

	response, err := client.Get(context.Background(), mockServer.URL)

	assert.NoError(t, err)
	assert.Equal(t, `{"message":"success"}`, response)
}

func TestPost(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)

		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)
		assert.Equal(t, `{"key":"value"}`, string(body))

		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"status":"created"}`))
	}))
	defer mockServer.Close()

	client := httpclient.NewClient()

	response, err := client.Post(context.Background(), mockServer.URL, []byte(`{"key":"value"}`))

	assert.NoError(t, err)
	assert.Equal(t, `{"status":"created"}`, response)
}

func TestGetWithQuery(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)

		// Assert query parameters
		query := r.URL.Query()
		assert.Equal(t, "value1", query.Get("key1"))
		assert.Equal(t, "value2", query.Get("key2"))

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"result":"query_success"}`))
	}))
	defer mockServer.Close()

	client := httpclient.NewClient()

	response, err := client.GetWithQuery(context.Background(), mockServer.URL, map[string]string{
		"key1": "value1",
		"key2": "value2",
	})

	assert.NoError(t, err)
	assert.Equal(t, `{"result":"query_success"}`, response)
}

func TestErrorResponse(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
	}))
	defer mockServer.Close()

	client := httpclient.NewClient()

	response, err := client.Get(context.Background(), mockServer.URL)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "received error response")
	assert.Contains(t, response, "something went wrong")
}
