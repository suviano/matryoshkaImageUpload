package storageImage

import (
	"bytes"
	"context"
	"errors"

	log "github.com/sirupsen/logrus"
)

// WriteImage save images.
func WriteImage(prefix, filePath string, buf *bytes.Buffer) (err error) {
	if len(prefix) == 0 {
		return errors.New("empty prefix")
	}
	if buf == nil || buf.Len() == 0 {
		return errors.New("empty buffer")
	}

	fileName, ext, mimeTyp, err := fixImgExtension(filePath)
	if err != nil {
		log.Warningf("{WriteImage}{error solving image attr: %v}", err)
		return
	}

	imgMapBuf, err := generateImgsByScale(buf, prefix, fileName, ext, mimeTyp)
	if err != nil {
		log.Warningf("{WriteImage}{error generating imgs: %v}", err)
		return
	}

	ctx := context.Background()
	for _, imgBuffer := range imgMapBuf {
		err = storageClient.SaveImg(ctx, prefix, bucket, imgBuffer)
		if err != nil {
			log.Warningf("{WriteImage}{error saving image on storage: %v}", err)
			return
		}
	}
	return
}
