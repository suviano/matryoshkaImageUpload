package storageImage

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupImageTest(t *testing.T, body io.Reader) (*httptest.ResponseRecorder, *http.Request) {
	t.Helper()
	path := "/image"
	r, err := http.NewRequest("PUT", path, body)
	assert.Nil(t, err)

	rr := httptest.NewRecorder()
	return rr, r
}

func TestWriteImage(t *testing.T) {
	t.Run("BadRequestForNoFile", func(t *testing.T) {
		rr, r := setupImageTest(t, nil)
		WriteImage(rr, r)

		b, err := json.Marshal(ImageErrRes{Message: "error reading formfile"})
		assert.Nil(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, string(b), strings.Trim(rr.Body.String(), "\n"))
	})
}
