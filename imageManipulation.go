package storageImage

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"strings"

	"github.com/nfnt/resize"
	log "github.com/sirupsen/logrus"
)

var extensionsMap = map[string]string{
	"jpeg": mimeJpeg,
	"jpg":  mimeJpeg,
	"png":  "image/png",
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
	ext := sourceParts[len(sourceParts)-1]

	mimeTyp, valid := extensionsMap[ext]
	if valid == false {
		return "", "", fmt.Errorf("%s is not a valid format", ext)
	}

	return buildImgName(fileName, ext), mimeTyp, nil
}

const mimeJpeg = "image/jpeg"

func decodeImg(buf *bytes.Buffer, ext string) (img image.Image, err error) {
	switch ext {
	case mimeJpeg:
		img, err = jpeg.Decode(buf)
	case "image/png":
		img, err = png.Decode(buf)
	default:
		err = errors.New("{Unsupported image type}")
	}
	return
}

func encodeImg(buf *bytes.Buffer, img image.Image, mimeTyp string) error {
	var err error
	switch mimeTyp {
	case mimeJpeg:
		err = jpeg.Encode(buf, img, nil)
	case "image/png":
		err = png.Encode(buf, img)
	}
	return err
}

type bufMedia struct {
	Buf  *bytes.Buffer
	Size uint
}

func generateImgsByScale(buf *bytes.Buffer, mimeTyp string) (map[string]bufMedia, error) {
	var bufMap = map[string]bufMedia{
		"original": {
			Buf:  bytes.NewBuffer(buf.Bytes()),
			Size: 0,
		},
		"large": {
			Buf:  bytes.NewBuffer(nil),
			Size: 1200,
		},
		"medium": {
			Buf:  bytes.NewBuffer(nil),
			Size: 800,
		},
		"small": {
			Buf:  bytes.NewBuffer(nil),
			Size: 400,
		},
		"extraSmall": {
			Buf:  bytes.NewBuffer(nil),
			Size: 200,
		},
	}

	img, err := decodeImg(buf, mimeTyp)
	if err != nil {
		log.Warningf("{generateImgsByScale}{error decoding: %v}", err)
		return nil, err
	}

	for sizeName, imgBufResize := range bufMap {
		if sizeName != "original" {
			tempImg := resize.Resize(imgBufResize.Size, 0, img, resize.Lanczos3)
			err = encodeImg(bufMap[sizeName].Buf, tempImg, mimeTyp)
			if err != nil {
				log.Warningf("{generateImgsByScale}{error encoding: [%s] [%v]}", sizeName, err)
				return nil, err
			}
		}
	}

	return bufMap, err
}
