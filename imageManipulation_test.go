package storageImage

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"io"
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

func setupGenerateImgsTest(t *testing.T, fileName, ext, mimeTyp string) (map[string]*bufMedia, error) {
	b, err := ioutil.ReadFile(fmt.Sprintf("sample_image/%s.%s", fileName, ext))
	assert.Nil(t, err)
	buf := bytes.NewBuffer(b)
	return generateImgsByScale(buf, "someThing", fileName, ext, mimeTyp)
}

func Test_generateImgsByScale(t *testing.T) {
	t.Run("decodeOnlySupportedImageTypes", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)
		_, err := generateImgsByScale(buf, "", "", "", "asdfsad")
		assert.NotNil(t, err)
	})

	t.Run("willDecodeValidImageType", func(t *testing.T) {
		t.Run("jpeg", func(t *testing.T) {
			fileName := "cassie-boca-296277"
			ext := "jpg"
			mimeTyp := "image/jpeg"
			img, err := setupGenerateImgsTest(t, fileName, ext, mimeTyp)
			assert.Nil(t, err)
			for _, item := range img {
				assert.NotEqual(t, 0, item.Buf.Len())
				assert.NotNil(t, item.Buf)
				assert.Contains(t, item.Path, fileName)
			}
		})

		t.Run("png", func(t *testing.T) {
			fileName := "cassie-boca-296277"
			ext := "png"
			mimeTyp := "image/png"
			img, err := setupGenerateImgsTest(t, fileName, ext, mimeTyp)
			assert.Nil(t, err)
			for _, item := range img {
				assert.NotEqual(t, 0, item.Buf.Len())
				assert.NotNil(t, item.Buf)
				assert.Contains(t, item.Path, fileName)
			}
		})
	})

	t.Run("willNotDecodeImgSmallerThan1200", func(t *testing.T) {
		t.Run("jpeg", func(t *testing.T) {
			fileName := "small-cassie-boca-296277"
			ext := "jpg"
			_, err := setupGenerateImgsTest(t, fileName, ext, mimeJpeg)
			assert.NotNil(t, err)
			assert.Contains(t, err.Error(), "image has 400 of width, the minimum is 1200")
		})
		t.Run("png", func(t *testing.T) {
			fileName := "small-cassie-boca-296277"
			ext := "png"
			_, err := setupGenerateImgsTest(t, fileName, ext, mimePng)
			assert.NotNil(t, err)
			assert.Contains(t, err.Error(), "image has 400 of width, the minimum is 1200")
		})
	})
}

func Test_hasMinimumSize(t *testing.T) {
	t.Run("shouldHandleErrorDecodingImage", func(t *testing.T) {
		err := hasMinimumSize(nil, func(r io.Reader) (image.Config, error) {
			return image.Config{}, errors.New("error decoding image")
		})
		assert.NotNil(t, err)
	})
}
