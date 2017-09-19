package storageImage

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteImage(t *testing.T) {
	t.Run("BadRequestForNoFile", func(t *testing.T) {
		r, err := http.NewRequest("PUT", "/image/landspaces", nil)
		assert.Nil(t, err)

		rr := httptest.NewRecorder()
		WriteImage(rr, r)

		b, err := json.Marshal(imageRes{Message: "error reading formfile"})
		assert.Nil(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, string(b), strings.Trim(rr.Body.String(), "\n"))
	})

	/*TODO fix this test
	t.Run("SuccessFullySaveImage", func(t *testing.T) {
		fileName := "sample_image/cassie-boca-296277.jpg"
		file, err := os.Open(fileName)
		assert.Nil(t, err)
		defer file.Close()

		body := bytes.NewBuffer(nil)
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("object", fileName)
		assert.Nil(t, err)

		_, err = io.Copy(part, file)
		assert.Nil(t, err)

		err = writer.Close()
		assert.Nil(t, err)

		r, err := http.NewRequest("PUT", "/image/landspaces", body)
		r.Header.Set("Content-Type", writer.FormDataContentType())
		assert.Nil(t, err)
		rr := httptest.NewRecorder()
		WriteImage(rr, r)

		bytJSON, err := json.Marshal(ImageRes{Message: "error reading formfile"})
		assert.Nil(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, string(bytJSON), strings.Trim(rr.Body.String(), "\n"))
	})*/
}
