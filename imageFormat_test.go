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
			"jpeg": {"rocket.jpg", "groot.jpeg", "makesomuchmoney.jpeg"},
			"png":  {"motoko\\ kusanagi.png", "batou.png", "hideo\\ kuze.png"},
			"bmp":  {"thehuman.bmp", "yellowDogoo.bmp", "pink.bmp", "coolVampire.bmp"},
			"gif":  {"ulfric-stormcloak.gif", "general tullius.gif", "esbern.gif", "arngeir.gif"},
			"tiff": {"guts.tiff", "griffith.tiff", "casca.tiff", "judeau.tiff", "puck.tiff"},
		}
		for index, formatType := range formats {
			for _, fileName := range formatType {
				format, err := defineFormat(fileName)
				assert.Nil(t, err)
				assert.Equal(t, index, format)
			}
		}
	})

	t.Run("InvalidFormatShouldThrowException", func(t *testing.T) {
		invalidSource := "i.wrong"
		format, err := defineFormat(invalidSource)
		assert.NotNil(t, err)
		assert.EqualError(t, errors.New("wrong is not a valid format"), err.Error())
		assert.Empty(t, format)
	})

	t.Run("ShouldNotFormatMalformedImageSource", func(t *testing.T) {
		malformerSource := "sdfs"
		format, err := defineFormat(malformerSource)
		assert.NotNil(t, err)
		assert.EqualError(t, fmt.Errorf(
			"%s is malformed", malformerSource), err.Error())
		assert.Empty(t, format)
	})
}
