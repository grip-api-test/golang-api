package integration_test

import (
	"bytes"
	"encoding/json"
	"gin-gonic-api/app/domain/dao"
	"gin-gonic-api/app/domain/dto"
	"gin-gonic-api/app/router"
	"gin-gonic-api/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetAccounts_ReturnsEmptyList(t *testing.T) {
	var response dto.ApiResponse[[]dao.Account]
	init := config.Init()
	app := router.Init(init)

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/accounts", nil)
	app.ServeHTTP(recorder, request)
	err := json.Unmarshal(recorder.Body.Bytes(), &response)

	assert.Nil(t, err)
	assert.Equal(t, "SUCCESS", response.ResponseKey)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Empty(t, response.Data)
}

func Test_CreateAndGetAccounts_ReturnsCreatedAcccount(t *testing.T) {
	var response dto.ApiResponse[[]dao.Account]
	init := config.Init()
	app := router.Init(init)

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/accounts", bytes.NewBufferString("{\"name\":\"Test Account\"}"))
	app.ServeHTTP(recorder, request)

	recorder = httptest.NewRecorder()
	request, _ = http.NewRequest("GET", "/api/accounts", nil)
	app.ServeHTTP(recorder, request)
	err := json.Unmarshal(recorder.Body.Bytes(), &response)

	assert.Nil(t, err)
	assert.Equal(t, "SUCCESS", response.ResponseKey)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, response.Data, dao.Account{Name: "Test Account"})
}


func Test_GetAccountById_ReturnsCreatedAcccount(t *testing.T) {
	var response dto.ApiResponse[dao.Account]
	init := config.Init()
	app := router.Init(init)

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/accounts", bytes.NewBufferString("{\"name\":\"Test Account\"}"))
	app.ServeHTTP(recorder, request)

	recorder = httptest.NewRecorder()
	request, _ = http.NewRequest("GET", "/api/accounts/1", nil)
	app.ServeHTTP(recorder, request)
	err := json.Unmarshal(recorder.Body.Bytes(), &response)

	assert.Nil(t, err)
	assert.Equal(t, "SUCCESS", response.ResponseKey)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, response.Data.Name, "Test Account")
}
