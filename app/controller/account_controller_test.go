package controller

import (
	"bytes"
	"encoding/json"
	"gin-gonic-api/app/domain/dao"
	"gin-gonic-api/app/domain/dto"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTest(body string) (*gin.Context, *httptest.ResponseRecorder, *MockAccountService) {
	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	request, _ := http.NewRequest("GET", "/", bytes.NewBufferString(body))
	context.Request = request
	mockAccountService := new(MockAccountService)
	return context, recorder, mockAccountService
}

func Test_CreateAccountWithInvalidPayload_ReturnsInvalidRequestError(t *testing.T) {
	var response dto.ApiResponse[dao.Account]
	context, recorder, mockAccountService := setupTest("")
	accountControllerImpl := AccountControllerInit(mockAccountService)

	accountControllerImpl.CreateAccount(context)

	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "INVALID_REQUEST", response.ResponseKey)
	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
}

func Test_CreateAccountWithValidPayload_CallsAccountService(t *testing.T) {
	var response dto.ApiResponse[dao.Account]
	context, recorder, mockAccountService := setupTest("{\"name\":\"Test Account\"}")
	accountControllerImpl := AccountControllerInit(mockAccountService)
	mockAccountService.On("CreateAccount", dao.Account{Name: "Test Account"}).
		Return(dao.Account{Name: "Test Account", Status: "Test"})

	accountControllerImpl.CreateAccount(context)

	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "SUCCESS", response.ResponseKey)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "Test", response.Data.Status)
	mockAccountService.AssertCalled(t, "CreateAccount", dao.Account{Name: "Test Account"})
}

func Test_AccountSummary_CallsAccountService(t *testing.T) {
	var response dto.ApiResponse[[]dao.Account]
	context, recorder, mockAccountService := setupTest("")
	accountControllerImpl := AccountControllerInit(mockAccountService)
	mockAccountService.On("GetAll").Return([]dao.Account{{Name: "Test"}})

	accountControllerImpl.AccountSummary(context)

	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "SUCCESS", response.ResponseKey)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, response.Data, dao.Account{Name: "Test"})
	mockAccountService.AssertCalled(t, "GetAll")
}

func Test_AccountDetails(t *testing.T) {
	var response dto.ApiResponse[dao.Account]
	context, recorder, mockAccountService := setupTest("")
	context.Params = []gin.Param{ {
		Key: "accountID",
		Value: "1",
		},
	}
	
	accountControllerImpl := AccountControllerInit(mockAccountService)
	mockAccountService.On("Get", 1).
		Return(dao.Account{Name: "Jane Doe", Status: "Test"})

	accountControllerImpl.AccountDetails(context)

	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "SUCCESS", response.ResponseKey)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "Test", response.Data.Status)
	assert.Equal(t, "Jane Doe", response.Data.Name)
	mockAccountService.AssertCalled(t, "Get", 1)
}
