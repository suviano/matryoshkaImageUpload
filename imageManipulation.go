package storageImage

import (
	"fmt"
	"strings"
)

var extensionsMap = map[string]string{
	"jpeg": "jpeg",
	"jpg":  "jpeg",
	"png":  "png",
	"bmp":  "bmp",
	"tiff": "tiff",
}

func buildImgName(fileName, extension string) string {
	return fmt.Sprintf("%s.%s", fileName, extension)
}

func fixImgExtension(source string) (string, string, error) {
	sourceParts := strings.Split(source, ".")
	if len(sourceParts) <= 1 {
		return "", "", fmt.Errorf("%s is malformed", source)
	}

	fileName := strings.Join(sourceParts[:len(sourceParts)-1], "/")
	originalExt := sourceParts[len(sourceParts)-1]

	ext, valid := extensionsMap[originalExt]
	if valid == false {
		return "", "", fmt.Errorf("%s is not a valid format", originalExt)
	}

	return buildImgName(fileName, ext), ext, nil
}

func generateImgsByScale(fileName, ext string) {
	return
}
