package storageImage

import (
	"testing"
	"net/http"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"encoding/json"
	"strings"
)

func TestWriteImage(t *testing.T) {
	path := "/"
	r, err := http.NewRequest("PUT", path, nil)
	assert.Nil(t, err)

	rr := httptest.NewRecorder()
	WriteImage(rr, r)

	b, err := json.Marshal(ImageErrRes{Message: "error reading formfile"})
	assert.Nil(t, err)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, string(b), strings.Trim(rr.Body.String(), "\n"))
}