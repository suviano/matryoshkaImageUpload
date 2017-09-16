package storageImage

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_defineFormat(t *testing.T) {
	t.Run("ImageWithCorrectFormatShouldReturnCorrectType", func(t *testing.T) {
		formats := map[string][]string{
			"png":  {"motoko\\ kusanagi.png", "batou.png", "hideo\\ kuze.png"},
			"bmp":  {"thehuman.bmp", "yellowDogoo.bmp", "pink.bmp", "coolVampire.bmp"},
			"jpeg": {"ulfric-stormcloak.jpg", "general tullius.jpeg", "esbern.jpeg", "arngeir.jpg"},
			"tiff": {"guts.tiff", "griffith.tiff", "casca.tiff", "judeau.tiff", "puck.tiff"},
		}
		for index, formatType := range formats {
			for _, fileName := range formatType {
				img, _, err := fixImgExtension(fileName)
				assert.Nil(t, err)
				assert.Contains(t, img, index)
			}
		}
	})

	t.Run("InvalidFormatShouldThrowException", func(t *testing.T) {
		invalidSource := "i.wrong"
		img, _, err := fixImgExtension(invalidSource)
		assert.NotNil(t, err)
		assert.EqualError(t, errors.New("wrong is not a valid format"), err.Error())
		assert.Empty(t, img)
	})

	t.Run("ShouldNotFormatMalformedImageSource", func(t *testing.T) {
		malformerSource := "sdfs"
		format, _, err := fixImgExtension(malformerSource)
		assert.NotNil(t, err)
		assert.EqualError(t, fmt.Errorf(
			"%s is malformed", malformerSource), err.Error())
		assert.Empty(t, format)
	})
}
