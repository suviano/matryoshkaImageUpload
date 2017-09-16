package storageImage

import (
	"fmt"
	"strings"
)

func defineFormat(source string) (string, error) {
	sourceParts := strings.Split(source, ".")
	if len(sourceParts) <= 1 {
		return "", fmt.Errorf("%s is malformed", source)
	}
	extension := sourceParts[len(sourceParts)-1]
	switch extension {
	case "jpeg", "jpg":
		return "jpeg", nil
	case "png":
		return "png", nil
	case "bmp":
		return "bmp", nil
	case "gif":
		return "gif", nil
	case "tiff":
		return "tiff", nil
	default:
		return "", fmt.Errorf("%s is not a valid format", extension)
	}
}
