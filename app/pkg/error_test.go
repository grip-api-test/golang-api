package pkg

import (
	"gin-gonic-api/app/constant"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PanicExceptionOnInvalidRequest_ThrowsWithErrorMessage(t *testing.T) {
	assert.PanicsWithError(t, "INVALID_REQUEST: Invalid Request", func() { PanicException(constant.InvalidRequest) }, "The code did not panic")
}
