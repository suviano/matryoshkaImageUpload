package matryoshka

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteImage(t *testing.T) {
	sampleBuf, err := ioutil.ReadFile("sample_image/small-cassie-boca-296277.jpg")
	assert.Nil(t, err)
	t.Run("EmptyPrefix", func(t *testing.T) {
		err := WriteImage("", "", "", nil)
		assert.NotNil(t, err)
		assert.Error(t, err, "empty prefix")
	})
	t.Run("EmptyBucket", func(t *testing.T) {
		err := WriteImage("aprefix", "prettyImg.jpeg", "", nil)
		assert.NotNil(t, err)
		assert.Error(t, err, "empty bucket")
	})
	t.Run("EmptyBuffer", func(t *testing.T) {
		err := WriteImage("aprefix", "prettyImg.jpeg", "abucket", nil)
		assert.NotNil(t, err)
		assert.Error(t, err, "empty buffer")
	})
	t.Run("malformedImagePath", func(t *testing.T) {
		malformedFilePath := "notextension"
		err := WriteImage("aprefix", malformedFilePath, "abucket", bytes.NewBuffer(sampleBuf))
		assert.NotNil(t, err)
		assert.EqualError(t, err, fmt.Sprintf("%s is malformed", malformedFilePath))
	})
	t.Run("errorGeneratingScaledImages", func(t *testing.T) {
		filePath := "sample_image/small-cassie-boca-296277.jpg"
		err := WriteImage("aprefix", filePath, "abucket", bytes.NewBuffer(sampleBuf))
		assert.NotNil(t, err)
		assert.EqualError(t, err, "image has 400 of width, the minimum is 1200")
	})
}
