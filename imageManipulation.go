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

func fixImgExtension(source string) (string, string, string, error) {
	sourceParts := strings.Split(source, ".")
	if len(sourceParts) <= 1 {
		return "", "", "", fmt.Errorf("%s is malformed", source)
	}

	fileName := strings.Join(sourceParts[:len(sourceParts)-1], "/")
	ext := sourceParts[len(sourceParts)-1]

	mimeTyp, valid := extensionsMap[ext]
	if valid == false {
		return "", "", "", fmt.Errorf("%s is not a valid format", ext)
	}

	return fileName, ext, mimeTyp, nil
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
	Buf     *bytes.Buffer
	Size    uint
	Path    string
	MimeTyp string
}

func generateImgsByScale(buf *bytes.Buffer, fileName, ext, mimeTyp string) (map[string]bufMedia, error) {
	var bufMap = map[string]bufMedia{
		"original": {
			Buf:     bytes.NewBuffer(buf.Bytes()),
			Size:    0,
			Path:    fmt.Sprintf("%s-original.%s", fileName, ext),
			MimeTyp: mimeTyp,
		},
		"large": {
			Buf:     bytes.NewBuffer(nil),
			Size:    1200,
			Path:    fmt.Sprintf("%s-large-%s", fileName, ext),
			MimeTyp: mimeTyp,
		},
		"medium": {
			Buf:     bytes.NewBuffer(nil),
			Size:    800,
			Path:    fmt.Sprintf("%s-medium-%s", fileName, ext),
			MimeTyp: mimeTyp,
		},
		"small": {
			Buf:     bytes.NewBuffer(nil),
			Size:    400,
			Path:    fmt.Sprintf("%s-small-%s", fileName, ext),
			MimeTyp: mimeTyp,
		},
		"extraSmall": {
			Buf:     bytes.NewBuffer(nil),
			Size:    200,
			Path:    fmt.Sprintf("%s-extraSmall-%s", fileName, ext),
			MimeTyp: mimeTyp,
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
