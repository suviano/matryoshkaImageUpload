package matryoshka

import (
	"fmt"
	"strings"
)

const (
	mimeJpeg = "image/jpeg"
	mimePng  = "image/png"
)

var extensionsMap = map[string]string{
	"jpeg": "jpeg",
	"jpg":  "jpeg",
	"png":  "png",
}

var extMimeTypMap = map[string]string{
	"jpeg": mimeJpeg,
	"jpg":  mimeJpeg,
	"png":  mimePng,
}

func solveImgInfo(filePath string) (string, string, string, error) {
	sourceParts := strings.Split(filePath, ".")
	if len(sourceParts) < 2 {
		return "", "", "", fmt.Errorf("%s is malformed", filePath)
	}

	fileName := strings.Join(sourceParts[0:len(sourceParts)-1], ".")
	ext := sourceParts[len(sourceParts)-1]

	mimeTyp, valid := extMimeTypMap[ext]
	if valid == false {
		return "", "", "", fmt.Errorf("%s is not a valid format", ext)
	}

	return fileName, extensionsMap[ext], mimeTyp, nil
}
