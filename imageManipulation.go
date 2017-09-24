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

const (
	mimeJpeg = "image/jpeg"
	mimePng  = "image/png"
)

var extensionsMap = map[string]string{
	"jpeg": mimeJpeg,
	"jpg":  mimeJpeg,
	"png":  mimePng,
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
	Path    string `json:"path"`
	MimeTyp string
}

func generateImgsByScale(buf *bytes.Buffer, prefix, fileName, ext, mimeTyp string) (map[string]*bufMedia, error) {
	bufMap := map[string]*bufMedia{
		"large":  {Size: 1200},
		"medium": {Size: 800},
		"small":  {Size: 400},
		"xsmall": {Size: 200},
		"original": {
			Size: 0,
			Buf:  bytes.NewBuffer(buf.Bytes()),
			Path: fmt.Sprintf("%s/%s.%s", prefix, fileName, ext),
		}}

	img, err := decodeImg(buf, mimeTyp)
	if err != nil {
		log.Warningf("{generateImgsByScale}{error decoding: %v}", err)
		return nil, err
	}

	for sizeName, imgBufResize := range bufMap {
		bufMap[sizeName].MimeTyp = mimeTyp
		if sizeName != "original" {
			bufMap[sizeName].Buf = bytes.NewBuffer(nil)
			bufMap[sizeName].Path = fmt.Sprintf("%s/%s-%s.%s", prefix, sizeName, fileName, ext)
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
