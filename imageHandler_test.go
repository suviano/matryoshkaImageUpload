package storageImage

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteImage(t *testing.T) {
	sampleBuf, err := ioutil.ReadFile("sample_image/small-cassie-boca-296277.jpg")
	assert.Nil(t, err)
	t.Run("AnyPrefixProvided", func(t *testing.T) {
		err := WriteImage("", "", nil)
		assert.NotNil(t, err)
		assert.Error(t, err, "empty prefix")
	})
	t.Run("EmptyBuffer", func(t *testing.T) {
		err := WriteImage("aprefix", "prettyImg.jpeg", nil)
		assert.NotNil(t, err)
		assert.Error(t, err, "empty buffer")
	})
	t.Run("malformedImagePath", func(t *testing.T) {
		malformedFilePath := "notextension"
		err := WriteImage("aprefix", malformedFilePath, bytes.NewBuffer(sampleBuf))
		assert.NotNil(t, err)
		assert.Errorf(t, err, "%s is malformed", malformedFilePath)
	})
	t.Run("errorGeneratingScaledImages", func(t *testing.T) {
		filePath := "sample_image/small-cassie-boca-296277.jpg"
		err := WriteImage("aprefix", filePath, bytes.NewBuffer(sampleBuf))
		assert.NotNil(t, err)
		assert.Error(t, err, "image has 400 of width, the minimum is 1200")
	})
}
