package response_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/EkaRahadi/go-codebase/internal/dto"
	"github.com/EkaRahadi/go-codebase/internal/helper/response"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestResponseOKPlain(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	response.ResponseOKPlain(c)

	// Validate the HTTP status code and response body
	assert.Equal(t, http.StatusOK, w.Code)
	var res dto.SuccessResponse
	err := json.Unmarshal(w.Body.Bytes(), &res)
	assert.NoError(t, err)
	assert.Equal(t, "success", res.Message)
	assert.Nil(t, res.Data)
	assert.Equal(t, "200", res.Code)
}

func TestResponseOKData(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	mockData := map[string]interface{}{
		"key": "value",
	}

	response.ResponseOKData(c, mockData)

	assert.Equal(t, http.StatusOK, w.Code)
	var res dto.SuccessResponse
	err := json.Unmarshal(w.Body.Bytes(), &res)
	assert.NoError(t, err)
	assert.Equal(t, "success", res.Message)
	assert.Equal(t, mockData, res.Data)
	assert.Equal(t, "200", res.Code)
}

func TestResponseOK(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	mockData := map[string]interface{}{
		"key": "value",
	}

	response.ResponseOK(c, "operation successful", mockData)

	assert.Equal(t, http.StatusOK, w.Code)
	var res dto.SuccessResponse
	err := json.Unmarshal(w.Body.Bytes(), &res)
	assert.NoError(t, err)
	assert.Equal(t, "operation successful", res.Message)
	assert.Equal(t, mockData, res.Data)
	assert.Equal(t, "200", res.Code)
}

func TestResponseSuccessJSON(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	mockData := map[string]interface{}{
		"field1": "value1",
	}

	response.ResponseSuccessJSON(c, http.StatusCreated, "created successfully", mockData)

	assert.Equal(t, http.StatusCreated, w.Code)
	var res dto.SuccessResponse
	err := json.Unmarshal(w.Body.Bytes(), &res)
	assert.NoError(t, err)
	assert.Equal(t, "created successfully", res.Message)
	assert.Equal(t, mockData, res.Data)
	assert.Equal(t, "201", res.Code)
}

func TestResponseSuccessJSONCustom(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	mockData := map[string]interface{}{
		"field2": "value2",
	}
	customCode := "CUSTOM-200"

	response.ResponseSuccessJSONCustom(c, http.StatusAccepted, "accepted", customCode, mockData)

	assert.Equal(t, http.StatusAccepted, w.Code)
	var res dto.SuccessResponse
	err := json.Unmarshal(w.Body.Bytes(), &res)
	assert.NoError(t, err)
	assert.Equal(t, "accepted", res.Message)
	assert.Equal(t, mockData, res.Data)
	assert.Equal(t, customCode, res.Code)
}
