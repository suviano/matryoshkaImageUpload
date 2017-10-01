package matryoshka

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_defineFormat(t *testing.T) {
	t.Run("ImageWithCorrectFormatShouldReturnCorrectType", func(t *testing.T) {
		formats := map[string][]string{
			"jpeg": {"ulfric-stormcloak.jpg", "general tullius.jpeg", "esbern.jpeg", "arngeir.jpg"},
			"png":  {"guts.png", "griffith.png", "casca.png", "judeau.png", "puck.png"},
		}
		for ext, formatType := range formats {
			for _, fileName := range formatType {
				name, solvedExt, mime, err := solveImgInfo(fileName)
				assert.Nil(t, err)
				assert.NotEmpty(t, mime)
				assert.Equal(t, ext, solvedExt)
				assert.Contains(t, fileName, name)
			}
		}
	})

	t.Run("InvalidFormatShouldThrowException", func(t *testing.T) {
		invalidSource := "i.wrong"
		img, _, _, err := solveImgInfo(invalidSource)
		assert.NotNil(t, err)
		assert.EqualError(t, errors.New("wrong is not a valid format"), err.Error())
		assert.Empty(t, img)
	})

	t.Run("ShouldNotFormatMalformedImageSource", func(t *testing.T) {
		malformerSource := "sdfs"
		format, _, _, err := solveImgInfo(malformerSource)
		assert.NotNil(t, err)
		assert.EqualError(t, fmt.Errorf(
			"%s is malformed", malformerSource), err.Error())
		assert.Empty(t, format)
	})
}
