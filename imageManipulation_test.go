package storageImage

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_defineFormat(t *testing.T) {
	t.Run("ImageWithCorrectFormatShouldReturnCorrectType", func(t *testing.T) {
		formats := map[string][]string{
			"jpeg": {"ulfric-stormcloak.jpg", "general tullius.jpeg", "esbern.jpeg", "arngeir.jpg"},
			"png":  {"guts.png", "griffith.png", "casca.png", "judeau.png", "puck.png"},
		}
		for _, formatType := range formats {
			for _, fileName := range formatType {
				_, _, _, err := fixImgExtension(fileName)
				assert.Nil(t, err)
			}
		}
	})

	t.Run("InvalidFormatShouldThrowException", func(t *testing.T) {
		invalidSource := "i.wrong"
		img, _, _, err := fixImgExtension(invalidSource)
		assert.NotNil(t, err)
		assert.EqualError(t, errors.New("wrong is not a valid format"), err.Error())
		assert.Empty(t, img)
	})

	t.Run("ShouldNotFormatMalformedImageSource", func(t *testing.T) {
		malformerSource := "sdfs"
		format, _, _, err := fixImgExtension(malformerSource)
		assert.NotNil(t, err)
		assert.EqualError(t, fmt.Errorf(
			"%s is malformed", malformerSource), err.Error())
		assert.Empty(t, format)
	})
}

func Test_generateImgsByScale(t *testing.T) {
	t.Run("decodeOnlySupportedImageTypes", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)
		_, err := generateImgsByScale(buf, "", "", "", "asdfsad")
		assert.NotNil(t, err)
	})

	t.Run("willDecodeValidImageType", func(t *testing.T) {
		prefix := "someThing"
		setupGenerateImgsTest := func(t *testing.T, fileName, ext, mimeTyp string) {
			b, err := ioutil.ReadFile(fmt.Sprintf("sample_image/%s.%s", fileName, ext))
			assert.Nil(t, err)
			buf := bytes.NewBuffer(b)
			img, err := generateImgsByScale(buf, prefix, fileName, ext, mimeTyp)
			assert.Nil(t, err)
			for _, item := range img {
				assert.NotEqual(t, 0, item.Buf.Len())
				assert.NotNil(t, item.Buf)
				assert.Contains(t, item.Path, fileName)
			}
		}

		t.Run("jpeg", func(t *testing.T) {
			fileName := "cassie-boca-296277"
			ext := "jpg"
			mimeTyp := "image/jpeg"
			setupGenerateImgsTest(t, fileName, ext, mimeTyp)
		})

		t.Run("png", func(t *testing.T) {
			fileName := "cassie-boca-296277"
			ext := "png"
			mimeTyp := "image/png"
			setupGenerateImgsTest(t, fileName, ext, mimeTyp)
		})
	})
}
